package actions

import (
	"os"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/wilhelm-murdoch/glaze"
)

func ActionFmt(profilePath string) error {
	p := glaze.NewParser()
	p.Open(profilePath)

	if p.HasErrors() {
		p.WriteDiags()
		return nil
	}

	formatted := string(hclwrite.Format(p.File.Bytes))

	if err := os.WriteFile(profilePath, []byte(formatted), 0644); err != nil {
		p.AppendDiag(&hcl.Diagnostic{
			Severity: hcl.DiagError,
			Summary:  "Failed to write file",
			Detail:   err.Error(),
		})
	}

	if p.HasErrors() {
		p.WriteDiags()
		return nil
	}

	return nil
}
