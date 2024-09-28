package models

import (
	"os"

	"github.com/hashicorp/hcl/v2"
	"github.com/wilhelm-murdoch/go-collection"
	"github.com/zclconf/go-cty/cty"
)

const DefaultGlazeSesssionName = "default"

type Session struct {
	Name              string
	StartingDirectory string
	Options           map[string]string
	Envs              map[string]string
	Windows           collection.Collection[*Window]
}

func (s *Session) Decode(value cty.Value) hcl.Diagnostics {
	var diags hcl.Diagnostics

	s.Name = DefaultGlazeSesssionName
	if !value.GetAttr("name").IsNull() {
		s.Name = value.GetAttr("name").AsString()
	}

	if !value.GetAttr("starting_directory").IsNull() {
		s.StartingDirectory = value.GetAttr("starting_directory").AsString()
	} else {
		if pwd, err := os.Getwd(); err == nil {
			s.StartingDirectory = pwd
		}
	}

	if !value.GetAttr("envs").IsNull() {
		s.Envs = make(map[string]string)
		for name, value := range value.GetAttr("envs").AsValueMap() {
			s.Envs[name] = value.AsString()
		}
	}

	if !value.GetAttr("options").IsNull() {
		s.Options = make(map[string]string)
		for name, value := range value.GetAttr("options").AsValueMap() {
			s.Options[name] = value.AsString()
		}
	}

	if !value.GetAttr("windows").IsNull() {
		if value.GetAttr("windows").CanIterateElements() {
			it := value.GetAttr("windows").ElementIterator()

			for it.Next() {
				_, value := it.Element()

				window := new(Window)

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
