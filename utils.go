package glaze

import (
	"errors"
	"io/fs"
	"os"
)

func FileExists(path string) bool {
	fileInfo, err := os.Stat(path)
	if err != nil || errors.Is(err, fs.ErrNotExist) || fileInfo.IsDir() {
		return false
	}

	return true
}
