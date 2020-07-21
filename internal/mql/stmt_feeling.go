package mql

import "bytes"

// FeelingStatement represents a command for looking at a room or object.
type FeelingStatement struct {
	Ident  string // feeling. If not found, respond with something like "what?"
	Object string
}

func (s *FeelingStatement) KeyTok() Token {
	return IDENT
}

// parseFeelingStatement parses a feeling command and returns a Statement AST object.
// This function assumes the IDENT token has already been consumed.
// TODO add target for feelings
func (p *Parser) parseFeelingStatement() (*FeelingStatement, error) {
	p.Unscan()
	stmt := &FeelingStatement{}
	var err error

	stmt.Ident, err = p.ParseIdent()
	if err != nil {
		return nil, err
	}

	tok, _, _ := p.ScanIgnoreWhitespace()
	switch tok {
	case AT, TO:
		obj, err := p.ParseIdent()
		if err != nil {
			return nil, err
		}
		stmt.Object = obj
	}
	return stmt, nil
}

func (s *FeelingStatement) String() string {
	var buf bytes.Buffer

	if s.Ident != "" {
		_, _ = buf.WriteString(s.Ident)
	}

	return buf.String()
}
