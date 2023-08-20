package tmux

import (
	"fmt"
	"strconv"
	"strings"
)

type Session struct {
	Id                int
	Name              string
	StartingDirectory string
}

func (s *Session) Target() string {
	return s.Name
}

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
		Layout:   parts[3],
		IsActive: parts[4] == "1",
		Session:  s,
	}, nil
}

func (s *Session) Kill() error {
	cmd, err := NewCommand("kill-session", "-t", s.Target())
	if err != nil {
		return err
	}

	return cmd.Exec()
}
