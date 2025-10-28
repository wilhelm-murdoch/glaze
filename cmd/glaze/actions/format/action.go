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
)

// Run is an action that will reformat the given glaze definition file to match
// a canonical format and style, ensuring consistency. If a `stdout` flag is not
// passed through via the cli, this command will attempt to overwrite the given
// file with reformatted output.
func Run(ctx *cli.Context) error {
	profilePath, err := profile.ResolveProfilePath(ctx.Args().First())
	if err != nil {
		return err
	}

	diagsManager := diagnostics.NewDiagnosticsManager(profilePath)
	if diagsManager.HasErrors() {
		return diagsManager.Write()
	}

	if diagsManager.HasErrors() {
		return diagsManager.Write()
	}

	glazeParser, parserDiags := parser.NewParser(profilePath)
	if parserDiags.HasErrors() {
		diagsManager.Extend(parserDiags)
		return diagsManager.Write()
	}

	formatted := string(hclwrite.Format(glazeParser.File.Bytes))

	if ctx.Bool("stdout") {
		fmt.Print(formatted)
		return nil
	}

	if err := os.WriteFile(profilePath, []byte(formatted), 0o644); err != nil {
		diagsManager.Append(&hcl.Diagnostic{
			Severity: hcl.DiagError,
			Summary:  "Failed to write file",
			Detail:   err.Error(),
		})
	}

	if diagsManager.HasErrors() {
		return diagsManager.Write()
	}

	return nil
}
