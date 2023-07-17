package glaze

import "fmt"

type Session struct {
	Name    string
	Windows []Window
}

func (s Session) String() string {
	return fmt.Sprintf(`new-session -d -s "%s"`, s.Name)
}

type WindowOption struct {
	Option string
	Value  any
}

func NewWindowOption(option string, value any) WindowOption {
	return WindowOption{
		Option: option,
		Value:  value,
	}
}

func (wo WindowOption) String() string {
	return fmt.Sprintf(`set-window-option -t "%s" %s %s`, "{name}", wo.Option, wo.Value)
}

type Window struct {
	Name    string
	Layout  string
	Focus   bool
	Options []WindowOption
	Panes   []Pane
}

func (w Window) String() string {
	return fmt.Sprintf(`new-window -t "%s" -n "%s"`, "{name}", w.Name)
}

type Pane struct {
	Name     string
	Commands []string
	Focus    bool
}

func (p Pane) String() string {
	return fmt.Sprintf(`split-window -t "%s" -n "%s"`, "{name}", p.Name)
}
