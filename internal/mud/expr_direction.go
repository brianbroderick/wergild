package mud

import (
	"github.com/brianbroderick/wergild/internal/mql"
)

// DirectionExpression represents a command for looking at a room or object.
type DirectionExpression struct {
	mob  *Mob
	stmt *mql.DirectionStatement
}

func (s *DirectionExpression) Execute() {
	room := WorldInstance.getRoom(s.mob.CurrentRoom)

	if val, ok := room.ExitMap[mql.Tokens[s.stmt.Token]]; ok {
		s.mob.CurrentRoom = val
		room.exitRoom(s.mob, s.stmt.Token.String())
		newRoom := WorldInstance.getRoom(s.mob.CurrentRoom)
		newRoom.showTo(s.mob)
		newRoom.enterRoom(s.mob)
		return
	}

	s.mob.conn.Write("You're unable to go that way.\n")
}
