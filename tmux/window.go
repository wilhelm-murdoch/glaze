package tmux

import (
	"fmt"
	"strconv"
	"strings"
)

type Window struct {
	Id       int
	Index    int
	Name     string
	Layout   string
	IsActive bool
	Session  *Session
}

func (w Window) Target() string {
	return fmt.Sprintf(`%s:@%d`, w.Session.Name, w.Id)
}

func (w *Window) Split(name, splitType, startingDirectory, size, placement, full string) (Pane, error) {
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
	if splitType == SplitVertical {
		splitTypeFlag = "-v"
	}

	args = append(args, splitTypeFlag)

	if size != "" {
		args = append(args, "-l", size)
	}

	if placement != "" {
		args = append(args, "-b", placement)
	}

	if full != "" {
		args = append(args, "-f")
	}

	fmt.Println("size:", size, "placement:", placement, "full:", full)

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

func (w Window) SelectLayout(layout string) error {
	cmd, err := NewCommand("selectl", "-t", w.Target(), layout)
	if err != nil {
		return err
	}

	return cmd.Exec()
}
