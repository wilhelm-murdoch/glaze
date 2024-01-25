package tmux

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/wilhelm-murdoch/glaze/tmux/enums"
)

// Session represents a tmux session.
type Session struct {
	Id                int
	Name              string
	StartingDirectory string
}

// Target returns the target session by its name.
func (s *Session) Target() string {
	return s.Name
}

// NewWindow creates a new window in the current session and returns it.
func (s *Session) NewWindow(windowName string) (*Window, error) {
	var window *Window

	format := []string{
		"#{window_id}",
		"#{window_index}",
		"#{window_name}",
		"#{window_layout}",
		"#{window_active}",
	}

	args := []string{
		"neww",
		"-d",
		"-t", fmt.Sprintf("%s:", s.Name),
		"-n", windowName,
		"-F", strings.Join(format, ";"),
		"-P",
	}

	cmd, err := NewCommand(args...)
	if err != nil {
		return window, err
	}

	output, err := cmd.ExecWithOutput()
	if err != nil {
		return window, err
	}

	parts := strings.SplitN(output, ";", 5)

	id, err := strconv.Atoi(strings.Replace(parts[0], "@", "", -1))
	if err != nil {
		return window, err
	}

	index, err := strconv.Atoi(parts[1])
	if err != nil {
		return window, err
	}

	return &Window{
		Id:       id,
		Index:    index,
		Name:     parts[2],
		Layout:   enums.LayoutFromString(parts[3]),
		IsActive: parts[4] == "1",
		Session:  s,
	}, nil
}

// Kill closes the current session.
func (s *Session) Kill() error {
	cmd, err := NewCommand("kill-session", "-t", s.Target())
	if err != nil {
		return err
	}

	return cmd.Exec()
}
