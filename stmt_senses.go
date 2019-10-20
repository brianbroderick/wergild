package main

import (
	"bytes"
	"fmt"
	"strings"
)

//*******
// LOOK
//*******

// LookStatement represents a command for looking at a room or object.
type LookStatement struct {
	player *Player
	ident  string // object that is being looked at, ex: `chair` in `look at chair`
	token  Token  // usually a direction. i.e. `north` in `look north`
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
	case AT, IN:
		stmt.token = tok
	case EOF:
		stmt.token = EOF
		p.Unscan()
		return stmt, nil
	default:
		stmt.token = NIL
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
	case NIL:
		_, _ = buf.WriteString(s.ident)
	}

	return buf.String()
}

func (s *LookStatement) setPlayer(player *Player) {
	s.player = player
}

func (s *LookStatement) execute() {
	currentRoom := ServerInstance.getRoom(s.player.CurrentRoom)

	switch s.token {
	case EOF:
		currentRoom.showTo(s.player)
		s.player.connection.Write("\n")
		return
	case NORTH, SOUTH, EAST, WEST, UP, DOWN:
		directions := [6]Token{NORTH, SOUTH, EAST, WEST, UP, DOWN}
		for _, direction := range directions {
			if s.token == direction {
				s.player.connection.Write(fmt.Sprintf("You look %s.\n", strings.ToLower(tokens[s.token])))
				return
			}
		}
	case NIL:
		s.player.connection.Write("Look AT or IN something, or what?\n")
		return
	default:
		points, err := queryPointOfInterest(s.player.CurrentRoom, s.ident)
		if err != nil {
			s.player.connection.Write("There's nothing interesting about that.\n")
		}

		if len(points) == 0 {
			s.player.connection.Write(s.ident + " wasn't found.\n")
			return
		}

		for _, point := range points {
			s.player.connection.Write(point.Desc + "\n")
		}
		return
	}

	s.player.connection.Write("Please try looking a different way.\n")
}

//*******
// LISTEN
//*******

// ListenStatement represents a command for listening to a room or object.
type ListenStatement struct {
	player *Player
	ident  string // object that is being listened to, ex: `clock` in `listen to clock`
	token  Token  // usually a direction. i.e. `north` in `listen north`
}

// parseListenStatement parses a listen command and returns a Statement AST object.
// This function assumes the LISTEN token has already been consumed.
func (p *Parser) parseListenStatement() (*ListenStatement, error) {
	stmt := &ListenStatement{}
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
	case TO, IN:
		stmt.token = tok
	case EOF:
		stmt.token = EOF
		p.Unscan()
		return stmt, nil
	default:
		stmt.token = NIL
		p.Unscan()
		return stmt, nil
	}

	stmt.ident, err = p.ParseIdent()
	if err != nil {
		return nil, err
	}

	return stmt, nil
}

func (s *ListenStatement) String() string {
	var buf bytes.Buffer
	_, _ = buf.WriteString("LISTEN")

	directions := [6]Token{NORTH, SOUTH, EAST, WEST, UP, DOWN}
	for _, direction := range directions {
		if s.token == direction {
			_, _ = buf.WriteString(" " + tokens[s.token])
		}
	}

	switch s.token {
	case TO, IN:
		_, _ = buf.WriteString(" " + tokens[s.token] + " ")
		_, _ = buf.WriteString(s.ident)
	case NIL:
		_, _ = buf.WriteString(s.ident)
	}

	return buf.String()
}

func (s *ListenStatement) setPlayer(player *Player) {
	s.player = player
}

func (s *ListenStatement) execute() {
	currentRoom := ServerInstance.getRoom(s.player.CurrentRoom)

	switch s.token {
	case EOF:
		currentRoom.showListenTo(s.player)
		s.player.connection.Write("\n")
		return
	case NORTH, SOUTH, EAST, WEST, UP, DOWN:
		directions := [6]Token{NORTH, SOUTH, EAST, WEST, UP, DOWN}
		for _, direction := range directions {
			if s.token == direction {
				s.player.connection.Write(fmt.Sprintf("You listen %s.\n", strings.ToLower(tokens[s.token])))
				return
			}
		}
	case NIL:
		s.player.connection.Write("Listen TO something, or what?\n")
		return
	default:
		points, err := queryPointOfInterestListen(s.player.CurrentRoom, s.ident)
		if err != nil {
			s.player.connection.Write("There's nothing interesting about that.\n")
			return
		}

		if len(points) == 0 {
			s.player.connection.Write("You listened to " + s.ident + ", but there wasn't anything unusual about it.\n")
			return
		}

		for _, point := range points {
			s.player.connection.Write(point.Listen + "\n")
		}
	}
}

//*******
// SMELL
//*******

// SmellStatement represents a command for smelling a room or object.
type SmellStatement struct {
	player *Player
	ident  string // object that is being smelled, ex: `chair` in `smell chair`
	token  Token  // usually not used, but you might try to `smell north`
}

func (s *SmellStatement) String() string {
	var buf bytes.Buffer
	_, _ = buf.WriteString("SMELL")

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
	case NIL:
		_, _ = buf.WriteString(s.ident)
	}

	return buf.String()
}

