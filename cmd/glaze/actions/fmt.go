package actions

import (
	"fmt"
	"os"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/urfave/cli/v2"
	"github.com/wilhelm-murdoch/glaze"
)

func ActionFmt(ctx *cli.Context) error {
	profilePath := ctx.Args().First()

	parser := glaze.NewParser(profilePath)

	if parser.HasErrors() {
		return parser.WriteDiags()
	}

	formatted := string(hclwrite.Format(parser.File.Bytes))

	if ctx.Bool("stdout") {
		fmt.Print(formatted)
		return nil
	}

	if err := os.WriteFile(profilePath, []byte(formatted), 0644); err != nil {
		parser.AppendDiag(&hcl.Diagnostic{
			Severity: hcl.DiagError,
			Summary:  "Failed to write file",
			Detail:   err.Error(),
		})
	}

	if parser.HasErrors() {
		return parser.WriteDiags()
	}

	return nil
}
