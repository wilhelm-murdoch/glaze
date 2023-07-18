package glaze

import (
	"os"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/hashicorp/hcl/v2/hclparse"
	"github.com/wilhelm-murdoch/glaze/models"
)

type Parser struct {
	parser      *hclparse.Parser
	File        *hcl.File
	diags       hcl.Diagnostics
	diagsWriter hcl.DiagnosticWriter
}

func NewParser() *Parser {
	return &Parser{
		parser: hclparse.NewParser(),
	}
}

func (p *Parser) Open(path string) {
	file, diags := p.parser.ParseHCLFile(path)

	p.File = file
	p.diagsWriter = hcl.NewDiagnosticTextWriter(os.Stdout, map[string]*hcl.File{path: p.File}, 78, true)

	if diags.HasErrors() {
		p.diags = p.diags.Extend(diags)
		return
	}
}

func (p *Parser) Decode(s hcldec.Spec) []*models.Session {
	decoded, diags := hcldec.Decode(p.File.Body, s, nil)
	if diags.HasErrors() {
		p.diags = p.diags.Extend(diags)
		return nil
	}

	var sessions []*models.Session

	it := decoded.ElementIterator()
	for it.Next() {
		_, value := it.Element()

		session := new(models.Session)
		diags := session.Decode(value)

		if diags.HasErrors() {
			p.diags = diags.Extend(diags)
			continue
		}

		sessions = append(sessions, session)
	}

	return sessions
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
