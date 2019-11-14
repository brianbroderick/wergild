package mud

import "bytes"

import "fmt"

// FeelingStatement represents a command for looking at a room or object.
type FeelingStatement struct {
	mob    *Mob
	ident  string // feeling. If not found, respond with something like "what?"
	object string
}

func (s *FeelingStatement) execute() {
	me := ""
	you := ""
	switch s.ident {
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

// parseFeelingStatement parses a feeling command and returns a Statement AST object.
// This function assumes the IDENT token has already been consumed.
// TODO add target for feelings
func (p *Parser) parseFeelingStatement() (*FeelingStatement, error) {
	p.Unscan()
	stmt := &FeelingStatement{}
	var err error

	stmt.ident, err = p.ParseIdent()
	if err != nil {
		return nil, err
	}

	tok, _, _ := p.ScanIgnoreWhitespace()
	switch tok {
	case AT, TO:
		obj, err := p.ParseIdent()
		if err != nil {
			return nil, err
		}
		stmt.object = obj
	}
	return stmt, nil
}

func (s *FeelingStatement) String() string {
	var buf bytes.Buffer

	if s.ident != "" {
		_, _ = buf.WriteString(s.ident)
	}

	return buf.String()
}

func (s *FeelingStatement) setMob(mob *Mob) {
	s.mob = mob
}
