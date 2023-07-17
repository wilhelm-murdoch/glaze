package glaze

import (
	"fmt"
	"strings"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/zclconf/go-cty/cty"
)

const (
	LayoutEventHorizontal = "even-horizontal"
	LayoutEvenVertical    = "even-vertical"
	LayoutMainHorizontal  = "main-horizontal"
	LayoutMainVertical    = "main-vertical"
	LayoutTiled           = "tiled"
	LayoutUnknown         = "unknown"
)

func ListContains(list []string, value string) bool {
	for _, v := range list {
		if v == value {
			return true
		}
	}

	return false
}

func JoinWithOr(choices []string) string {
	return joinWith(choices, "or")
}

func JoinWithAnd(choices []string) string {
	return joinWith(choices, "and")
}

func joinWith(choices []string, conjunction string) string {
	length := len(choices)
	if length == 0 {
		return ""
	} else if length == 1 {
		return choices[0]
	} else if length == 2 {
		return fmt.Sprintf(`%s %s %s`, choices[0], conjunction, choices[1])
	}

	return fmt.Sprintf(`%s %s %s`, strings.Join(choices[:length-1], ", "), conjunction, choices[length-1])
}

var (
	LayoutList = []string{
		LayoutEventHorizontal,
		LayoutEvenVertical,
		LayoutMainHorizontal,
		LayoutMainVertical,
		LayoutTiled,
	}

	PrimaryGlazeSpec = &hcldec.BlockListSpec{
		TypeName: "session",
		MinItems: 1,
		Nested: &hcldec.ObjectSpec{
			"name": &hcldec.BlockLabelSpec{
				Index: 0,
				Name:  "name",
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
							if !ListContains(LayoutList, value.AsString()) {
								return hcl.Diagnostics{{
									Severity: hcl.DiagError,
									Summary:  `Invalid layout specified`,
									Detail:   fmt.Sprintf(`The layout of "%s" is not supported among: %s`, value.AsString(), JoinWithOr(LayoutList)),
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
