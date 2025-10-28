package save

import (
	"errors"

	"github.com/urfave/cli/v2"
)

type Action struct {
	ctx *cli.Context
}

func NewAction(ctx *cli.Context) *Action {
	return &Action{
		ctx: ctx,
	}
}

// Run attempts to save the specified tmux session as a canonical glaze definition file.
func (a *Action) Run() error {
	return errors.New(
		"dynamically saving an existing tmux layout as a glaze definition file is not yet supported",
	)
}
