package mud

import (
	"fmt"

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

	if s.stmt.Mob != "" {
		s.NewMob(room)
		return
	}

	if s.stmt.Name != "" {
		s.NewRoom(room)
		return
	} else {
		s.UpdateRoom(room)
		return
	}

}

func (s *ImagineExpression) NewRoom(room *Room) {
	err := room.newRoom(s.stmt)
	if err != nil {
		s.mob.myMessageToChannel(err.Error() + "\n\n")
	} else {

		s.mob.myMessageToChannel("The area around you expands...\n\n")

		// Reload room
		room = WorldInstance.getRoom(s.mob.CurrentRoom)
		room.showTo(s.mob)
	}
}

func (s *ImagineExpression) UpdateRoom(room *Room) {
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
}

func (s *ImagineExpression) NewMob(room *Room) {
	CreateNPCMob(room, s.stmt)

	s.mob.myMessageToChannel(fmt.Sprintf("The area in front of you begins to shift...\n\nBefore your eyes something springs to existence...\n\nIt's %s!\n\n", s.stmt.Mob))
}
