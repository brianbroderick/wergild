package mud

import "github.com/brianbroderick/wergild/internal/mql"

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

//******
// IMAGINE
//******

// ImagineExpression allows you to create rooms, objects, and mobs
type ImagineExpression struct {
	mob  *Mob
	stmt *mql.ImagineStatement
}

func (s *ImagineExpression) Execute() {
	myStr := "You start to concentrate really hard.\n"
	s.mob.myMessageToChannel(myStr)
}
