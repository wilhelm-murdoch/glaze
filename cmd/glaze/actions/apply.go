package actions

import (
	"fmt"

	"github.com/wilhelm-murdoch/glaze"
	"github.com/zclconf/go-cty/cty/gocty"
)

func ActionApply(profilePath string) error {
	p := glaze.NewParser()
	p.Open(profilePath)

	decoded := p.Decode(glaze.PrimaryGlazeSpec)

	if p.HasErrors() {
		p.WriteDiags()
		return nil
	}

	var sessions []glaze.Session

	sit := decoded.ElementIterator()
	for sit.Next() {
		_, e := sit.Element()

		var name string
		if !e.GetAttr("name").IsNull() {
			name = e.GetAttr("name").AsString()
		}

		session := glaze.Session{
			Name: name,
		}

		wit := e.GetAttr("windows").ElementIterator()
		for wit.Next() {
			_, e := wit.Element()

			var name string
			if !e.GetAttr("name").IsNull() {
				name = e.GetAttr("name").AsString()
			}

			var layout string
			if !e.GetAttr("layout").IsNull() {
				layout = e.GetAttr("layout").AsString()
			}

			var focus bool
			if !e.GetAttr("focus").IsNull() {
				gocty.FromCtyValue(e.GetAttr("focus"), &focus)
			}

			window := glaze.Window{
				Name:   name,
				Layout: layout,
				Focus:  focus,
			}

			var options []glaze.WindowOption
			if !e.GetAttr("options").IsNull() {
				if e.GetAttr("options").CanIterateElements() {
					vit := e.GetAttr("options").ElementIterator()
					for vit.Next() {
						option, value := vit.Element()
						if value.Type().FriendlyName() == "string" {
							options = append(options, glaze.NewWindowOption(option.AsString(), value.AsString()))
						}
					}
				}
			}

			window.Options = options

			pit := e.GetAttr("panes").ElementIterator()
			for pit.Next() {
				_, e := pit.Element()

				var name string
				if !e.GetAttr("name").IsNull() {
					name = e.GetAttr("name").AsString()
				}

				var focus bool
				if !e.GetAttr("focus").IsNull() {
					gocty.FromCtyValue(e.GetAttr("focus"), &focus)
				}

				var commands []string
				if !e.GetAttr("commands").IsNull() {
					if e.GetAttr("commands").CanIterateElements() {
						vit := e.GetAttr("commands").ElementIterator()
						for vit.Next() {
							_, v := vit.Element()
							if v.Type().FriendlyName() == "string" {
								commands = append(commands, v.AsString())
							}
						}
					}
				}

				pane := glaze.Pane{
					Name:     name,
					Focus:    focus,
					Commands: commands,
				}

				window.Panes = append(window.Panes, pane)

			}

			session.Windows = append(session.Windows, window)
		}

		sessions = append(sessions, session)
	}

	for _, session := range sessions {
		fmt.Println(session)
		for _, window := range session.Windows {
			fmt.Println(window)
			for _, option := range window.Options {
				fmt.Println(option)
			}
			for _, pane := range window.Panes {
				fmt.Println(pane)
			}
		}
	}

	return nil
}
