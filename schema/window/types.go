package window

import (
	"github.com/wilhelm-murdoch/glaze/schema/pane"
	"github.com/wilhelm-murdoch/glaze/tmux/enums"
	"github.com/wilhelm-murdoch/go-collection"
)

type Name string
type Value string
type Directory string
type Envs map[Name]Value
type Options map[Name]Value
type Focus bool

type Window struct {
	Name              Name
	StartingDirectory Directory
	Envs              Envs
	Options           Options
	Panes             collection.Collection[*pane.Pane]
	Layout            enums.Layout
	Focus             Focus
}
