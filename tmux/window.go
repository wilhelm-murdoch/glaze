package tmux

import (
	"strconv"
	"strings"
)

type Window struct {
	Id       int
	Index    int
	Name     string
	Layout   string
	IsActive bool
}

func (w Window) Split(name, splitType, StartingDirectory string) (Pane, error) {
	var pane Pane

	format := []string{
		"#{pane_id}",
		"#{pane_index}",
		"#{pane_title}",
		"#{pane_active}",
	}

	splitTypeFlag := "-h"
	if splitType == SplitVertical {
		splitTypeFlag = "-v"
	}

	output, err := ExecWithOutput("splitw", splitTypeFlag, "-Pd", "-t", w.Name, "-c", StartingDirectory, "-F", strings.Join(format, ";"))
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

	if err = Exec("selectp", "-T", name, "-t", parts[0]); err != nil {
		return pane, err
	}

	return Pane{
		Id:                id,
		Index:             index,
		Name:              name,
		StartingDirectory: StartingDirectory,
		IsActive:          parts[3] == "1",
	}, nil
}

func (w Window) Kill() error {
	return Exec("kill-window", "-t", w.Name)
}

func (w Window) SelectLayout(layout string) error {
	return Exec("selectl", "-t", w.Name, layout)
}
