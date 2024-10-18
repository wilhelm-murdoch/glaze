package tmux

import (
	"fmt"

	"github.com/wilhelm-murdoch/glaze/schema/pane"
)

// Pane represents a tmux pane.
type Pane struct {
	Window            *Window
	Name              string
	StartingDirectory string
	IsActive          bool
	IsFirst           bool
	Index             int
	Id                PaneId
}

// Target returns the target pane by its composite id of session name, window id, and pane id.
func (p Pane) Target() string {
	return fmt.Sprintf(`%s:%d.%d`, p.Window.Session.Name, p.Window.Index, p.Index)
}

// SendKeys sends the given keystrokes to the current pane.
func (p Pane) SendKeys(keys pane.Command) error {
	cmd, err := NewCommand(p.Window.Session.Client, "send", "-t", p.Target(), fmt.Sprint(keys), "Enter")
	if err != nil {
		return err
	}

	return cmd.Exec()
}

// SetEnv sets the given environment variable to the given value in the current pane.
func (p Pane) SetEnv(key pane.Name, value pane.Value) error {
	cmd, err := NewCommand(p.Window.Session.Client, "setenv", "-t", p.Name, fmt.Sprint(key), fmt.Sprint(value))
	if err != nil {
		return err
	}

	return cmd.Exec()
}

func (p Pane) Select() error {
	cmd, err := NewCommand(p.Window.Session.Client, "selectp", "-t", p.Target())
	if err != nil {
		return err
	}

	return cmd.Exec()
}

// Kill closes the current pane.
func (p Pane) Kill() error {
	cmd, err := NewCommand(p.Window.Session.Client, "killp", "-t", p.Target())
	if err != nil {
		return err
	}

	return cmd.Exec()
}
