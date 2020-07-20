package mql

import "bytes"

// QuitStatement represents a command for exiting the game.
type QuitStatement struct{}

func (s *QuitStatement) String() string {
	var buf bytes.Buffer
	_, _ = buf.WriteString("QUIT")

	return buf.String()
}

func (s *QuitStatement) KeyTok() Token {
	return QUIT
}

// parseQuitStatement parses a look command and returns a Statement AST object.
// This function assumes the QUIT token has already been consumed.
func (p *Parser) parseQuitStatement() (*QuitStatement, error) {
	return &QuitStatement{}, nil
}
