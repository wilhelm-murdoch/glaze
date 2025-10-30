package profile

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/wilhelm-murdoch/glazier/pkg/files"
)

// ResolveProfilePath attempts to determine the most reasonable path to a glaze definition file.
// If a `profilePath` is not given, it assumes we wish to look within the current path in which
// the glaze cli was executed. Failing that, it attempts to read a path from a `GLAZE_PATH`
// environment variable.
func ResolveProfilePath(profilePath string) (string, error) {
	if profilePath == "" {
		cwd, err := os.Getwd()
		if err != nil {
			return profilePath, fmt.Errorf("could not read current working directory: %s", err)
		}

		profilePath = filepath.Join(cwd, ".glaze")

		if !files.FileExists(profilePath) && os.Getenv("GLAZE_PATH") != "" {
			profilePath = filepath.Join(os.Getenv("GLAZE_PATH"), ".glaze")
		}
	}

	if !files.FileExists(profilePath) {
		return profilePath, fmt.Errorf(
			"profile `%s` not found on the specified path, the current directory, or the GLAZE_PATH environment variable",
			profilePath,
		)
	}

	return profilePath, nil
}
