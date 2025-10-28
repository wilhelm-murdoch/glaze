package files

import (
	"errors"
	"io/fs"
	"os"
	"strings"
)

// FileExists is a utility function that simply checks if the
// given path is not only a file, but that it exists and is
// readable.
func FileExists(path string) bool {
	fileInfo, err := os.Stat(path)
	if err != nil || errors.Is(err, fs.ErrNotExist) || fileInfo.IsDir() {
		return false
	}

	return true
}

// ExpandPath is a utility function that determines whether the
// given path is a shortcut to a user's home directory. If it is
// it returns the associated absolute path.
func ExpandPath(path string) string {
	if strings.HasPrefix(path, "~/") {
		userHome, err := os.UserHomeDir()
		if err != nil {
			return path
		}

		return strings.Replace(path, "~", userHome, 1)
	}

	return path
}
