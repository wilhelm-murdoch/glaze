package window

import (
	"github.com/wilhelm-murdoch/go-collection"

	"github.com/wilhelm-murdoch/glaze/internal/schema/pane"
	"github.com/wilhelm-murdoch/glaze/internal/tmux/enums"
)

type (
	Name      string
	Value     string
	Directory string
	Envs      map[Name]Value
	Options   map[Name]Value
	Focus     bool
)

type Window struct {
	Name              Name
	StartingDirectory Directory
	Envs              Envs
	Options           Options
	Panes             collection.Collection[*pane.Pane]
	Layout            enums.Layout
	Focus             Focus
}
