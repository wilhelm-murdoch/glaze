package models

import (
	"os"

	"github.com/hashicorp/hcl/v2"
	"github.com/zclconf/go-cty/cty"
	"github.com/zclconf/go-cty/cty/gocty"
)

const DefaultGlazePaneName = "default"

type Pane struct {
	Name              string
	StartingDirectory string
	Size              string
	Envs              map[string]string
	Hooks             map[string]string
	Commands          []string
	Focus             bool
}

func (p *Pane) Decode(value cty.Value) hcl.Diagnostics {
	var diags hcl.Diagnostics

	p.Name = DefaultGlazePaneName
	if !value.GetAttr("name").IsNull() {
		p.Name = value.GetAttr("name").AsString()
	}

	if !value.GetAttr("focus").IsNull() {
		gocty.FromCtyValue(value.GetAttr("focus"), &p.Focus)
	}

	if !value.GetAttr("starting_directory").IsNull() {
		p.StartingDirectory = value.GetAttr("starting_directory").AsString()
	} else {
		if pwd, err := os.Getwd(); err == nil {
			p.StartingDirectory = pwd
		}
	}

	if !value.GetAttr("size").IsNull() {
		p.Size = value.GetAttr("size").AsString()
	}

	if !value.GetAttr("envs").IsNull() {
		p.Envs = make(map[string]string)
		for name, value := range value.GetAttr("envs").AsValueMap() {
			p.Envs[name] = value.AsString()
		}
	}

	if !value.GetAttr("hooks").IsNull() {
		p.Hooks = make(map[string]string)
		for name, value := range value.GetAttr("hooks").AsValueMap() {
			p.Hooks[name] = value.AsString()
		}
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
