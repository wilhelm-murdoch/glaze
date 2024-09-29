package pane

import (
	"os"

	"github.com/hashicorp/hcl/v2"
	"github.com/zclconf/go-cty/cty"
	"github.com/zclconf/go-cty/cty/gocty"
)

const DefaultGlazePaneName = "default"

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

func (p *Pane) Decode(value cty.Value) hcl.Diagnostics {
	var diags hcl.Diagnostics

	p.Name = DefaultGlazePaneName
	if !value.GetAttr("name").IsNull() {
		p.Name = Name(value.GetAttr("name").AsString())
	}

	if !value.GetAttr("focus").IsNull() {
		gocty.FromCtyValue(value.GetAttr("focus"), &p.Focus)
	}

	if !value.GetAttr("starting_directory").IsNull() {
		p.StartingDirectory = Directory(value.GetAttr("starting_directory").AsString())
	} else {
		if pwd, err := os.Getwd(); err == nil {
			p.StartingDirectory = Directory(pwd)
		}
	}

	if !value.GetAttr("size").IsNull() {
		p.Size = Size(value.GetAttr("size").AsString())
	}

	if !value.GetAttr("envs").IsNull() {
		p.Envs = make(Envs)
		for name, value := range value.GetAttr("envs").AsValueMap() {
			p.Envs[Name(name)] = Value(value.AsString())
		}
	}

	if !value.GetAttr("hooks").IsNull() {
		p.Hooks = make(Hooks)
		for name, value := range value.GetAttr("hooks").AsValueMap() {
			p.Hooks[Name(name)] = Value(value.AsString())
		}
	}

	if !value.GetAttr("options").IsNull() {
		p.Options = make(Options)
		for name, value := range value.GetAttr("options").AsValueMap() {
			p.Options[Name(name)] = Value(value.AsString())
		}
	}

	if !value.GetAttr("commands").IsNull() {
		if value.GetAttr("commands").CanIterateElements() {
			cit := value.GetAttr("commands").ElementIterator()
			for cit.Next() {
				_, c := cit.Element()
				if c.Type().FriendlyName() == "string" {
					p.Commands = append(p.Commands, Command(c.AsString()))
				}
			}
		}
	}

	return diags
}
