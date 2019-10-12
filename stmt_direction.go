package main

import "bytes"

// DirectionStatement represents a command for looking at a room or object.
type DirectionStatement struct {
	player *Player
}

func (s *DirectionStatement) String() string {
	var buf bytes.Buffer
	// _, _ = buf.WriteString("QUIT")

	return buf.String()
}

func (s *DirectionStatement) setPlayer(player *Player) {
	s.player = player
}

// parseDirectionStatement parses a look command and returns a Statement AST object.
// This function assumes the QUIT token has already been consumed.
func (p *Parser) parseDirectionStatement() (*DirectionStatement, error) {
	return &DirectionStatement{}, nil
}

func (s *DirectionStatement) execute() {

}
