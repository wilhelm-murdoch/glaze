package parser

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/hashicorp/hcl/v2/hclparse"
	"github.com/zclconf/go-cty/cty"
	"github.com/zclconf/go-cty/cty/function"
	"github.com/zclconf/go-cty/cty/function/stdlib"

	"github.com/wilhelm-murdoch/glaze/internal/schema/session"
	"github.com/wilhelm-murdoch/glaze/pkg/files"
)

type Parser struct {
	File   *hcl.File
	parser *hclparse.Parser
}

func NewParser(path string) (*Parser, hcl.Diagnostics) {
	parser := hclparse.NewParser()
	file, diags := parser.ParseHCLFile(path)

	if diags.HasErrors() {
		return nil, diags
	}

	return &Parser{
		parser: parser,
		File:   file,
	}, nil
}

func (p *Parser) Decode(
	spec hcldec.Spec,
	ctx *hcl.EvalContext,
) (*session.Session, hcl.Diagnostics) {
	var diags hcl.Diagnostics

	decoded, diags := hcldec.Decode(p.File.Body, spec, ctx)
	if diags.HasErrors() {
		return nil, diags
	}

	session := new(session.Session)

	it := decoded.ElementIterator()
	for it.Next() {
		_, value := it.Element()
		if diagsDecode := session.Decode(value); diagsDecode.HasErrors() {
			diags = diags.Extend(diagsDecode)
			continue
		}
	}

	return session, diags
}

const glazeEnvPrefix = "GLAZE_ENV_"

func ResolveProfilePath(profilePath string) (string, error) {
	if profilePath == "" {
		cwd, err := os.Getwd()
		if err != nil {
			return profilePath, fmt.Errorf("could not read current working directory: %s", err)
		}

		profilePath = filepath.Join(cwd, ".glaze")

		if !files.FileExists(profilePath) && os.Getenv("GLAZE_PATH") != "" {
			profilePath = filepath.Join(os.Getenv("GLAZE_PATH"), ".glaze")
		}
	}

	if !files.FileExists(profilePath) {
		return profilePath, fmt.Errorf(
			"profile `%s` not found on the specified path, the current directory, or the GLAZE_PATH environment variable",
			profilePath,
		)
	}

	return profilePath, nil
}

func CollectVariables(flaggedVariables []string) (map[string]cty.Value, error) {
	variables := make(map[string]cty.Value)

	// We import environmental variables first:
	{
		for _, env := range os.Environ() {
			if !strings.HasPrefix(env, glazeEnvPrefix) {
				continue
			}

			env = strings.TrimPrefix(env, glazeEnvPrefix)

			if !strings.Contains(env, "=") {
				continue
			}

			parts := strings.SplitN(env, "=", 2)

			variables[parts[0]] = cty.StringVal(parts[1])
		}
	}

	// Next, import all variables passed by the --var flag:
	{
		for _, flag := range flaggedVariables {
			parts := strings.SplitN(flag, "=", 2)

			variables[strings.TrimSpace(parts[0])] = cty.StringVal(parts[1])
		}
	}

	// Finally, we add some default variables that might be useful:
	{
		cwd, err := os.Getwd()
		if err != nil {
			return variables, fmt.Errorf("could not read current working directory: %s", err)
		}

		variables["path"] = cty.ObjectVal(map[string]cty.Value{
			"base": cty.StringVal(filepath.Base(cwd)),
			"cwd":  cty.StringVal(cwd),
		})
	}

	return variables, nil
}

func BuildEvalContext(variables map[string]cty.Value) *hcl.EvalContext {
	return &hcl.EvalContext{
		Variables: variables,
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
}
