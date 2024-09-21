package glaze

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/urfave/cli/v2"
	"github.com/zclconf/go-cty/cty"
	"github.com/zclconf/go-cty/cty/function"
	"github.com/zclconf/go-cty/cty/function/stdlib"
)

type Command interface {
	Common
	Run(*cli.Context) error
}

type Common struct{}

func (c *Common) collectVariables() (map[string]cty.Value, error) {
	// env variables with the GLAZE_ENV_ prefix
	// --var flag values
	// default values
	return nil, nil
}

func (c *Common) buildEvalContext(variables map[string]string) (*hcl.EvalContext, error) {
	ctx := &hcl.EvalContext{
		Variables: map[string]cty.Value{
			"foo": cty.StringVal("bar"),
		},
		Functions: map[string]function.Function{
			"replace":      stdlib.ReplaceFunc,
			"regexreplace": stdlib.RegexReplaceFunc,
			"upper":        stdlib.UpperFunc,
			"lower":        stdlib.LowerFunc,
			"reverse":      stdlib.ReverseFunc,
			"len":          stdlib.LengthFunc,
			"substr":       stdlib.SubstrFunc,
			"join":         stdlib.JoinFunc,
			"title":        stdlib.TitleFunc,
			"trim":         stdlib.TrimFunc,
			"trimspace":    stdlib.TrimSpaceFunc,
			"trimsuffix":   stdlib.TrimSuffixFunc,
			"trimprefix":   stdlib.TrimPrefixFunc,
			"chomp":        stdlib.ChompFunc,
		},
	}

	return ctx, nil
}
