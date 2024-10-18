package glaze

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"strings"

	"github.com/kr/pretty"
)

func FileExists(path string) bool {
	fileInfo, err := os.Stat(path)
	if err != nil || errors.Is(err, fs.ErrNotExist) || fileInfo.IsDir() {
		return false
	}

	return true
}

func JoinWithOr(choices []string) string {
	return joinWith(choices, "or")
}

func JoinWithAnd(choices []string) string {
	return joinWith(choices, "and")
}

func joinWith(choices []string, conjunction string) string {
	length := len(choices)
	switch length {
	case 0:
		return ""
	case 1:
		return choices[0]
	case 2:
		return fmt.Sprintf(`%s %s %s`, choices[0], conjunction, choices[1])
	}

	return fmt.Sprintf(`%s %s %s`, strings.Join(choices[:length-1], ", "), conjunction, choices[length-1])
}

func Prettier(values ...any) {
	for _, value := range values {
		fmt.Printf("%# v\n", pretty.Formatter(value))
	}
}

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
