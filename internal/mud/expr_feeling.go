package mud

import (
	"fmt"

	"github.com/brianbroderick/wergild/internal/mql"
)

// FeelingExpression represents a command for looking at a room or object.
type FeelingExpression struct {
	mob  *Mob
	stmt *mql.FeelingStatement
}

func (s *FeelingExpression) Execute() {
	me := ""
	you := ""
	switch s.stmt.Ident {
	case "bow":
		if s.stmt.Object == "" {
			me = "You bow.\n"
			you = s.mob.Name + " bows.\n"
		} else {
			me = fmt.Sprintf("You bow to %s.\n", s.stmt.Object)
			you = fmt.Sprintf("%s bows to %s.\n", s.mob.Name, s.stmt.Object)
		}
	case "laugh":
		if s.stmt.Object == "" {
			me = "You fall down laughing.\n"
			you = s.mob.Name + " falls down laughing.\n"
		} else {
			me = fmt.Sprintf("You laugh at %s.\n", s.stmt.Object)
			you = fmt.Sprintf("%s laughs at %s.\n", s.mob.Name, s.stmt.Object)
		}
	default:
		s.mob.myMessageToChannel("what?\n")
	}

	if me != "" {
		s.mob.myMessageToChannel(me)
	}

	room := WorldInstance.getRoom(s.mob.CurrentRoom)
	for _, mob := range room.MobMap {
		if s.mob.Slug == mob.Slug {
			continue
		}
		mob.yourMessageToChannel(you)
	}
}
