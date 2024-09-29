package schema

import (
	"regexp"
	"strconv"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/wilhelm-murdoch/glaze"
	"github.com/wilhelm-murdoch/glaze/tmux/enums"
	"github.com/zclconf/go-cty/cty"
)

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
			"menu": &hcldec.ValidateSpec{
				Wrapped: &hcldec.ObjectSpec{
					"name": &hcldec.AttrSpec{
						Name: "name",
						Type: cty.String,
					},
					"bind": &hcldec.AttrSpec{
						Name: "bind",
						Type: cty.String,
					},
					"shell-script": &hcldec.ValidateSpec{
						Wrapped: &hcldec.AttrSpec{
							Name: "shell-script",
							Type: cty.String,
						},
						Func: func(value cty.Value) hcl.Diagnostics {
							return nil
						},
					},
					"items": &hcldec.ValidateSpec{
						Wrapped: &hcldec.AttrSpec{
							Name: "items",
							Type: cty.Map(cty.String),
						},
						Func: func(value cty.Value) hcl.Diagnostics {
							return nil
						},
					},
				},
				Func: func(value cty.Value) hcl.Diagnostics {
					return nil
				},
			},
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
					return glaze.DirectoryDiagnostic("starting directory", value)
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
							return glaze.DirectoryDiagnostic("starting directory", value)
						},
					},
					"layout": &hcldec.ValidateSpec{
						Wrapped: &hcldec.AttrSpec{
							Name: "layout",
							Type: cty.String,
						},
						Func: func(value cty.Value) hcl.Diagnostics {
							return glaze.ContainsDiagnostic("layout", value, enums.LayoutList)
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
										return glaze.DirectoryDiagnostic("starting directory", value)
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
