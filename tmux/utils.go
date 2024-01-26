package tmux

import (
	"os"
	"os/exec"
	"strings"
)

// IsInstalled returns true if tmux is installed on the system and also returns
// the path to the associated binary.
func IsInstalled() (bool, string) {
	path, err := exec.LookPath("tmux")
	return err == nil, path
}

// IsInsideTmux checks if we are inside a tmux session. We assume we are in
// a tmux session when the `$TMUX` environment variable is set.
func IsInsideTmux() bool {
	return os.Getenv("TMUX") != ""
}

// ExpandPath expands the given path. If the path starts with "~/", we replace
// it with the absolute path to the user's home directory.
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

// Contains returns true if the given value is in the given list.
func Contains(haystack []string, needle string) bool {
	for _, v := range haystack {
		if string(v) == needle {
			return true
		}
	}

	return false
}
