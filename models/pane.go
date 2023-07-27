package models

import (
	"os"

	"github.com/hashicorp/hcl/v2"
	"github.com/wilhelm-murdoch/glaze/tmux"
	"github.com/zclconf/go-cty/cty"
	"github.com/zclconf/go-cty/cty/gocty"
)

type Pane struct {
	Name              string
	StartingDirectory string
	Commands          []string
	IsActive          bool
	Split             string
	Size              string
	Placement         string
	Full              string
}

func (p *Pane) Decode(value cty.Value) hcl.Diagnostics {
	var diags hcl.Diagnostics

	if !value.GetAttr("name").IsNull() {
		p.Name = value.GetAttr("name").AsString()
	}

	if !value.GetAttr("focus").IsNull() {
		gocty.FromCtyValue(value.GetAttr("focus"), &p.IsActive)
	}

	if !value.GetAttr("starting_directory").IsNull() {
		p.StartingDirectory = value.GetAttr("starting_directory").AsString()
	} else {
		if pwd, err := os.Getwd(); err == nil {
			p.StartingDirectory = pwd
		}
	}

	if !value.GetAttr("split").IsNull() {
		p.Split = value.GetAttr("split").AsString()
	} else {
		p.Split = tmux.SplitVertical
	}

	if !value.GetAttr("size").IsNull() {
		p.Size = value.GetAttr("size").AsString()
	}

	if !value.GetAttr("full").IsNull() {
		gocty.FromCtyValue(value.GetAttr("full"), &p.Full)
	}

	if !value.GetAttr("placement").IsNull() {
		p.Placement = value.GetAttr("placement").AsString()
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
