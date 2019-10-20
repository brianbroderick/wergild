package main

import "bytes"

// FeelingStatement represents a command for looking at a room or object.
type FeelingStatement struct {
	player *Player
	ident  string // feeling. If not found, respond with something like "what?"
}

func (s *FeelingStatement) execute() {
	switch s.ident {
	case "laugh":
		s.player.connection.Write("You fall down laughing.\n")
	default:
		s.player.connection.Write("what? \n")
	}
}

// parseFeelingStatement parses a feeling command and returns a Statement AST object.
// This function assumes the IDENT token has already been consumed.
// TODO add target for feelings
func (p *Parser) parseFeelingStatement() (*FeelingStatement, error) {
	p.Unscan()
	stmt := &FeelingStatement{}
	var err error

	stmt.ident, err = p.ParseIdent()
	if err != nil {
		return nil, err
	}

	return stmt, nil
}

func (s *FeelingStatement) String() string {
	var buf bytes.Buffer

	if s.ident != "" {
		_, _ = buf.WriteString(s.ident)
	}

	return buf.String()
}

func (s *FeelingStatement) setPlayer(player *Player) {
	s.player = player
}
