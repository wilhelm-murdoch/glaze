package tmux

import (
	"strconv"
	"strings"

	"github.com/wilhelm-murdoch/glaze/tmux/enums"
)

// Session represents a tmux session.
type Session struct {
	Client            Client
	Name              string
	StartingDirectory string
	Id                int
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
		"-t", s.Name,
		"-n", windowName,
		"-F", strings.Join(format, ";"),
		"-P",
	}

	cmd, err := NewCommand(s.Client, args...)
	if err != nil {
		return window, err
	}

	output, err := cmd.ExecWithOutput()
	if err != nil {
		return window, err
	}

	parts := strings.SplitN(output, ";", len(format))

	id, err := strconv.Atoi(strings.ReplaceAll(parts[0], "@", ""))
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
		IsFirst:  index == 0,
		Session:  s,
	}, nil
}

func (s *Session) SetOption(option string, value string) error {
	cmd, err := NewCommand(s.Client, "set", option, value)
	if err != nil {
		return err
	}

	return cmd.Exec()
}

// Kill closes the current session.
func (s *Session) Kill() error {
	cmd, err := NewCommand(s.Client, "kill-session", "-t", s.Target())
	if err != nil {
		return err
	}

	return cmd.Exec()
}
