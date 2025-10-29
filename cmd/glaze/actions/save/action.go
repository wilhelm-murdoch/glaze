package save

import (
	"github.com/urfave/cli/v2"

	"github.com/wilhelm-murdoch/glaze/cmd/glaze/actions"
)

type Action struct {
	actions.BaseAction
}

func NewAction(ctx *cli.Context) (*Action, error) {
	base, err := actions.NewBaseAction(ctx)
	if err != nil {
		return nil, err
	}

	return &Action{
		BaseAction: *base,
	}, nil
}
