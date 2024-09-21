package actions

import (
	"github.com/urfave/cli/v2"
	"github.com/wilhelm-murdoch/glaze"
)

type Save struct {
	glaze.Common
}

func (s Save) Run(ctx *cli.Context) error {
	// profilePath := ctx.Args().First()
	return nil
}
