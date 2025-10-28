package files

import (
	"errors"
	"io/fs"
	"os"
	"strings"
)

func FileExists(path string) bool {
	fileInfo, err := os.Stat(path)
	if err != nil || errors.Is(err, fs.ErrNotExist) || fileInfo.IsDir() {
		return false
	}

	return true
}

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
