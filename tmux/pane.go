package tmux

import "fmt"

type Pane struct {
	Id                int
	Index             int
	Name              string
	StartingDirectory string
	IsActive          bool
	WindowId          int
	SessionId         int
}

func (p Pane) SendKeys(keys string) error {
	return Exec("send", "-t", fmt.Sprintf(`%d:%d.%d`, p.SessionId, p.WindowId, p.Index), keys, "Enter")
}

func (p Pane) SetEnv(key string, value string) error {
	return Exec("setenv", "-t", p.Name, key, value)
}
