package mud

import (
	"github.com/brianbroderick/wergild/internal/mql"
)

// QuitExpression represents a command for exiting the game.
type QuitExpression struct {
	mob *Mob
}

func (s *QuitExpression) Execute() {
	WorldInstance.getRoom(s.mob.CurrentRoom).exitWorld(s.mob)
	s.mob.conn.Write("You slowly fade away.\n")
	s.mob.conn.conn.Close()
	ServerInstance.onClientConnectionClosed(s.mob.conn, nil)
}

//********
// IMAGINE
//********

// ImagineExpression allows you to create rooms, objects, and mobs
type ImagineExpression struct {
	mob  *Mob
	stmt *mql.ImagineStatement
}

func (s *ImagineExpression) Execute() {
	myStr := "You start to concentrate really hard.\n"
	s.mob.myMessageToChannel(myStr)

	room := WorldInstance.getRoom(s.mob.CurrentRoom)
	if s.stmt.Name != "" {

		room.newRoom(s.stmt)

		s.mob.myMessageToChannel("The area around you expands...\n\n")
		room.showTo(s.mob)
	}

	if s.stmt.Name == "" {
		room.updateRoom(s.stmt)

		if s.stmt.Location != "" {
			s.mob.myMessageToChannel("The area around you begins to shift...\n\n")
			room.showTo(s.mob)
		} else if s.stmt.Listen != "" {
			s.mob.myMessageToChannel("The sound around you begins to shift...\n\n")
			room.showListenTo(s.mob)
		} else if s.stmt.Smell != "" {
			s.mob.myMessageToChannel("The smell around you begins to shift...\n\n")
			room.showSmellTo(s.mob)
		}
		return
	}
}
