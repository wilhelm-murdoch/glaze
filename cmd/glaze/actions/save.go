package actions

import (
	"fmt"

	"github.com/urfave/cli/v2"
	"github.com/wilhelm-murdoch/glaze"
)

func Save(ctx *cli.Context) error {
	profilePath, err := glaze.ResolveProfilePath(ctx.Args().First())
	if err != nil {
		return err
	}

	fmt.Println(profilePath)

	return nil
}
