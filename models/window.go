package models

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/zclconf/go-cty/cty"
	"github.com/zclconf/go-cty/cty/gocty"
)

type Window struct {
	Name    string
	Layout  string
	Focus   bool
	Options []Option
	Panes   []Pane
}

func (w *Window) Decode(value cty.Value) hcl.Diagnostics {
	var diags hcl.Diagnostics

	if !value.GetAttr("name").IsNull() {
		w.Name = value.GetAttr("name").AsString()
	}

	if !value.GetAttr("layout").IsNull() {
		w.Layout = value.GetAttr("layout").AsString()
	}

	if !value.GetAttr("focus").IsNull() {
		gocty.FromCtyValue(value.GetAttr("focus"), &w.Focus)
	}

	if !value.GetAttr("panes").IsNull() {
		if value.GetAttr("panes").CanIterateElements() {
			it := value.GetAttr("panes").ElementIterator()

			for it.Next() {
				_, value := it.Element()

				pane := new(Pane)
				diags = pane.Decode(value)
				if diags.HasErrors() {
					diags = diags.Extend(diags)
					continue
				}

				w.Panes = append(w.Panes, *pane)
			}
		}
	}

	return diags
}
