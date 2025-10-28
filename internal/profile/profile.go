package profile

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/wilhelm-murdoch/glaze/pkg/files"
)

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
