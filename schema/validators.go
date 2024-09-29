package schema

import (
	"fmt"

	"github.com/hashicorp/hcl/v2"
	"github.com/wilhelm-murdoch/glaze"
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
					out = out.Extend(glaze.ContainsDiagnostic(fmt.Sprintf(`session option "%s"`, option), value, choices))
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