// parseSmellStatement parses a look command and returns a Statement AST object.
// This function assumes the Smell token has already been consumed.
func (p *Parser) parseSmellStatement() (*SmellStatement, error) {
	stmt := &SmellStatement{}
	var err error

	tok, _, lit := p.ScanIgnoreWhitespace()

	directions := [6]Token{NORTH, SOUTH, EAST, WEST, UP, DOWN}
	for _, direction := range directions {
		if tok == direction {
			stmt.token = tok
			return stmt, nil
		}
	}

	switch tok {
	case IN:
		stmt.token = tok
	case EOF:
		stmt.token = EOF
		p.Unscan()
		return stmt, nil
	default:
		stmt.token = NIL
		stmt.ident = lit
		return stmt, nil
	}

	stmt.ident, err = p.ParseIdent()
	if err != nil {
		return nil, err
	}

	return stmt, nil
}

func (s *SmellStatement) execute() {
	currentRoom := ServerInstance.getRoom(s.player.CurrentRoom)

	switch s.token {
	case EOF:
		currentRoom.showSmellTo(s.player)
		s.player.connection.Write("\n")
		return
	case NORTH, SOUTH, EAST, WEST, UP, DOWN:
		directions := [6]Token{NORTH, SOUTH, EAST, WEST, UP, DOWN}
		for _, direction := range directions {
			if s.token == direction {
				s.player.connection.Write(fmt.Sprintf("You smell %s. But what does that smell like?\n", strings.ToLower(tokens[s.token])))
				return
			}
		}
	case IN:
		s.player.connection.Write("Smelling in things is not recommended at this time. Who knows what might be in there?\n")
		return
	default:
		points, err := queryPointOfInterestSmell(s.player.CurrentRoom, s.ident)
		if err != nil {
			s.player.connection.Write("There's nothing interesting about that.\n")
			return
		}

		if len(points) == 0 {
			s.player.connection.Write("After trying to smell " + s.ident + ", you don't notice anything interesting.\n")
			return
		}

		for _, point := range points {
			s.player.connection.Write(point.Smell + "\n")
		}
	}
}

func (s *SmellStatement) setPlayer(player *Player) {
	s.player = player
}

//*******
// TOUCH
//*******

// TouchStatement represents a command for touching an object.
type TouchStatement struct {
	player *Player
	ident  string // object that is being touched, ex: `chair` in `touch chair`
	token  Token  // usually not used, but you might try to `touch north`
}

func (s *TouchStatement) String() string {
	var buf bytes.Buffer
	_, _ = buf.WriteString("TOUCH")

	directions := [6]Token{NORTH, SOUTH, EAST, WEST, UP, DOWN}
	for _, direction := range directions {
		if s.token == direction {
			_, _ = buf.WriteString(" " + tokens[s.token])
		}
	}

	switch s.token {
	case IN:
		_, _ = buf.WriteString(" " + tokens[s.token] + " ")
		_, _ = buf.WriteString(s.ident)
	case NIL:
		_, _ = buf.WriteString(s.ident)
	}

	return buf.String()
}

// parseTouchStatement parses a look command and returns a Statement AST object.
// This function assumes the Smell token has already been consumed.
func (p *Parser) parseTouchStatement() (*TouchStatement, error) {
	stmt := &TouchStatement{}
	var err error

	tok, _, lit := p.ScanIgnoreWhitespace()

	directions := [6]Token{NORTH, SOUTH, EAST, WEST, UP, DOWN}
	for _, direction := range directions {
		if tok == direction {
			stmt.token = tok
			return stmt, nil
		}
	}

	switch tok {
	case IN:
		stmt.token = tok
	case EOF:
		stmt.token = EOF
		p.Unscan()
		return stmt, nil
	default:
		stmt.token = NIL
		stmt.ident = lit
		return stmt, nil
	}

	stmt.ident, err = p.ParseIdent()
	if err != nil {
		return nil, err
	}

	return stmt, nil
}

func (s *TouchStatement) execute() {
	switch s.token {
	case EOF:
		s.player.connection.Write("Maybe you should keep your hands to yourself.\n")
		return
	case NORTH, SOUTH, EAST, WEST, UP, DOWN:
		directions := [6]Token{NORTH, SOUTH, EAST, WEST, UP, DOWN}
		for _, direction := range directions {
			if s.token == direction {
				s.player.connection.Write(fmt.Sprintf("You touch %s. But what does that feel like?\n", strings.ToLower(tokens[s.token])))
				return
			}
		}
	case IN:
		s.player.connection.Write("Touching in things is not recommended at this time. Who knows what might be in there?\n")
		return
	default:
		points, err := queryPointOfInterestTouch(s.player.CurrentRoom, s.ident)
		if err != nil {
			s.player.connection.Write("There's nothing interesting about that.\n")
			return
		}

		if len(points) == 0 {
			s.player.connection.Write("After trying to touch " + s.ident + ", you don't notice anything interesting.\n")
			return
		}

		for _, point := range points {
			s.player.connection.Write(point.Touch + "\n")
		}
	}
}

func (s *TouchStatement) setPlayer(player *Player) {
	s.player = player
}
