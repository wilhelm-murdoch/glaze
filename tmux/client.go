package tmux

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/wilhelm-murdoch/glaze/tmux/enums"
	"github.com/wilhelm-murdoch/go-collection"
)

// Client represents a tmux client.
type Client struct {
	socketPath     string
	socketName     string
	CurrentSession *Session
}

// NewClient returns a new client.
func NewClient(socketPath, socketName string) Client {
	return Client{
		socketPath: socketPath,
		socketName: socketName,
	}
}

// Attach attaches to the given session. If we are inside a tmux session,
// we switch to the given session.
func (c *Client) Attach(session *Session) error {
	var args []string

	if !IsInsideTmux() {
		args = append(args, "attach", "-t", session.Target())
	} else {
		args = append(args, "switchc", "-t", session.Target())
	}

	cmd, err := NewCommand(*c, args...)
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

// Sessions returns a collection of active sessions.
func (c Client) Sessions() (collection.Collection[*Session], error) {
	var sessions collection.Collection[*Session]

	format := []string{
		"#{session_id}",
		"#{session_name}",
		"#{session_path}",
	}

	args := []string{
		"ls",
		"-F", strings.Join(format, ";"),
	}

	cmd, err := NewCommand(c, args...)
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
			Client:            c,
			Id:                id,
			Name:              strings.TrimSpace(parts[1]),
			StartingDirectory: strings.TrimSpace(parts[2]),
		})
	}

	return sessions, err
}

// Windows returns a collection of windows for the given session.
func (c Client) Windows(session *Session) (collection.Collection[*Window], error) {
	var windows collection.Collection[*Window]

	format := []string{
		"#{window_id}",
		"#{window_index}",
		"#{window_name}",
		"#{window_layout}",
		"#{window_active}",
	}

	args := []string{
		"lsw",
		"-F", strings.Join(format, ";"),
		"-t", session.Target(),
	}

	cmd, err := NewCommand(c, args...)
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
			Layout:   enums.LayoutFromString(parts[3]),
			IsActive: parts[4] == "1",
			Session:  session,
		})
	}

	return windows, nil
}

// Panes returns a collection of panes for the given window.
func (c Client) Panes(window *Window) (collection.Collection[*Pane], error) {
	var panes collection.Collection[*Pane]

	format := []string{
		"#{pane_id}",
		"#{pane_index}",
		"#{pane_title}",
		"#{pane_active}",
		"#{pane_current_path}",
	}

	args := []string{
		"lsp",
		"-F", strings.Join(format, ";"),
		"-t", window.Target(),
	}

	cmd, err := NewCommand(c, args...)
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

// NewSession creates a new session with the given name and starting directory.
func (c Client) NewSession(sessionName, startingDirectory string) (*Session, error) {
	var session *Session

	cmd, err := NewCommand(c, "new", "-d", "-s", sessionName, "-c", startingDirectory)
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

// NewSessionIfNotExists creates a new session with the given name and starting
// directory if it does not already exist.
func (c Client) NewSessionIfNotExists(sessionName, startingDirectory string) (*Session, error) {
	sessions, _ := c.Sessions()
	exists := sessions.Find(func(i int, s *Session) bool {
		return s.Name == sessionName
	})

	if exists == nil {
		return c.NewSession(sessionName, startingDirectory)
	}

	return exists, nil
}

// KillSession kills the given session.
func (c Client) KillSessionByName(sessionName string) error {
	sessions, _ := c.Sessions()

	found := sessions.Find(func(i int, s *Session) bool {
		return s.Name == sessionName
	})

	if found == nil {
		return fmt.Errorf(`session "%s" not found`, sessionName)
	}

	cmd, err := NewCommand(c, "kill-session", "-t", found.Target())
	if err != nil {
		return err
	}

	return cmd.Exec()
}

// FindSessionByName returns the session with the given name if it exists.
func (c Client) FindSessionByName(sessionName string) (*Session, error) {
	sessions, _ := c.Sessions()

	found := sessions.Find(func(i int, s *Session) bool {
		return s.Name == sessionName
	})

	if found != nil {
		return found, nil
	}

	return nil, fmt.Errorf(`session "%s" not found`, sessionName)
}

// SessionExists returns true if a session with the given name exists.
func (c Client) SessionExists(sessionName string) bool {
	sessions, _ := c.Sessions()

	return sessions.Find(func(i int, s *Session) bool {
		return s.Name == sessionName
	}) != nil
}
