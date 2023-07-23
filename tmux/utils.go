package tmux

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func Exec(args ...string) error {
	tmux, ok := IsInstalled()
	if !ok {
		return fmt.Errorf("tmux is not installed")
	}

	cmd := exec.Command(tmux, args...)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

func ExecWithOutput(args ...string) (string, error) {
	tmux, ok := IsInstalled()
	if !ok {
		return "", fmt.Errorf("tmux is not installed")
	}

	cmd := exec.Command(tmux, args...)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}

	return strings.TrimSuffix(string(output), "\n"), nil
}

func IsInstalled() (string, bool) {
	path, err := exec.LookPath("tmux")
	return path, err == nil
}

func IsInsideTmux() bool {
	if os.Getenv("TMUX") != "" {
		return true
	} else {
		return false
	}
}
