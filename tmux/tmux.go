package tmux

import (
	"os"
	"os/exec"
	"strconv"
	"strings"
)

const (
	LayoutEventHorizontal = "even-horizontal"
	LayoutEvenVertical    = "even-vertical"
	LayoutMainHorizontal  = "main-horizontal"
	LayoutMainVertical    = "main-vertical"
	LayoutTiled           = "tiled"
	LayoutUnknown         = "unknown"
)

var (
	LayoutList = []string{
		LayoutEventHorizontal,
		LayoutEvenVertical,
		LayoutMainHorizontal,
		LayoutMainVertical,
		LayoutTiled,
	}
)

func LayoutExists(value string) bool {
	for _, v := range LayoutList {
		if string(v) == value {
			return true
		}
	}

	return false
}

type Tmux struct{}

func (t Tmux) IsInstalled() (string, bool) {
	path, err := exec.LookPath("tmux")
	return path, err == nil
}

func (t Tmux) IsInsideTmux() bool {
	if os.Getenv("TMUX") != "" {
		return true
	} else {
		return false
	}
}

func (t Tmux) execute(cmd *exec.Cmd) (string, error) {
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}

	return strings.TrimSuffix(string(output), "\n"), nil
}

func (t Tmux) SendKeys(pane Pane, keys string) (string, error) {
	return t.execute(exec.Command("tmux", "send-keys", "-t", pane.Name, keys, "Enter"))
}

func (tmux Tmux) Sessions() ([]Session, error) {
	var sessions []Session
	output, err := tmux.execute(exec.Command("tmux", "list-sessions"))
	if err != nil {
		return sessions, err
	}

	for _, line := range strings.Split(output, "\n") {
		parts := strings.SplitN(line, ":", 2)
		if len(parts) == 2 {
			sessions = append(sessions, Session{
				Name: strings.TrimSpace(parts[0]),
			})
		}
	}

	return sessions, err
}

func (t Tmux) Windows(session Session) ([]Window, error) {
	var windows []Window

	parts := []string{
		"#{window_id}",
		"#{window_name}",
		"#{window_layout}",
		"#{pane_current_path}",
	}

	output, err := t.execute(exec.Command("tmux", "list-windows", "-F", strings.Join(parts, ";"), "-t", session.Name))
	if err != nil {
		return windows, err
	}

	for _, window := range strings.Split(output, "\n") {
		parts = strings.Split(window, ";")

		id, err := strconv.Atoi(strings.Replace(parts[0], "@", "", -1))
		if err != nil {
			return windows, err
		}

		windows = append(windows, Window{
			Id:   id,
			Name: parts[1],
		})
	}

	return windows, nil
}

func (t Tmux) Panes(window Window) ([]Pane, error) {
	var panes []Pane

	parts := []string{
		"#{pane_id}",
		"#{pane_title}",
		"#{pane_current_path}",
	}

	output, err := t.execute(exec.Command("tmux", "list-panes", "-F", strings.Join(parts, ";"), "-t", window.Name))
	if err != nil {
		return panes, err
	}

	for _, p := range strings.Split(output, "\n") {
		parts := strings.Split(p, ";")

		id, err := strconv.Atoi(strings.Replace(parts[0], "@", "", -1))
		if err != nil {
			return panes, err
		}

		panes = append(panes, Pane{
			Id:               id,
			Name:             parts[1],
			CurrentDirectory: parts[2],
		})
	}

	return panes, nil
}

func (t Tmux) NewSession(name string, root string, windowName string) (string, error) {
	return t.execute(exec.Command("tmux", "new", "-Pd", "-s", name, "-n", windowName, "-c", root))
}

func (t Tmux) SessionExists(name string) bool {
	output, err := t.execute(exec.Command("tmux", "has-session", "-t", name))
	return output == "" && err == nil
}

func (t Tmux) AttachToSession(name string) (string, error) {
	return t.execute(exec.Command("tmux", "attach", "-d", "-t", name))
}

func (t Tmux) SelectLayout(window Window, layout string) (string, error) {
	return t.execute(exec.Command("tmux", "select-layout", "-t", window.Name, string(layout)))
}

func (t Tmux) SetEnv(pane Pane, key string, value string) (string, error) {
	return t.execute(exec.Command("tmux", "setenv", "-t", pane.Name, key, value))
}
