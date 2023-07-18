package actions

import (
	"github.com/k0kubun/pp"
	"github.com/wilhelm-murdoch/glaze"
)

func ActionApply(profilePath string) error {
	p := glaze.NewParser()
	p.Open(profilePath)

	sessions := p.Decode(glaze.PrimaryGlazeSpec)

	if p.HasErrors() {
		p.WriteDiags()
		return nil
	}

	pp.Print(sessions)

	return nil
}
