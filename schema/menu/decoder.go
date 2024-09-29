package menu

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/zclconf/go-cty/cty"
	"github.com/zclconf/go-cty/cty/gocty"
)

const DefaultGlazeMenuName = "menu"

func (m *Menu) Decode(value cty.Value) hcl.Diagnostics {
	var diags hcl.Diagnostics

	m.Name = DefaultGlazeMenuName
	if !value.GetAttr("name").IsNull() {
		m.Name = Name(value.GetAttr("name").AsString())
	}

	if !value.GetAttr("shell-script").IsNull() {
		m.ShellScript = ShellScript(value.GetAttr("shell-script").AsString())
	}

	if !value.GetAttr("bind").IsNull() {
		m.Bind = Bind(value.GetAttr("bind").AsString())
	}

	if !value.GetAttr("items").IsNull() {
		m.Items = []Item{}

		it := value.GetAttr("items").ElementIterator()

		for it.Next() {
			_, value := it.Element()

			var name Name
			if !value.GetAttr("name").IsNull() {
				name = Name(value.GetAttr("name").AsString())
			}

			var bind Bind
			if !value.GetAttr("bind").IsNull() {
				bind = Bind(value.GetAttr("bind").AsString())
			}

			var command Command
			if !value.GetAttr("command").IsNull() {
				command = Command(value.GetAttr("command").AsString())
			}

			var disabled Disabled
			if !value.GetAttr("disabled").IsNull() {
				gocty.FromCtyValue(value.GetAttr("disabled"), &disabled)
			}

			m.Items = append(m.Items, Item{
				Name:     name,
				Bind:     bind,
				Command:  command,
				Disabled: disabled,
			})
		}
	}

	return diags
}
