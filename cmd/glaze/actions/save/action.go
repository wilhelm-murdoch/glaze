package save

import (
	"fmt"

	"github.com/urfave/cli/v2"
	"github.com/wilhelm-murdoch/glaze/internal/parser"
)

func Run(ctx *cli.Context) error {
	profilePath, err := parser.ResolveProfilePath(ctx.Args().First())
	if err != nil {
		return err
	}

	fmt.Println(profilePath)

	return nil
}
