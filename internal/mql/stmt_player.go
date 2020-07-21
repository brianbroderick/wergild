package mql

import "bytes"

// ScoreStatement represents a command for looking at a room or object.
type ScoreStatement struct{}

func (s *ScoreStatement) String() string {
	var buf bytes.Buffer
	_, _ = buf.WriteString("SCORE")

	return buf.String()
}

func (s *ScoreStatement) KeyTok() Token {
	return SCORE
}

// parseScoreStatement parses a look command and returns a Statement AST object.
// This function assumes the Score token has already been consumed.
func (p *Parser) parseScoreStatement() (*ScoreStatement, error) {
	return &ScoreStatement{}, nil
}
