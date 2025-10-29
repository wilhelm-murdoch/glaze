package diagnostics

import (
	"os"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclparse"
)

type DiagnosticsManager struct {
	Diagnostics      hcl.Diagnostics
	DiagnosticWriter hcl.DiagnosticWriter
}

func (dm *DiagnosticsManager) Write() error {
	return dm.DiagnosticWriter.WriteDiagnostics(dm.Diagnostics)
}

func (dm *DiagnosticsManager) Extend(diags hcl.Diagnostics) hcl.Diagnostics {
	dm.Diagnostics = dm.Diagnostics.Extend(diags)
	return dm.Diagnostics
}

func (dm *DiagnosticsManager) Append(diag *hcl.Diagnostic) hcl.Diagnostics {
	return dm.Diagnostics.Append(diag)
}

func (dm *DiagnosticsManager) HasErrors() bool {
	return dm.Diagnostics.HasErrors()
}

func NewDiagnosticsManager(filePath string) *DiagnosticsManager {
	parser := hclparse.NewParser()
	file, diags := parser.ParseHCLFile(filePath)

	diagsManager := &DiagnosticsManager{
		DiagnosticWriter: hcl.NewDiagnosticTextWriter(
			os.Stdout,
			map[string]*hcl.File{filePath: file},
			78,
			true,
		),
	}

	if diags.HasErrors() {
		diagsManager.Extend(diags)
	}

	return diagsManager
}
