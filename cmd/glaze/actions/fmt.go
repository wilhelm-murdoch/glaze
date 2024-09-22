package actions

import (
	"fmt"
	"os"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/urfave/cli/v2"
	"github.com/wilhelm-murdoch/glaze"
)

func Fmt(ctx *cli.Context) error {
	profilePath, err := glaze.ResolveProfilePath(ctx.Args().First())
	if err != nil {
		return err
	}

	diagsManager := glaze.NewDiagnosticsManager(profilePath)
	if diagsManager.HasErrors() {
		return diagsManager.Write()
	}

	if diagsManager.HasErrors() {
		return diagsManager.Write()
	}

	parser, parserDiags := glaze.NewParser(profilePath)
	if parserDiags.HasErrors() {
		diagsManager.Extend(parserDiags)
		return diagsManager.Write()
	}

	formatted := string(hclwrite.Format(parser.File.Bytes))

	if ctx.Bool("stdout") {
		fmt.Print(formatted)
		return nil
	}

	if err := os.WriteFile(profilePath, []byte(formatted), 0644); err != nil {
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
