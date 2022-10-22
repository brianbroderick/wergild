package mud

import "github.com/brianbroderick/wergild/internal/mql"

// HelpExpression represents a command for looking at a room or object.
type HelpExpression struct {
	mob  *Mob
	stmt *mql.HelpStatement
}

func (s *HelpExpression) Execute() {
	// room := WorldInstance.getRoom(s.mob.CurrentRoom)

	s.mob.myMessageToChannel("Help\n")
}
