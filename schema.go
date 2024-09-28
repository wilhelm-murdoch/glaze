package glaze

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/wilhelm-murdoch/glaze/tmux/enums"
	"github.com/zclconf/go-cty/cty"
)

func validateOptions[OT enums.OptionTyper[OT]](value cty.Value) hcl.Diagnostics {
	var out hcl.Diagnostics

	if !value.IsNull() {
		for option, value := range value.AsValueMap() {
			var optionType OT

			if !optionType.IsKnown(option) {
				out = out.Append(&hcl.Diagnostic{
					Severity: hcl.DiagError,
					Summary:  "Invalid session option specified",
					Detail:   fmt.Sprintf(`The session option of "%s" does not exist.`, option),
				})
			}

			validator, ok := optionType.GetValidator(option)
			if !ok {
				out = out.Append(&hcl.Diagnostic{
					Severity: hcl.DiagError,
					Summary:  "Invalid validator specified",
					Detail:   fmt.Sprintf(`The session option "%s" does not have a defined validator.`, option),
				})

				continue
			}

			ok, choices := validator(value.AsString())
			if !ok {
				if len(choices) > 0 {
					out = out.Extend(ContainsDiagnostic(fmt.Sprintf(`session option "%s"`, option), value, choices))
					continue
				}

				out = out.Append(&hcl.Diagnostic{
					Severity: hcl.DiagError,
					Summary:  "Invalid session option value specified",
					Detail:   fmt.Sprintf(`The value "%s" for session option "%s" is not valid.`, value.AsString(), option),
				})
			}
		}
	}

	return out
}

var (
	envsSpec = &hcldec.AttrSpec{
		Name: "env",
		Type: cty.Map(cty.String),
	}

	hooksSpec = &hcldec.AttrSpec{
		Name: "hooks",
		Type: cty.Map(cty.String),
	}

	PrimaryGlazeSpec = &hcldec.BlockListSpec{
		TypeName: "session",
		MinItems: 1,
		MaxItems: 1,
		Nested: &hcldec.ObjectSpec{
			"name": &hcldec.AttrSpec{
				Name: "name",
				Type: cty.String,
			},
			"envs": envsSpec,
			"options": &hcldec.ValidateSpec{
				Wrapped: &hcldec.AttrSpec{
					Name: "options",
					Type: cty.Map(cty.String),
				},
				Func: func(value cty.Value) hcl.Diagnostics {
					return validateOptions[enums.OptionsSession](value)
				},
			},
			"hooks": hooksSpec,
			"starting_directory": &hcldec.ValidateSpec{
				Wrapped: &hcldec.AttrSpec{
					Name: "starting_directory",
					Type: cty.String,
				},
				Func: func(value cty.Value) hcl.Diagnostics {
					return DirectoryDiagnostic("starting directory", value)
				},
			},
			"windows": &hcldec.BlockListSpec{
				TypeName: "window",
				MinItems: 1,
				Nested: &hcldec.ObjectSpec{
					"name": &hcldec.AttrSpec{
						Name: "name",
						Type: cty.String,
					},
					"envs": envsSpec,
					"options": &hcldec.ValidateSpec{
						Wrapped: &hcldec.AttrSpec{
							Name: "options",
							Type: cty.Map(cty.String),
						},
						Func: func(value cty.Value) hcl.Diagnostics {
							return validateOptions[enums.OptionsWindow](value)
						},
					},
					"hooks": hooksSpec,
					"focus": &hcldec.AttrSpec{
						Name: "focus",
						Type: cty.Bool,
					},
					"starting_directory": &hcldec.ValidateSpec{
						Wrapped: &hcldec.AttrSpec{
							Name: "starting_directory",
							Type: cty.String,
						},
						Func: func(value cty.Value) hcl.Diagnostics {
							return DirectoryDiagnostic("starting directory", value)
						},
					},
					"layout": &hcldec.ValidateSpec{
						Wrapped: &hcldec.AttrSpec{
							Name: "layout",
							Type: cty.String,
						},
						Func: func(value cty.Value) hcl.Diagnostics {
							return ContainsDiagnostic("layout", value, enums.LayoutList)
						},
					},
					"panes": &hcldec.BlockListSpec{
						TypeName: "pane",
						MinItems: 1,
						Nested: &hcldec.ValidateSpec{
							Wrapped: &hcldec.ObjectSpec{
								"name": &hcldec.AttrSpec{
									Name: "name",
									Type: cty.String,
								},
								"envs": envsSpec,
								"options": &hcldec.ValidateSpec{
									Wrapped: &hcldec.AttrSpec{
										Name: "options",
										Type: cty.Map(cty.String),
									},
									Func: func(value cty.Value) hcl.Diagnostics {
										return validateOptions[enums.OptionsPane](value)
									},
								},
								"hooks": hooksSpec,
								"focus": &hcldec.AttrSpec{
									Name: "focus",
									Type: cty.Bool,
								},
								"size": &hcldec.ValidateSpec{
									Wrapped: &hcldec.AttrSpec{
										Name: "size",
										Type: cty.String,
									},
									Func: func(value cty.Value) hcl.Diagnostics {
										var diags hcl.Diagnostics

										if !value.IsNull() {
											input := value.AsString()

											matched := regexp.MustCompile(`^(\\d+|\\d+%)$`).MatchString(input)

											if input[len(input)-1] == '%' {
												input = input[:len(input)-1]
											}

											_, err := strconv.Atoi(input)

											if err != nil || !matched {
												diags = diags.Append(&hcl.Diagnostic{
													Severity: hcl.DiagError,
													Summary:  `Invalid size specified`,
													Detail:   `The size value must be either an integer or a percentage.`,
												})
											}
										}

										return diags
									},
								},
								"starting_directory": &hcldec.ValidateSpec{
									Wrapped: &hcldec.AttrSpec{
										Name: "starting_directory",
										Type: cty.String,
									},
									Func: func(value cty.Value) hcl.Diagnostics {
										return DirectoryDiagnostic("starting directory", value)
									},
								},
								"commands": &hcldec.AttrSpec{
									Name: "commands",
									Type: cty.List(cty.String),
								},
							},
							Func: func(value cty.Value) hcl.Diagnostics {
								var diags hcl.Diagnostics

								// Placeholder for future schema validations ...

								return diags
							},
						},
					},
				},
			},
		},
	}
)
