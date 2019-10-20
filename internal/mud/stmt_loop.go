package mud

import (
	"bytes"
	"strconv"
)

// LoopStatement represents a command for repeating another command X times
type LoopStatement struct {
	player *Player
	i      int
	nStmt  Statement // msg to parse and run i times
}

// parseLoopStatement parses an integer to be looped on and returns a Statement AST object.
// This function assumes the INTEGER token has already been consumed.
// TODO actually loop
func (p *Parser) parseLoopStatement() (*LoopStatement, error) {
	p.Unscan()
	stmt := &LoopStatement{}
	var err error

	stmt.i, err = p.ParseInt(1, 50) // max of 50
	if err != nil {
		return nil, err
	}

	nStmt, err := p.ParseStatement()
	if err != nil {
		return nil, err
	}

	stmt.nStmt = nStmt

	return stmt, nil
}

func (s *LoopStatement) String() string {
	var buf bytes.Buffer

	if s.i != 0 {
		_, _ = buf.WriteString(strconv.Itoa(s.i))
	}

	return buf.String()
}

func (s *LoopStatement) setPlayer(player *Player) {
	s.player = player
}

func (s *LoopStatement) execute() {
	s.nStmt.setPlayer(s.player)

	for i := 0; i < s.i; i++ {
		s.nStmt.execute()
	}
}
