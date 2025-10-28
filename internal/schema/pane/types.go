package pane

type Value string
type Name string
type Directory string
type Size string
type Envs map[Name]Value
type Hooks map[Name]Value
type Options map[Name]Value
type Command string
type Focus bool

type Pane struct {
	Name              Name
	StartingDirectory Directory
	Size              Size
	Envs              Envs
	Hooks             Hooks
	Options           Options
	Commands          []Command
	Focus             Focus
}
