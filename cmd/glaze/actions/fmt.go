package actions

import (
	"os"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/wilhelm-murdoch/glaze"
)

func ActionFmt(profilePath string) error {
	parser := glaze.NewParser(profilePath)

	if parser.HasErrors() {
		parser.WriteDiags()
		return nil
	}

	formatted := string(hclwrite.Format(parser.File.Bytes))

	if err := os.WriteFile(profilePath, []byte(formatted), 0644); err != nil {
		parser.AppendDiag(&hcl.Diagnostic{
			Severity: hcl.DiagError,
			Summary:  "Failed to write file",
			Detail:   err.Error(),
		})
	}

	if parser.HasErrors() {
		parser.WriteDiags()
		return nil
	}

	return nil
}
