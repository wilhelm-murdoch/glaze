package models

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/zclconf/go-cty/cty"
	"github.com/zclconf/go-cty/cty/gocty"
)

type Pane struct {
	Name     string
	Commands []string
	Focus    bool
}

func (p *Pane) Decode(value cty.Value) hcl.Diagnostics {
	var diags hcl.Diagnostics

	if !value.GetAttr("name").IsNull() {
		p.Name = value.GetAttr("name").AsString()
	}

	if !value.GetAttr("focus").IsNull() {
		gocty.FromCtyValue(value.GetAttr("focus"), &p.Focus)
	}

	if !value.GetAttr("commands").IsNull() {
		if value.GetAttr("commands").CanIterateElements() {
			cit := value.GetAttr("commands").ElementIterator()
			for cit.Next() {
				_, c := cit.Element()
				if c.Type().FriendlyName() == "string" {
					p.Commands = append(p.Commands, c.AsString())
				}
			}
		}
	}

	return diags
}
