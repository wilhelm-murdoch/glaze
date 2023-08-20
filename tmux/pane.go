package tmux

import "fmt"

type Pane struct {
	Id                int
	Index             int
	Name              string
	StartingDirectory string
	IsActive          bool
	Window            *Window
}

func (p Pane) Target() string {
	return fmt.Sprintf(`%s:@%d.%%%d`, p.Window.Session.Name, p.Window.Id, p.Id)
}

func (p Pane) SendKeys(keys string) error {
	cmd, err := NewCommand("send", "-t", p.Target(), keys, "Enter")
	if err != nil {
		return err
	}

	return cmd.Exec()
}

func (p Pane) SetEnv(key string, value string) error {
	cmd, err := NewCommand("setenv", "-t", p.Name, key, value)
	if err != nil {
		return err
	}

	return cmd.Exec()
}

func (p Pane) Kill() error {
	cmd, err := NewCommand("killp", "-t", p.Target())
	if err != nil {
		return err
	}

	return cmd.Exec()
}
