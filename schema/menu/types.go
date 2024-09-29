package menu

import (
	"fmt"
)

type Name string
type Value string
type Bind string
type Command string
type ShellScript Command
type Disabled bool

type Items []Item

func (i Items) String() string {
	var out string
	for _, item := range i {
		out += fmt.Sprintf("%s ", item)
	}

	return out
}

type Item struct {
	Name     Name
	Bind     Bind
	Command  Command
	Disabled Disabled
}

func (i Item) String() string {
	var namePrefix string
	if i.Disabled {
		namePrefix = "-"
	}

	return fmt.Sprintf("%s%s %s \"%s\"", namePrefix, i.Name, i.Bind, i.Command)
}

type Menu struct {
	Name        Name
	Bind        Bind
	Items       Items
	ShellScript ShellScript
}

func (m Menu) CommandArgs() []string {
	return []string{
		"bind",
		fmt.Sprint(m.Bind),
		fmt.Sprintf("display-menu -T %s %s", m.Name, m.Items),
	}
}
