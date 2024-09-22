package glaze

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/hashicorp/hcl/v2"
	"github.com/zclconf/go-cty/cty"
	"github.com/zclconf/go-cty/cty/function"
	"github.com/zclconf/go-cty/cty/function/stdlib"
)

const glazeEnvPrefix = "GLAZE_ENV_"

func ResolveProfilePath(profilePath string) (string, error) {
	if profilePath == "" {
		cwd, err := os.Getwd()
		if err != nil {
			return profilePath, fmt.Errorf("could not read current working directory: %s", err)
		}

		profilePath = filepath.Join(cwd, ".glaze")

		if !FileExists(profilePath) && os.Getenv("GLAZE_PATH") != "" {
			profilePath = filepath.Join(os.Getenv("GLAZE_PATH"), ".glaze")
		}
	}

	if !FileExists(profilePath) {
		return profilePath, fmt.Errorf("profile `%s` not found on the specified path, the current directory, or the GLAZE_PATH environment variable", profilePath)
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
