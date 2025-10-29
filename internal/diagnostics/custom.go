package diagnostics

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"slices"
	"strings"

	"github.com/hashicorp/hcl/v2"
	"github.com/zclconf/go-cty/cty"

	"github.com/wilhelm-murdoch/glaze/pkg/files"
)

func ContainsDiagnostic(field string, value cty.Value, list []string) hcl.Diagnostics {
	var out hcl.Diagnostics

	if !value.IsNull() && !slices.Contains(list, value.AsString()) {
		return hcl.Diagnostics{{
			Severity: hcl.DiagError,
			Summary:  fmt.Sprintf(`Invalid %s specified`, field),
			Detail: fmt.Sprintf(
				`The %s value of "%s" is not supported among: %s.`,
				field,
				value.AsString(),
				strings.Join(list, ", "),
			),
		}}
	}

	return out
}

func DirectoryDiagnostic(field string, value cty.Value) hcl.Diagnostics {
	var out hcl.Diagnostics

	if !value.IsNull() {
		fileInfo, err := os.Stat(files.ExpandPath(value.AsString()))
		if err != nil || errors.Is(err, fs.ErrNotExist) || !fileInfo.IsDir() {
			return hcl.Diagnostics{{
				Severity: hcl.DiagError,
				Summary:  fmt.Sprintf(`Invalid %s specified`, field),
				Detail: fmt.Sprintf(
					`The %s of "%s" does not exist or is not a directory.`,
					field,
					value.AsString(),
				),
			}}
		}
	}

	return out
}

func FileDiagnostic(field string, value cty.Value) hcl.Diagnostics {
	var out hcl.Diagnostics

	if !value.IsNull() {
		fileInfo, err := os.Stat(files.ExpandPath(value.AsString()))
		if err != nil || errors.Is(err, fs.ErrNotExist) || fileInfo.IsDir() {
			return hcl.Diagnostics{{
				Severity: hcl.DiagError,
				Summary:  fmt.Sprintf(`Invalid %s specified`, field),
				Detail: fmt.Sprintf(
					`The %s of "%s" does not exist, cannot be accessed or is a directory.`,
					field,
					value.AsString(),
				),
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
