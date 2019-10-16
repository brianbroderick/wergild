package main

import (
	"bytes"
	"fmt"
)

// DirectionStatement represents a command for looking at a room or object.
type DirectionStatement struct {
	player *Player
	token  Token
}

// parseDirectionStatement parses a look command and returns a Statement AST object.
// This function assumes the QUIT token has already been consumed.
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
	room := ServerInstance.getRoom(s.player.CurrentRoom)

	fmt.Printf("tok: %s \n", tokens[s.token])

	if val, ok := room.ExitMap[tokens[s.token]]; ok {
		// fmt.Printf("%v", val.destination)
		s.player.CurrentRoom = val
		newRoom := ServerInstance.getRoom(s.player.CurrentRoom)
		newRoom.showTo(s.player)
	}
}

// func moveTo(player *Player, arguments []string) {
// 	room := ServerInstance.getRoom(player.CurrentRoom)

// 	if val, ok := room.exits[arguments[0]]; ok {
// 		// fmt.Printf("%v", val.destination)
// 		player.CurrentRoom = val.destination
// 		newRoom := ServerInstance.getRoom(player.CurrentRoom)
// 		newRoom.showTo(player)
// 	}
// }

func (s *DirectionStatement) setPlayer(player *Player) {
	s.player = player
}
