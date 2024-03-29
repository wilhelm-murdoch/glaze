package tmux

import "fmt"

// Pane represents a tmux pane.
type Pane struct {
	Id                int
	Index             int
	Name              string
	StartingDirectory string
	IsActive          bool
	IsFirst           bool
	Window            *Window
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

// Kill closes the current pane.
func (p Pane) Kill() error {
	cmd, err := NewCommand(p.Window.Session.Client, "killp", "-t", p.Target())
	if err != nil {
		return err
	}

	return cmd.Exec()
}
