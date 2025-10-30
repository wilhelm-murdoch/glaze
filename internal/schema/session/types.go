package session

import (
	"github.com/wilhelm-murdoch/go-collection"

	"github.com/wilhelm-murdoch/glaze/internal/schema/window"
)

type (
	Name      string         // Name represents the name of a tmux session.
	Value     string         // Value represents a generic string value used in various session configurations.
	Directory string         // Directory represents a file system path, typically for a session's starting directory.
	Envs      map[Name]Value // Envs is a map of environment variable names to their corresponding values for a session.
	Commands  []string       // Commands is a slice of commands to be executed within a session.
)

type Session struct {
	Name              Name
	StartingDirectory Directory
	Envs              Envs
	Windows           collection.Collection[*window.Window]
	Commands          Commands
}
