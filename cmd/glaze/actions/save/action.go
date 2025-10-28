package save

import (
	"errors"

	"github.com/urfave/cli/v2"
)

// Run attempts to save the specified tmux session as a canonical glaze definition file.
func Run(ctx *cli.Context) error {
	return errors.New(
		"dynamically saving an existing tmux layout as a glaze definition file is not yet supported",
	)
}
