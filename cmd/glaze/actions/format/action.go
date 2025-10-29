package format

import (
	"fmt"
	"os"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/urfave/cli/v2"

	"github.com/wilhelm-murdoch/glaze/internal/diagnostics"
	"github.com/wilhelm-murdoch/glaze/internal/parser"
	"github.com/wilhelm-murdoch/glaze/internal/profile"
	"github.com/wilhelm-murdoch/glaze/internal/schema"
)

type Action struct {
	ctx          *cli.Context
	diagsManager *diagnostics.DiagnosticsManager
	parser       *parser.Parser
	profilePath  string
}

func NewAction(ctx *cli.Context) (*Action, error) {
	profilePath, err := profile.ResolveProfilePath(ctx.Args().First())
	if err != nil {
		return nil, err
	}

	diagsManager := diagnostics.NewDiagnosticsManager(profilePath)
	if diagsManager.HasErrors() {
		return nil, diagsManager.Write()
	}

	parser, parserDiags := parser.NewParser(profilePath)
	if parserDiags.HasErrors() {
		diagsManager.Extend(parserDiags)
		return nil, diagsManager.Write()
	}

	return &Action{
		ctx:          ctx,
		diagsManager: diagsManager,
		parser:       parser,
		profilePath:  profilePath,
	}, nil
}

// Run is an action that will reformat the given glaze definition file to match
// a canonical format and style, ensuring consistency. If a `stdout` flag is not
// passed through via the cli, this command will attempt to overwrite the given
// file with reformatted output.
func (a *Action) Run() error {
	formatted := string(hclwrite.Format(a.parser.File.Bytes))

	if a.ctx.Bool("validate") {
		if valid := a.isGlazeDefintionValid(); !valid {
			return a.diagsManager.Write()
		}
	}

	if a.ctx.Bool("stdout") {
		fmt.Print(formatted)
		return nil
	}

	if err := os.WriteFile(a.profilePath, []byte(formatted), 0o644); err != nil {
		a.diagsManager.Append(&hcl.Diagnostic{
			Severity: hcl.DiagError,
			Summary:  "Failed to write file",
			Detail:   err.Error(),
		})
	}

	if a.diagsManager.HasErrors() {
		return a.diagsManager.Write()
	}

	return nil
}

// isGlazeDefintionValid checks if the given glaze definition file and any variable
// flags yield a valid result when run through the schema parser.
func (a *Action) isGlazeDefintionValid() bool {
	variables, err := parser.CollectVariables(a.ctx.StringSlice("var"))
	if err != nil {
		a.diagsManager.Append(&hcl.Diagnostic{
			Severity: hcl.DiagError,
			Summary:  fmt.Sprintf("could not parse specified variables: %s", err),
			Detail:   err.Error(),
		})

		return false
	}

	_, decodeDiags := a.parser.Decode(
		schema.PrimaryGlazeSpec,
		parser.BuildEvalContext(variables),
	)

	if decodeDiags.HasErrors() {
		a.diagsManager.Extend(decodeDiags)
		return false
	}

	return true
}
