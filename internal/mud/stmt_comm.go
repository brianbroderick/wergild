package mud

import (
	"bytes"
)

// SayStatement represents a command for exiting the game.
type SayStatement struct {
	mob      *Mob
	sentence string
	token    Token
}

func (s *SayStatement) String() string {
	var buf bytes.Buffer
	_, _ = buf.WriteString("SAY")

	return buf.String()
}

// parseSayStatement parses a look command and returns a Statement AST object.
// This function assumes the QUIT token has already been consumed.
func (p *Parser) parseSayStatement() (*SayStatement, error) {
	stmt := &SayStatement{}
	stmt.token, _, stmt.sentence = p.ScanSentence()

	return stmt, nil
}

func (s *SayStatement) execute() {
	room := WorldInstance.getRoom(s.mob.CurrentRoom)

	yourDescriptor := ""
	myDescriptor := ""
	switch s.token {
	case SENTENCE:
		yourDescriptor = s.mob.Name + " says: "
		myDescriptor = "You say: "
	case EXCLAIM:
		yourDescriptor = s.mob.Name + " exclaims: "
		myDescriptor = "You exclaim: "
	case QUESTION:
		yourDescriptor = s.mob.Name + " asks: "
		myDescriptor = "You ask: "
	}

	if s.sentence == "" {
		s.mob.myMessageToChannel("You open your mouth to speak, but nothing comes out.\n")
		return
	}

	myStr := myDescriptor + s.sentence + "\n"
	s.mob.myMessageToChannel(myStr)

	yourStr := yourDescriptor + s.sentence + "\n"
	for _, mob := range room.MobMap {
		if s.mob.Slug != mob.Slug {
			mob.yourMessageToChannel(yourStr)
		}
	}
}

func (s *SayStatement) setMob(mob *Mob) {
	s.mob = mob
}
