package main

import (
	"bytes"
	"fmt"
	"strings"
)

// LookStatement represents a command for looking at a room or object.
type LookStatement struct {
	player *Player
	ident  string // object that is being looked at, ex: `chair` in `look at chair`
	room   int
	token  Token // usually a direction. i.e. `north` in `look north`
}

// parseLookStatement parses a look command and returns a Statement AST object.
// This function assumes the LOOK token has already been consumed.
func (p *Parser) parseLookStatement() (*LookStatement, error) {
	stmt := &LookStatement{}
	var err error

	tok, _, _ := p.ScanIgnoreWhitespace()

	directions := [6]Token{NORTH, SOUTH, EAST, WEST, UP, DOWN}
	for _, direction := range directions {
		if tok == direction {
			stmt.token = tok
			return stmt, nil
		}
	}

	switch tok {
	case AT:
		stmt.token = tok
	case IN:
		stmt.token = tok
	default:
		p.Unscan()
		return stmt, nil
	}

	stmt.ident, err = p.ParseIdent()
	if err != nil {
		return nil, err
	}

	return stmt, nil
}

func (s *LookStatement) String() string {
	var buf bytes.Buffer
	_, _ = buf.WriteString("LOOK")

	directions := [6]Token{NORTH, SOUTH, EAST, WEST, UP, DOWN}
	for _, direction := range directions {
		if s.token == direction {
			_, _ = buf.WriteString(" " + tokens[s.token])
		}
	}

	switch s.token {
	case AT, ON, IN:
		_, _ = buf.WriteString(" " + tokens[s.token] + " ")
		_, _ = buf.WriteString(s.ident)
	}

	// if s.ident != "" {
	// 	_, _ = buf.WriteString(" AT ")
	// 	_, _ = buf.WriteString(s.ident)
	// }

	return buf.String()
}

func (s *LookStatement) setPlayer(player *Player) {
	s.player = player
}

func (s *LookStatement) execute() {
	currentRoom := ServerInstance.getRoom(s.player.CurrentRoom)

	fmt.Printf("TOK: %s -- %s \n", s.token, s.ident)

	if s.token == ILLEGAL && s.ident == "" {
		currentRoom.showTo(s.player)
		s.player.connection.Write("\n")
	} else if s.token != ILLEGAL && s.ident == "" {
		directions := [6]Token{NORTH, SOUTH, EAST, WEST, UP, DOWN}
		for _, direction := range directions {
			if s.token == direction {
				s.player.connection.Write(fmt.Sprintf("You look %s.\n", strings.ToLower(tokens[s.token])))
			}
		}
	} else if s.ident != "" {
		for _, item := range currentRoom.Inventory {
			if containsString(item.Keys, s.ident) {
				s.player.connection.Write(item.Description + "\n")
				return
			}

			// if len(arguments) > 2 && containsString(item.Keys, arguments[2]) {
			// 	player.connection.Write(item.Description + "\n")
			// 	return
			// }
		}
		s.player.connection.Write("You didn't see anything unusual about " + s.ident + ".\n")
	} else {
		s.player.connection.Write("Look AT or IN something, or what?\n")
	}
}

// func look(player *Player, arguments []string) {
// 	currentRoom := ServerInstance.getRoom(player.CurrentRoom)

// 	if len(arguments) == 0 {
// 		currentRoom.showTo(player)
// 		player.connection.Write("\n")
// 	} else if containsString([]string{"at", "in"}, arguments[0]) {
// 		for _, item := range currentRoom.Inventory {
// 			if containsString(item.Keys, arguments[1]) {
// 				player.connection.Write(item.Description + "\n")
// 				return
// 			}

// 			if len(arguments) > 2 && containsString(item.Keys, arguments[2]) {
// 				player.connection.Write(item.Description + "\n")
// 				return
// 			}
// 		}
// 		player.connection.Write("You didn't see anything unusual about " + arguments[1] + ".\n")
// 	} else if arguments[0] == "travel" && arguments[1] == "up" {
// 		player.connection.Write("You look up towards the sky.\n")
// 	} else if arguments[0] == "travel" && arguments[1] == "down" {
// 		player.connection.Write("You look down towards the ground.\n")
// 	} else if arguments[0] == "travel" {
// 		player.connection.Write(fmt.Sprintf("You look towards the %s.\n", arguments[1]))
// 	} else {
// 		player.connection.Write("Look AT or IN something, or what?\n")
// 	}

// }
