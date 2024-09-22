package glaze

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/hashicorp/hcl/v2/hclparse"
	"github.com/wilhelm-murdoch/glaze/models"
)

type Parser struct {
	File   *hcl.File
	parser *hclparse.Parser
}

func NewParser(path string) (*Parser, hcl.Diagnostics) {
	parser := hclparse.NewParser()
	file, diags := parser.ParseHCLFile(path)

	if diags.HasErrors() {
		return nil, diags
	}

	return &Parser{
		parser: parser,
		File:   file,
	}, nil
}

func (p *Parser) Decode(spec hcldec.Spec, ctx *hcl.EvalContext) (*models.Session, hcl.Diagnostics) {
	var (
		session *models.Session
		diags   hcl.Diagnostics
	)

	decoded, diags := hcldec.Decode(p.File.Body, spec, ctx)
	if diags.HasErrors() {
		return session, diags
	}

	it := decoded.ElementIterator()
	for it.Next() {
		_, value := it.Element()

		session = new(models.Session)
		if diagsDecode := session.Decode(value); diagsDecode.HasErrors() {
			diags = diags.Extend(diagsDecode)
			continue
		}
	}

	return session, diags
}
