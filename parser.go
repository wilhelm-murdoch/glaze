package glaze

import (
	"os"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/hashicorp/hcl/v2/hclparse"
	"github.com/wilhelm-murdoch/glaze/models"
)

type Parser struct {
	diagsWriter hcl.DiagnosticWriter
	File        *hcl.File
	parser      *hclparse.Parser
	diags       hcl.Diagnostics
}

func NewParser(path string) *Parser {
	parser := hclparse.NewParser()
	file, diags := parser.ParseHCLFile(path)

	return &Parser{
		parser:      parser,
		File:        file,
		diags:       diags,
		diagsWriter: hcl.NewDiagnosticTextWriter(os.Stdout, map[string]*hcl.File{path: file}, 78, true),
	}
}

func (p *Parser) Decode(s hcldec.Spec, ctx *hcl.EvalContext) *models.Session {
	var session *models.Session

	decoded, diags := hcldec.Decode(p.File.Body, s, ctx)
	if diags.HasErrors() {
		p.diags = p.diags.Extend(diags)
		return session
	}

	it := decoded.ElementIterator()
	for it.Next() {
		_, value := it.Element()

		session = new(models.Session)
		diags := session.Decode(value)

		if diags.HasErrors() {
			p.diags = diags.Extend(diags)
			continue
		}
	}

	return session
}

func (p *Parser) AppendDiag(diag *hcl.Diagnostic) {
	p.diags = p.diags.Append(diag)
}

func (p *Parser) WriteDiags() error {
	if err := p.diagsWriter.WriteDiagnostics(p.diags); err != nil {
		return err
	}

	return nil
}

func (p *Parser) HasErrors() bool {
	return p.diags.HasErrors()
}
