package glaze

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"strings"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/wilhelm-murdoch/glaze/tmux"
	"github.com/zclconf/go-cty/cty"
)

var (
	PrimaryGlazeSpec = &hcldec.BlockListSpec{
		TypeName: "session",
		MinItems: 1,
		MaxItems: 1,
		Nested: &hcldec.ObjectSpec{
			"name": &hcldec.BlockLabelSpec{
				Index: 0,
				Name:  "name",
			},
			"reattach_on_start": &hcldec.AttrSpec{
				Name: "reattach_on_start",
				Type: cty.Bool,
			},
			"starting_directory": &hcldec.ValidateSpec{
				Wrapped: &hcldec.AttrSpec{
					Name: "starting_directory",
					Type: cty.String,
				},
				Func: func(value cty.Value) hcl.Diagnostics {
					if !value.IsNull() {
						fileInfo, err := os.Stat(value.AsString())
						if err != nil || errors.Is(err, fs.ErrNotExist) || !fileInfo.IsDir() {
							return hcl.Diagnostics{{
								Severity: hcl.DiagError,
								Summary:  `Invalid starting directory specified`,
								Detail:   fmt.Sprintf(`The starting directory "%s" does not exist or is not a directory`, value.AsString()),
							}}
						}
					}

					return nil
				},
			},
			"windows": &hcldec.BlockListSpec{
				TypeName: "window",
				MinItems: 1,
				Nested: &hcldec.ObjectSpec{
					"name": &hcldec.BlockLabelSpec{
						Index: 0,
						Name:  "name",
					},
					"focus": &hcldec.AttrSpec{
						Name: "focus",
						Type: cty.Bool,
					},
					"options": &hcldec.AttrSpec{
						Name:     "options",
						Type:     cty.Map(cty.String),
						Required: true,
					},
					"layout": &hcldec.ValidateSpec{
						Wrapped: &hcldec.AttrSpec{
							Name: "layout",
							Type: cty.String,
						},
						Func: func(value cty.Value) hcl.Diagnostics {
							if !tmux.LayoutExists(value.AsString()) {
								return hcl.Diagnostics{{
									Severity: hcl.DiagError,
									Summary:  `Invalid layout specified`,
									Detail:   fmt.Sprintf(`The layout of "%s" is not supported among: %s`, value.AsString(), strings.Join(tmux.LayoutList, ", ")),
								}}
							}

							return nil
						},
					},
					"panes": &hcldec.BlockListSpec{
						TypeName: "pane",
						MinItems: 1,
						Nested: &hcldec.ObjectSpec{
							"name": &hcldec.BlockLabelSpec{
								Index: 0,
								Name:  "name",
							},
							"focus": &hcldec.AttrSpec{
								Name: "focus",
								Type: cty.Bool,
							},
							"commands": &hcldec.AttrSpec{
								Name: "commands",
								Type: cty.List(cty.String),
							},
						},
					},
				},
			},
		},
	}
)
