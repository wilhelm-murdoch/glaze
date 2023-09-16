package tmux

import (
	"os"
	"os/exec"
)

func IsInstalled() (string, bool) {
	path, err := exec.LookPath("tmux")
	return path, err == nil
}

func IsInsideTmux() bool {
	return os.Getenv("TMUX") != ""
}
