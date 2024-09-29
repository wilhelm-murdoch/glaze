package window

import (
	"os"

	"github.com/hashicorp/hcl/v2"
	"github.com/wilhelm-murdoch/glaze/schema/pane"
	"github.com/wilhelm-murdoch/glaze/tmux/enums"
	"github.com/wilhelm-murdoch/go-collection"
	"github.com/zclconf/go-cty/cty"
	"github.com/zclconf/go-cty/cty/gocty"
)

const DefaultGlazeWindowName = "default"

type Window struct {
	Name              Name
	StartingDirectory Directory
	Envs              Envs
	Options           Options
	Panes             collection.Collection[*pane.Pane]
	Layout            enums.Layout
	Focus             Focus
}

func (w *Window) Decode(value cty.Value) hcl.Diagnostics {
	var diags hcl.Diagnostics

	w.Name = "default"
	if !value.GetAttr("name").IsNull() {
		w.Name = Name(value.GetAttr("name").AsString())
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
		w.StartingDirectory = Directory(value.GetAttr("starting_directory").AsString())
	} else {
		if pwd, err := os.Getwd(); err == nil {
			w.StartingDirectory = Directory(pwd)
		}
	}

	if !value.GetAttr("envs").IsNull() {
		w.Envs = make(Envs)
		for name, value := range value.GetAttr("envs").AsValueMap() {
			w.Envs[Name(name)] = Value(value.AsString())
		}
	}

	if !value.GetAttr("options").IsNull() {
		w.Options = make(Options)
		for name, value := range value.GetAttr("options").AsValueMap() {
			w.Options[Name(name)] = Value(value.AsString())
		}
	}

	if !value.GetAttr("panes").IsNull() {
		if value.GetAttr("panes").CanIterateElements() {
			it := value.GetAttr("panes").ElementIterator()

			for it.Next() {
				_, value := it.Element()

				pane := new(pane.Pane)

				if diags = pane.Decode(value); diags.HasErrors() {
					diags = diags.Extend(diags)
					continue
				}

				w.Panes.Push(pane)
			}
		}
	}

	return diags
}
