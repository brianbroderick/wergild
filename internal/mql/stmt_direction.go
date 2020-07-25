package mql

import (
	"bytes"

	"github.com/brianbroderick/logit"
)

// DirectionStatement represents a command for looking at a room or object.
type DirectionStatement struct {
	Token Token
}

// parseDirectionStatement parses a look command and returns a Statement AST object.
// This function assumes the direction token has already been consumed.
func (p *Parser) parseDirectionStatement() (*DirectionStatement, error) {
	p.Unscan()
	stmt := &DirectionStatement{}
	stmt.Token, _, _ = p.ScanIgnoreWhitespace()

	logit.Info(Tokens[stmt.Token])

	return stmt, nil
}

func (s *DirectionStatement) KeyTok() Token {
	return DIRECTION
}

func (s *DirectionStatement) String() string {
	var buf bytes.Buffer
	_, _ = buf.WriteString(Tokens[s.Token])

	return buf.String()
}
