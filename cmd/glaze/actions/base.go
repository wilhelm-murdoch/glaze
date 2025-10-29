package actions

import (
	"github.com/urfave/cli/v2"

	"github.com/wilhelm-murdoch/glaze/internal/diagnostics"
	ge "github.com/wilhelm-murdoch/glaze/internal/errors" // ge = "Glaze Errors"
	"github.com/wilhelm-murdoch/glaze/internal/parser"
	"github.com/wilhelm-murdoch/glaze/internal/profile"
)

type BaseAction struct {
	Context            *cli.Context
	DiagnosticsManager *diagnostics.DiagnosticsManager
	Parser             *parser.Parser
	ProfilePath        string
}

func NewBaseAction(ctx *cli.Context) (*BaseAction, error) {
	profilePath, err := profile.ResolveProfilePath(ctx.Args().First())
	if err != nil {
		return nil, err
	}

	diagsManager := diagnostics.NewDiagnosticsManager(profilePath)
	if diagsManager.HasErrors() {
		diagsManager.Write()
		return nil, ge.ErrorInvalidGlazeDefinition
	}

	parser, parserDiags := parser.NewParser(profilePath)
	if parserDiags.HasErrors() {
		diagsManager.Write()
		return nil, ge.ErrorInvalidGlazeDefinition
	}

	return &BaseAction{
		Context:            ctx,
		DiagnosticsManager: diagsManager,
		Parser:             parser,
		ProfilePath:        profilePath,
	}, nil
}

func (ba *BaseAction) Run() error {
	return ge.ErrorNotYetImplemented
}
