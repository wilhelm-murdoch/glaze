package window

import (
	"github.com/wilhelm-murdoch/go-collection"

	"github.com/wilhelm-murdoch/glazier/internal/schema/pane"
	"github.com/wilhelm-murdoch/glazier/internal/tmux/enums"
)

type (
	Name      string         // Name represents the name of a tmux window.
	Value     string         // Value represents a generic string value used in various window configurations.
	Directory string         // Directory represents a file system path, typically for a window's starting directory.
	Envs      map[Name]Value // Envs is a map of environment variable names to their corresponding values for a window.
	Options   map[Name]Value // Options is a map of window-specific options to their values.
	Focus     bool           // Focus indicates whether a window should be focused.
)

// Window represents the configuration for a single tmux window.
type Window struct {
	Name              Name
	StartingDirectory Directory
	Envs              Envs
	Options           Options
	Panes             collection.Collection[*pane.Pane]
	Layout            enums.Layout
	Focus             Focus
}
