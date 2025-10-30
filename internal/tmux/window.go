package tmux

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/wilhelm-murdoch/glaze/internal/schema/pane"
	"github.com/wilhelm-murdoch/glaze/internal/tmux/enums"
)

type WindowId int

// String is responsible for returning the string representation of the WindowId.
func (id WindowId) String() string {
	return fmt.Sprintf("@%d", int(id))
}

// Window represents a tmux window.
type Window struct {
	Session  *Session
	Name     string
	IsActive bool
	IsFirst  bool
	Id       int
	Index    int
	Layout   enums.Layout
}

// Target returns the target window by its composite id of session name
// and window id.
func (w Window) Target() string {
	return fmt.Sprintf(`%s:%d`, w.Session.Name, w.Index)
}

// Split splits the current window into two panes.
func (w *Window) Split(
	parentId string,
	name pane.Name,
	startingDirectory pane.Directory,
) (Pane, error) {
	var pane Pane

	format := []string{
		"#{pane_id}",
		"#{pane_index}",
		"#{pane_title}",
		"#{pane_active}",
	}

	args := []string{
		"splitw",
		"-Pd",
		"-t", parentId,
		"-c", fmt.Sprint(startingDirectory),
		"-F", strings.Join(format, ";"),
	}

	cmd, err := NewCommand(w.Session.Client, args...)
	if err != nil {
		return pane, err
	}

	output, err := cmd.ExecWithOutput()
	if err != nil {
		return pane, err
	}

	parts := strings.SplitN(output, ";", len(format))

	id, err := strconv.Atoi(strings.ReplaceAll(parts[0], "%", ""))
	if err != nil {
		return pane, err
	}

	index, err := strconv.Atoi(parts[1])
	if err != nil {
		return pane, err
	}

	cmd, err = NewCommand(w.Session.Client, "selectp", "-T", fmt.Sprint(name), "-t", parts[0])
	if err != nil {
		return pane, err
	}

	if err = cmd.Exec(); err != nil {
		return pane, err
	}

	args = []string{
		"show",
		"-gw",
		"-t", w.Target(),
		"pane-base-index",
	}

	baseIndexCmd, err := NewCommand(w.Session.Client, args...)
	if err != nil {
		return pane, err
	}

	baseIndexCmdOutput, err := baseIndexCmd.ExecWithOutput()
	if err != nil {
		return pane, err
	}

	baseIndexCmdParts := strings.Split(baseIndexCmdOutput, " ")
	if len(baseIndexCmdParts) != 2 {
		return pane, errors.New("could not determine pane base index")
	}

	return Pane{
		Id:                PaneId(id),
		Index:             index,
		Name:              fmt.Sprint(name),
		StartingDirectory: fmt.Sprint(startingDirectory),
		IsActive:          parts[3] == "1",
		IsFirst:           parts[1] == baseIndexCmdParts[1],
		Window:            w,
	}, nil
}

// Kill is responsible for closing the current window.
func (w Window) Kill() error {
	cmd, err := NewCommand(w.Session.Client, "killw", "-t", w.Target())
	if err != nil {
		return err
	}

	return cmd.Exec()
}

// Select is responsible for selecting the current window.
func (w Window) Select() error {
	cmd, err := NewCommand(w.Session.Client, "selectw", "-t", w.Target())
	if err != nil {
		return err
	}

	return cmd.Exec()
}

// SelectLayout is responsible for selecting the layout for the current window.
func (w Window) SelectLayout(layout enums.Layout) error {
	cmd, err := NewCommand(w.Session.Client, "selectl", "-t", w.Target(), fmt.Sprint(layout))
	if err != nil {
		return err
	}

	return cmd.Exec()
}
