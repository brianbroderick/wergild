package mud

import (
	"fmt"
	"strings"

	"github.com/brianbroderick/wergild/internal/mql"
)

//*******
// LOOK
//*******

// LookExpression represents a command for looking at a room or object.
type LookExpression struct {
	mob  *Mob
	stmt *mql.LookStatement
}

func (s *LookExpression) Execute() {
	currentRoom := WorldInstance.getRoom(s.mob.CurrentRoom)

	switch s.stmt.Token {
	case mql.EOF:
		currentRoom.showTo(s.mob)
		s.mob.conn.Write("\n")
		return
	case mql.NORTH, mql.SOUTH, mql.EAST, mql.WEST, mql.UP, mql.DOWN:
		directions := [6]mql.Token{mql.NORTH, mql.SOUTH, mql.EAST, mql.WEST, mql.UP, mql.DOWN}
		for _, direction := range directions {
			if s.stmt.Token == direction {
				s.mob.conn.Write(fmt.Sprintf("You look %s.\n", strings.ToLower(mql.Tokens[s.stmt.Token])))
				return
			}
		}
	case mql.NIL:
		s.mob.conn.Write("Look AT or IN something, or what?\n")
		return
	default:
		points, err := queryPointOfInterest(s.mob.CurrentRoom, s.stmt.Ident)
		if err != nil {
			s.mob.conn.Write("There's nothing interesting about that.\n")
		}

		if len(points) == 0 {
			s.mob.conn.Write(s.stmt.Ident + " wasn't found.\n")
			return
		}

		for _, point := range points {
			s.mob.conn.Write(point.Desc + "\n")
		}
		return
	}

	s.mob.conn.Write("Please try looking a different way.\n")
}

//*******
// LISTEN
//*******

// ListenExpression represents a command for listening to a room or object.
type ListenExpression struct {
	mob  *Mob
	stmt *mql.ListenStatement
}

func (s *ListenExpression) Execute() {
	currentRoom := WorldInstance.getRoom(s.mob.CurrentRoom)

	switch s.stmt.Token {
	case mql.EOF:
		currentRoom.showListenTo(s.mob)
		s.mob.conn.Write("\n")
		return
	case mql.NORTH, mql.SOUTH, mql.EAST, mql.WEST, mql.UP, mql.DOWN:
		directions := [6]mql.Token{mql.NORTH, mql.SOUTH, mql.EAST, mql.WEST, mql.UP, mql.DOWN}
		for _, direction := range directions {
			if s.stmt.Token == direction {
				s.mob.conn.Write(fmt.Sprintf("You listen %s.\n", strings.ToLower(mql.Tokens[s.stmt.Token])))
				return
			}
		}
	case mql.NIL:
		s.mob.conn.Write("Listen TO something, or what?\n")
		return
	default:
		points, err := queryPointOfInterestListen(s.mob.CurrentRoom, s.stmt.Ident)
		if err != nil {
			s.mob.conn.Write("There's nothing interesting about that.\n")
			return
		}

		if len(points) == 0 {
			s.mob.conn.Write("You listened to " + s.stmt.Ident + ", but there wasn't anything unusual about it.\n")
			return
		}

		for _, point := range points {
			s.mob.conn.Write(point.Listen + "\n")
		}
	}
}

//*******
// SMELL
//*******

// SmellExpression represents a command for smelling a room or object.
type SmellExpression struct {
	mob  *Mob
	stmt *mql.SmellStatement
}

func (s *SmellExpression) Execute() {
	currentRoom := WorldInstance.getRoom(s.mob.CurrentRoom)

	switch s.stmt.Token {
	case mql.EOF:
		currentRoom.showSmellTo(s.mob)
		s.mob.conn.Write("\n")
		return
	case mql.NORTH, mql.SOUTH, mql.EAST, mql.WEST, mql.UP, mql.DOWN:
		directions := [6]mql.Token{mql.NORTH, mql.SOUTH, mql.EAST, mql.WEST, mql.UP, mql.DOWN}
		for _, direction := range directions {
			if s.stmt.Token == direction {
				s.mob.conn.Write(fmt.Sprintf("You smell %s. But what does that smell like?\n", strings.ToLower(mql.Tokens[s.stmt.Token])))
				return
			}
		}
	case mql.IN:
		s.mob.conn.Write("Smelling in things is not recommended at this time. Who knows what might be in there?\n")
		return
	default:
		points, err := queryPointOfInterestSmell(s.mob.CurrentRoom, s.stmt.Ident)
		if err != nil {
			s.mob.conn.Write("There's nothing interesting about that.\n")
			return
		}

		if len(points) == 0 {
			s.mob.conn.Write("After trying to smell " + s.stmt.Ident + ", you don't notice anything interesting.\n")
			return
		}

		for _, point := range points {
			s.mob.conn.Write(point.Smell + "\n")
		}
	}
}

//*******
// TOUCH
//*******

// TouchExpression represents a command for touching an object.
type TouchExpression struct {
	mob  *Mob
	stmt *mql.TouchStatement
}

func (s *TouchExpression) Execute() {
	switch s.stmt.Token {
	case mql.EOF:
		s.mob.conn.Write("Maybe you should keep your hands to yourself.\n")
		return
	case mql.NORTH, mql.SOUTH, mql.EAST, mql.WEST, mql.UP, mql.DOWN:
		directions := [6]mql.Token{mql.NORTH, mql.SOUTH, mql.EAST, mql.WEST, mql.UP, mql.DOWN}
		for _, direction := range directions {
			if s.stmt.Token == direction {
				s.mob.conn.Write(fmt.Sprintf("You touch %s. But what does that feel like?\n", strings.ToLower(mql.Tokens[s.stmt.Token])))
				return
			}
		}
	case mql.IN:
		s.mob.conn.Write("Touching in things is not recommended at this time. Who knows what might be in there?\n")
		return
	default:
		points, err := queryPointOfInterestTouch(s.mob.CurrentRoom, s.stmt.Ident)
		if err != nil {
			s.mob.conn.Write("There's nothing interesting about that.\n")
			return
		}

		if len(points) == 0 {
			s.mob.conn.Write("After trying to touch " + s.stmt.Ident + ", you don't notice anything interesting.\n")
			return
		}

		for _, point := range points {
			s.mob.conn.Write(point.Touch + "\n")
		}
	}
}
