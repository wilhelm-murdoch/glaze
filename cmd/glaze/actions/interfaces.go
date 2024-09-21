package actions

import (
	"github.com/urfave/cli/v2"
	"github.com/wilhelm-murdoch/glaze"
)

type Actionable interface {
	glaze.Common
	Run(*cli.Context) error
}
