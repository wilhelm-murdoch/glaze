package save

import (
	"github.com/urfave/cli/v2"

	ge "github.com/wilhelm-murdoch/glaze/internal/errors" // ge = "Glaze Errors"
)

type Action struct {
	ctx *cli.Context
}

func NewAction(ctx *cli.Context) (*Action, error) {
	return &Action{
		ctx: ctx,
	}, nil
}

// Run attempts to save the specified tmux session as a canonical glaze definition file.
func (a *Action) Run() error {
	return ge.ErrorNotYetImplemented
}
