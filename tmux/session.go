package tmux

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/wilhelm-murdoch/glaze/schema/menu"
	"github.com/wilhelm-murdoch/glaze/schema/session"
	"github.com/wilhelm-murdoch/glaze/schema/window"
	"github.com/wilhelm-murdoch/glaze/tmux/enums"
	"github.com/wilhelm-murdoch/go-collection"
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
func (s *Session) NewWindow(windowName window.Name) (*Window, error) {
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
		"-n", fmt.Sprint(windowName),
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

	baseIndex, err := s.GetOption(enums.OptionsSessionBaseIndex)
	if err != nil {
		return window, err
	}

	return &Window{
		Id:       id,
		Index:    index,
		Name:     parts[2],
		Layout:   enums.LayoutFromString(parts[3]),
		IsActive: parts[4] == "1",
		IsFirst:  parts[1] == baseIndex.Value,
		Session:  s,
	}, nil
}

func (s *Session) SetMenu(menu *menu.Menu) error {
	cmd, err := NewCommand(s.Client, menu.CommandArgs()...)
	if err != nil {
		return err
	}

	return cmd.Exec()
}

func (s *Session) SetOption(option session.Name, value session.Value) error {
	return setOption[enums.OptionsSession](s.Client, "set", "-t", s.Target(), fmt.Sprint(option), fmt.Sprint(value))
}

func (s *Session) GetOption(option enums.OptionsSession) (Option[enums.OptionsSession], error) {
	return getOption[enums.OptionsSession](s.Client, "show", "-g", "-t", s.Target(), fmt.Sprint(option))
}

func (s *Session) ShowOptions() (collection.Collection[Option[enums.OptionsSession]], error) {
	return showOptions[enums.OptionsSession](s.Client, "show", "-g", "-t", s.Target())
}

// Kill closes the current session.
func (s *Session) Kill() error {
	cmd, err := NewCommand(s.Client, "kill-session", "-t", s.Target())
	if err != nil {
		return err
	}

	return cmd.Exec()
}
