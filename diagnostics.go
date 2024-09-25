package glaze

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"strings"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclparse"
	"github.com/wilhelm-murdoch/glaze/tmux"
	"github.com/zclconf/go-cty/cty"
)

type DiagnosticsManager struct {
	hcl.Diagnostics
	hcl.DiagnosticWriter
}

func (dm *DiagnosticsManager) Write() error {
	if err := dm.DiagnosticWriter.WriteDiagnostics(dm.Diagnostics); err != nil {
		return err
	}

	return nil
}

func (dm *DiagnosticsManager) Extend(diags hcl.Diagnostics) hcl.Diagnostics {
	dm.Diagnostics = dm.Diagnostics.Extend(diags)
	return dm.Diagnostics
}

func NewDiagnosticsManager(filePath string) *DiagnosticsManager {
	parser := hclparse.NewParser()
	file, diags := parser.ParseHCLFile(filePath)

	diagsManager := &DiagnosticsManager{
		DiagnosticWriter: hcl.NewDiagnosticTextWriter(os.Stdout, map[string]*hcl.File{filePath: file}, 78, true),
	}

	if diags.HasErrors() {
		diagsManager.Extend(diags)
	}

	return diagsManager
}

func ContainsDiagnostic(field string, value cty.Value, list []string) hcl.Diagnostics {
	var out hcl.Diagnostics

	if !value.IsNull() && !tmux.Contains(list, value.AsString()) {
		return hcl.Diagnostics{{
			Severity: hcl.DiagError,
			Summary:  fmt.Sprintf(`Invalid %s specified`, field),
			Detail:   fmt.Sprintf(`The %s value of "%s" is not supported among: %s.`, field, value.AsString(), strings.Join(list, ", ")),
		}}
	}

	return out
}

func DirectoryDiagnostic(field string, value cty.Value) hcl.Diagnostics {
	var out hcl.Diagnostics

	if !value.IsNull() {
		fileInfo, err := os.Stat(tmux.ExpandPath(value.AsString()))
		if err != nil || errors.Is(err, fs.ErrNotExist) || !fileInfo.IsDir() {
			return hcl.Diagnostics{{
				Severity: hcl.DiagError,
				Summary:  fmt.Sprintf(`Invalid %s specified`, field),
				Detail:   fmt.Sprintf(`The %s of "%s" does not exist or is not a directory.`, field, value.AsString()),
			}}
		}
	}

	return out
}

func WrongAttributeDiagnostic(field, have, want string) hcl.Diagnostic {
	return hcl.Diagnostic{
		Severity: hcl.DiagError,
		Summary:  fmt.Sprintf(`Invalid %s specified`, field),
		Detail:   fmt.Sprintf(`The %s value "%s" should be "%s".`, field, have, want),
	}
}
