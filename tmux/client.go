package tmux

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/wilhelm-murdoch/go-collection"
)

type Client struct {
	socketPath     string
	socketName     string
	CurrentSession *Session
}

func NewClient() Client {
	return Client{}
}

func NewClientWithSocket(socketPath string, socketName string) Client {
	return Client{
		socketPath: socketPath,
		socketName: socketName,
	}
}

func (c *Client) Attach(session *Session) error {
	var args []string

	if !IsInsideTmux() {
		args = append(args, "attach-session", "-t", session.Target())
	} else {
		args = append(args, "switch-client", "-t", session.Target())
	}

	cmd, err := NewCommand(args...)
	if err != nil {
		return err
	}

	if err := cmd.Exec(); err != nil {
		if strings.Contains(err.Error(), "can't find session") {
			return fmt.Errorf(`session "%s" not found`, session.Name)
		}

		return err
	}

	c.CurrentSession = session

	return nil
}

func (c Client) Sessions() (collection.Collection[*Session], error) {
	var sessions collection.Collection[*Session]

	format := []string{
		"#{session_id}",
		"#{session_name}",
		"#{session_path}",
	}

	cmd, err := NewCommand("ls", "-F", strings.Join(format, ";"))
	if err != nil {
		return sessions, err
	}

	output, err := cmd.ExecWithOutput()
	if err != nil {
		return sessions, err
	}

	for _, line := range strings.Split(output, "\n") {
		parts := strings.SplitN(line, ";", 3)

		id, err := strconv.Atoi(strings.Replace(parts[0], "$", "", -1))
		if err != nil {
			return sessions, err
		}

		sessions.Push(&Session{
			Id:                id,
			Name:              strings.TrimSpace(parts[1]),
			StartingDirectory: strings.TrimSpace(parts[2]),
		})
	}

	return sessions, err
}

func (c Client) Windows(session *Session) (collection.Collection[*Window], error) {
	var windows collection.Collection[*Window]

	format := []string{
		"#{window_id}",
		"#{window_index}",
		"#{window_name}",
		"#{window_layout}",
		"#{window_active}",
	}

	cmd, err := NewCommand("list-windows", "-F", strings.Join(format, ";"), "-t", session.Target())
	if err != nil {
		return windows, err
	}

	output, err := cmd.ExecWithOutput()
	if err != nil {
		return windows, err
	}

	for _, window := range strings.Split(output, "\n") {
		parts := strings.SplitN(window, ";", 5)

		id, err := strconv.Atoi(strings.Replace(parts[0], "@", "", -1))
		if err != nil {
			return windows, err
		}

		index, err := strconv.Atoi(parts[1])
		if err != nil {
			return windows, err
		}

		windows.Push(&Window{
			Id:       id,
			Index:    index,
			Name:     parts[2],
			Layout:   parts[3],
			IsActive: parts[4] == "1",
			Session:  session,
		})
	}

	return windows, nil
}

func (c Client) Panes(window *Window) (collection.Collection[*Pane], error) {
	var panes collection.Collection[*Pane]

	format := []string{
		"#{pane_id}",
		"#{pane_index}",
		"#{pane_title}",
		"#{pane_active}",
		"#{pane_current_path}",
	}

	cmd, err := NewCommand("list-panes", "-F", strings.Join(format, ";"), "-t", window.Target())
	if err != nil {
		return panes, err
	}

	output, err := cmd.ExecWithOutput()
	if err != nil {
		return panes, err
	}

	for _, pane := range strings.Split(output, "\n") {
		parts := strings.SplitN(pane, ";", 5)

		id, err := strconv.Atoi(strings.Replace(parts[0], "%", "", -1))
		if err != nil {
			return panes, err
		}

		index, err := strconv.Atoi(parts[1])
		if err != nil {
			return panes, err
		}

		panes.Push(&Pane{
			Id:                id,
			Index:             index,
			Name:              parts[2],
			StartingDirectory: parts[4],
			IsActive:          parts[3] == "1",
			Window:            window,
		})
	}

	return panes, nil
}

func (c Client) NewSession(sessionName, startingDirectory string) (*Session, error) {
	var session *Session

	cmd, err := NewCommand("new", "-d", "-s", sessionName, "-c", startingDirectory)
	if err != nil {
		return session, err
	}

	if err := cmd.Exec(); err != nil {
		return session, err
	}

	sessions, err := c.Sessions()
	if err != nil {
		return session, err
	}

	return sessions.Find(func(i int, s *Session) bool {
		return s.Name == sessionName
	}), nil
}

func (c Client) KillSessionByName(sessionName string) error {
	sessions, err := c.Sessions()
	if err != nil {
		return err
	}

	found := sessions.Find(func(i int, s *Session) bool {
		return s.Name == sessionName
	})

	if found == nil {
		return fmt.Errorf(`session "%s" not found`, sessionName)
	}

	cmd, err := NewCommand("kill-session", "-t", found.Target())
	if err != nil {
		return err
	}

	return cmd.Exec()
}

func (c Client) SessionExists(name string) bool {
	sessions, _ := c.Sessions()
	return sessions.Find(func(i int, s *Session) bool {
		return s.Name == name
	}) != nil
}
