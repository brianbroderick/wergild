package mud

import (
	"bytes"
)

// DirectionStatement represents a command for looking at a room or object.
type DirectionStatement struct {
	player *Player
	token  Token
}

// parseDirectionStatement parses a look command and returns a Statement AST object.
// This function assumes the direction token has already been consumed.
func (p *Parser) parseDirectionStatement() (*DirectionStatement, error) {
	p.Unscan()
	stmt := &DirectionStatement{}
	stmt.token, _, _ = p.ScanIgnoreWhitespace()

	return stmt, nil
}

func (s *DirectionStatement) String() string {
	var buf bytes.Buffer
	_, _ = buf.WriteString(tokens[s.token])

	return buf.String()
}

func (s *DirectionStatement) execute() {
	room := WorldInstance.getRoom(s.player.CurrentRoom)

	if val, ok := room.ExitMap[tokens[s.token]]; ok {
		s.player.CurrentRoom = val
		newRoom := WorldInstance.getRoom(s.player.CurrentRoom)
		newRoom.showTo(s.player)
		return
	}

	s.player.connection.Write("You're unable to go that way.\n")
}

func (s *DirectionStatement) setPlayer(player *Player) {
	s.player = player
}