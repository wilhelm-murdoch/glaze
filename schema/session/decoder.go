package session

import (
	"os"

	"github.com/hashicorp/hcl/v2"
	"github.com/wilhelm-murdoch/glaze/schema/window"
	"github.com/wilhelm-murdoch/go-collection"
	"github.com/zclconf/go-cty/cty"
)

const DefaultGlazeSesssionName = "default"

type Session struct {
	Name              Name
	StartingDirectory Directory
	Envs              Envs
	Windows           collection.Collection[*window.Window]
	Commands          Commands
}

func (s *Session) Decode(value cty.Value) hcl.Diagnostics {
	var diags hcl.Diagnostics

	s.Name = DefaultGlazeSesssionName
	if !value.GetAttr("name").IsNull() {
		s.Name = Name(value.GetAttr("name").AsString())
	}

	if !value.GetAttr("starting_directory").IsNull() {
		s.StartingDirectory = Directory(value.GetAttr("starting_directory").AsString())
	} else {
		if pwd, err := os.Getwd(); err == nil {
			s.StartingDirectory = Directory(pwd)
		}
	}

	if !value.GetAttr("envs").IsNull() {
		s.Envs = make(Envs)
		for name, value := range value.GetAttr("envs").AsValueMap() {
			s.Envs[Name(name)] = Value(value.AsString())
		}
	}

	if !value.GetAttr("commands").IsNull() {
		s.Commands = make(Commands, len(value.GetAttr("commands").AsValueSlice()))
		for _, command := range value.GetAttr("commands").AsValueSlice() {
			s.Commands = append(s.Commands, command.AsString())
		}
	}

	if !value.GetAttr("windows").IsNull() {
		if value.GetAttr("windows").CanIterateElements() {
			it := value.GetAttr("windows").ElementIterator()

			for it.Next() {
				_, value := it.Element()

				window := new(window.Window)
				if diags = window.Decode(value); diags.HasErrors() {
					diags = diags.Extend(diags)
					continue
				}

				s.Windows.Push(window)
			}
		}
	}

	return diags
}
