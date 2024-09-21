package actions

import (
	"fmt"

	"github.com/urfave/cli/v2"
	"github.com/wilhelm-murdoch/glaze"
)

type Save struct {
	glaze.Common
}

func (s Save) Run(ctx *cli.Context) error {
	profilePath, err := s.ResolveProfilePath(ctx.Args().First())
	if err != nil {
		return err
	}

	fmt.Println(profilePath)

	return nil
}
