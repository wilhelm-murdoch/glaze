package tmux

type Pane struct {
	Id                int
	Index             int
	Name              string
	StartingDirectory string
	IsActive          bool
}

func (c Client) SendKeys(pane Pane, keys string) error {
	return Exec("send-keys", "-t", pane.Name, keys, "Enter")
}

func (c Client) SetEnv(pane Pane, key string, value string) error {
	return Exec("setenv", "-t", pane.Name, key, value)
}
