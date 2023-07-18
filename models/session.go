package models

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/zclconf/go-cty/cty"
)

type Session struct {
	Name    string
	Windows []Window
}

func (s *Session) Decode(value cty.Value) hcl.Diagnostics {
	var diags hcl.Diagnostics

	if !value.GetAttr("name").IsNull() {
		s.Name = value.GetAttr("name").AsString()
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

				s.Windows = append(s.Windows, *window)
			}
		}
	}

	return diags
}
