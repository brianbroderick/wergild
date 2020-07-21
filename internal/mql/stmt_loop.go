package mql

import (
	"bytes"
	"strconv"
)

// LoopStatement represents a command for repeating another command X times
type LoopStatement struct {
	I     int
	NStmt Statement // msg to parse and run i times
}

func (s *LoopStatement) KeyTok() Token {
	return INTEGER
}

// parseLoopStatement parses an integer to be looped on and returns a Statement AST object.
// This function assumes the INTEGER token has already been consumed.
// TODO actually loop
func (p *Parser) parseLoopStatement() (*LoopStatement, error) {
	p.Unscan()
	stmt := &LoopStatement{}
	var err error

	stmt.I, err = p.ParseInt(1, 50) // max of 50
	if err != nil {
		return nil, err
	}

	nStmt, err := p.ParseStatement()
	if err != nil {
		return nil, err
	}

	stmt.NStmt = nStmt

	return stmt, nil
}

func (s *LoopStatement) String() string {
	var buf bytes.Buffer

	if s.I != 0 {
		_, _ = buf.WriteString(strconv.Itoa(s.I))
	}

	return buf.String()
}
