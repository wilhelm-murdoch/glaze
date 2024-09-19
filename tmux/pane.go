package tmux

import "fmt"

// Pane represents a tmux pane.
type Pane struct {
	Window            *Window
	Name              string
	StartingDirectory string
	IsActive          bool
	IsFirst           bool
	Index             int
	Id                int
}

// Target returns the target pane by its composite id of session name, window id, and pane id.
func (p Pane) Target() string {
	return fmt.Sprintf(`%s:@%d.%%%d`, p.Window.Session.Name, p.Window.Id, p.Id)
}

// SendKeys sends the given keystrokes to the current pane.
func (p Pane) SendKeys(keys string) error {
	cmd, err := NewCommand(p.Window.Session.Client, "send", "-t", p.Target(), keys, "Enter")
	if err != nil {
		return err
	}

	return cmd.Exec()
}

// SetEnv sets the given environment variable to the given value in the current pane.
func (p Pane) SetEnv(key string, value string) error {
	cmd, err := NewCommand(p.Window.Session.Client, "setenv", "-t", p.Name, key, value)
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
