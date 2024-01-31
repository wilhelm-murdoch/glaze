package glaze

import (
	"regexp"
	"strconv"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/wilhelm-murdoch/glaze/tmux/enums"
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
			"envs": &hcldec.AttrSpec{
				Name: "env",
				Type: cty.Map(cty.String),
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
			"windows": &hcldec.BlockListSpec{
				TypeName: "window",
				MinItems: 1,
				Nested: &hcldec.ObjectSpec{
					"name": &hcldec.BlockLabelSpec{
						Index: 0,
						Name:  "name",
					},
					"envs": &hcldec.AttrSpec{
						Name: "env",
						Type: cty.Map(cty.String),
					},
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
								"name": &hcldec.BlockLabelSpec{
									Index: 0,
									Name:  "name",
								},
								"envs": &hcldec.AttrSpec{
									Name: "env",
									Type: cty.Map(cty.String),
								},
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
								"placement": &hcldec.ValidateSpec{
									Wrapped: &hcldec.AttrSpec{
										Name: "placement",
										Type: cty.String,
									},
									Func: func(value cty.Value) hcl.Diagnostics {
										return ContainsDiagnostic("placement", value, enums.PlacementList)
									},
								},
								"full": &hcldec.ValidateSpec{
									Wrapped: &hcldec.AttrSpec{
										Name: "full",
										Type: cty.String,
									},
									Func: func(value cty.Value) hcl.Diagnostics {
										return ContainsDiagnostic("full", value, enums.FullList)
									},
								},
								"split": &hcldec.ValidateSpec{
									Wrapped: &hcldec.AttrSpec{
										Name: "split",
										Type: cty.String,
									},
									Func: func(value cty.Value) hcl.Diagnostics {
										return ContainsDiagnostic("split", value, enums.SplitList)
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

								split := value.GetAttr("split")
								full := value.GetAttr("full")
								placement := value.GetAttr("placement")
								if split.IsNull() || enums.SplitFromString(split.AsString()) == enums.SplitVertical {
									if !placement.IsNull() && enums.PlacementFromString(placement.AsString()) == enums.PlacementAbove {
										diags = diags.Append(WrongAttributeDiagnostic("placement", placement.AsString(), enums.PlacementLeftString))
									}

									if !full.IsNull() && enums.FullFromString(full.AsString()) == enums.FullHeight {
										diags = diags.Append(WrongAttributeDiagnostic("full", full.AsString(), enums.FullWidthString))
									}

									return diags
								}

								if !split.IsNull() && enums.SplitFromString(split.AsString()) == enums.SplitHorizontal {
									if !placement.IsNull() && enums.PlacementFromString(placement.AsString()) == enums.PlacementLeft {
										diags = diags.Append(WrongAttributeDiagnostic("placement", placement.AsString(), enums.PlacementAboveString))

									}

									if !full.IsNull() && enums.FullFromString(full.AsString()) == enums.FullWidth {
										diags = diags.Append(WrongAttributeDiagnostic("full", full.AsString(), enums.FullHeightString))
									}

									return diags
								}

								return nil
							},
						},
					},
				},
			},
		},
	}
)
