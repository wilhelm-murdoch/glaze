package tmux

import (
	"os"
	"os/exec"
	"strings"
)

func IsInstalled() (string, bool) {
	path, err := exec.LookPath("tmux")
	return path, err == nil
}

func IsInsideTmux() bool {
	return os.Getenv("TMUX") != ""
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
