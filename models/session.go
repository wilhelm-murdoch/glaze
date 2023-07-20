package models

import (
	"os"

	"github.com/hashicorp/hcl/v2"
	"github.com/wilhelm-murdoch/go-collection"
	"github.com/zclconf/go-cty/cty"
	"github.com/zclconf/go-cty/cty/gocty"
)

type Session struct {
	Name              string
	ReattachOnStart   bool
	StartingDirectory string
	Windows           collection.Collection[*Window]
}

func (s *Session) Decode(value cty.Value) hcl.Diagnostics {
	var diags hcl.Diagnostics

	if !value.GetAttr("name").IsNull() {
		s.Name = value.GetAttr("name").AsString()
	}

	if !value.GetAttr("reattach_on_start").IsNull() {
		gocty.FromCtyValue(value.GetAttr("reattach_on_start"), &s.ReattachOnStart)
	} else {
		s.ReattachOnStart = true
	}

	if !value.GetAttr("starting_directory").IsNull() {
		s.StartingDirectory = value.GetAttr("starting_directory").AsString()
	} else {
		if pwd, err := os.Getwd(); err == nil {
			s.StartingDirectory = pwd
		}
	}

	if !value.GetAttr("windows").IsNull() {
		if value.GetAttr("windows").CanIterateElements() {
			it := value.GetAttr("windows").ElementIterator()

			for it.Next() {
				_, value := it.Element()

				window := new(Window)
				diags = window.Decode(value)
				if diags.HasErrors() {
					diags = diags.Extend(diags)
					continue
				}

				s.Windows.Push(window)
			}
		}
	}

	return diags
}
