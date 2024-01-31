package models

import (
	"os"

	"github.com/hashicorp/hcl/v2"
	"github.com/wilhelm-murdoch/glaze/tmux/enums"
	"github.com/wilhelm-murdoch/go-collection"
	"github.com/zclconf/go-cty/cty"
	"github.com/zclconf/go-cty/cty/gocty"
)

type Window struct {
	Name              string
	Layout            enums.Layout
	Focus             bool
	StartingDirectory string
	Envs              map[string]string
	Panes             collection.Collection[*Pane]
}

func (w *Window) Decode(value cty.Value) hcl.Diagnostics {
	var diags hcl.Diagnostics

	if !value.GetAttr("name").IsNull() {
		w.Name = value.GetAttr("name").AsString()
	}

	if !value.GetAttr("layout").IsNull() {
		w.Layout = enums.LayoutFromString(value.GetAttr("layout").AsString())
	} else {
		w.Layout = enums.LayoutTiled
	}

	if !value.GetAttr("focus").IsNull() {
		gocty.FromCtyValue(value.GetAttr("focus"), &w.Focus)
	}

	if !value.GetAttr("starting_directory").IsNull() {
		w.StartingDirectory = value.GetAttr("starting_directory").AsString()
	} else {
		if pwd, err := os.Getwd(); err == nil {
			w.StartingDirectory = pwd
		}
	}

	if !value.GetAttr("envs").IsNull() {
		for name, value := range value.GetAttr("envs").AsValueMap() {
			w.Envs[name] = value.AsString()
		}
	}

	if !value.GetAttr("panes").IsNull() {
		if value.GetAttr("panes").CanIterateElements() {
			it := value.GetAttr("panes").ElementIterator()

			for it.Next() {
				_, value := it.Element()

				pane := new(Pane)
				if diags = pane.Decode(value); diags.HasErrors() {
					diags = diags.Extend(diags)
					continue
				}

				for name, value := range w.Envs {
					pane.Envs[name] = value
				}

				w.Panes.Push(pane)
			}
		}
	}

	return diags
}
