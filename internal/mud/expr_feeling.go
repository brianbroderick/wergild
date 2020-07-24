package mud

import "fmt"

// FeelingExpression represents a command for looking at a room or object.
type FeelingExpression struct {
	mob    *Mob
	ident  string // feeling. If not found, respond with something like "what?"
	object string
}

func (s *FeelingExpression) Execute() {
	me := ""
	you := ""
	switch s.ident {
	case "bow":
		if s.object == "" {
			me = "You bow.\n"
			you = s.mob.Name + " bows.\n"
		} else {
			me = fmt.Sprintf("You bow to %s.\n", s.object)
			you = fmt.Sprintf("%s bows to %s.\n", s.mob.Name, s.object)
		}
	case "laugh":
		if s.object == "" {
			me = "You fall down laughing.\n"
			you = s.mob.Name + " falls down laughing.\n"
		} else {
			me = fmt.Sprintf("You laugh at %s.\n", s.object)
			you = fmt.Sprintf("%s laughs at %s.\n", s.mob.Name, s.object)
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
