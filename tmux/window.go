package tmux

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/wilhelm-murdoch/glaze/tmux/enums"
)

// Window represents a tmux window.
type Window struct {
	Id       int
	Index    int
	Name     string
	Layout   enums.Layout
	IsActive bool
	Session  *Session
}

// Target returns the target window by its composite id of session name
// and window id.
func (w Window) Target() string {
	return fmt.Sprintf(`%s:@%d`, w.Session.Name, w.Id)
}

// Split splits the current window into two panes.
func (w *Window) Split(splitType enums.Split, placement enums.Placement, full enums.Full, name, startingDirectory, size string) (Pane, error) {
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
		"-t", w.Name,
		"-c", startingDirectory,
		"-F", strings.Join(format, ";"),
	}

	splitTypeFlag := "-h"
	if splitType == enums.SplitVertical {
		splitTypeFlag = "-v"
	}

	args = append(args, splitTypeFlag)

	if size != "" {
		args = append(args, "-l", size)
	}

	if placement != enums.PlacementUnknown {
		args = append(args, "-b", fmt.Sprint(placement))
	}

	if full != enums.FullUnknown {
		args = append(args, "-f")
	}

	cmd, err := NewCommand(args...)
	if err != nil {
		return pane, err
	}

	output, err := cmd.ExecWithOutput()
	if err != nil {
		return pane, err
	}

	parts := strings.SplitN(output, ";", 4)

	id, err := strconv.Atoi(strings.Replace(parts[0], "%", "", -1))
	if err != nil {
		return pane, err
	}

	index, err := strconv.Atoi(parts[1])
	if err != nil {
		return pane, err
	}

	cmd, err = NewCommand("selectp", "-T", name, "-t", parts[0])
	if err != nil {
		return pane, err
	}

	if err = cmd.Exec(); err != nil {
		return pane, err
	}

	return Pane{
		Id:                id,
		Index:             index,
		Name:              name,
		StartingDirectory: startingDirectory,
		IsActive:          parts[3] == "1",
		Window:            w,
	}, nil
}

func (w Window) Reindex() error {
	return nil
}

func (w Window) Kill() error {
	cmd, err := NewCommand("killw", "-t", w.Target())
	if err != nil {
		return err
	}

	return cmd.Exec()
}

func (w Window) SelectLayout(layout enums.Layout) error {
	cmd, err := NewCommand("selectl", "-t", w.Target(), fmt.Sprint(layout))
	if err != nil {
		return err
	}

	return cmd.Exec()
}
