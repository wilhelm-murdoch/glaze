package tmux

import "fmt"

type PaneId int

func (id PaneId) String() string {
	return fmt.Sprintf("%%%d", int(id))
}

type SessionId int

func (id SessionId) String() string {
	return fmt.Sprintf("$%d", int(id))
}

type WindowId int

func (id WindowId) String() string {
	return fmt.Sprintf("@%d", int(id))
}
