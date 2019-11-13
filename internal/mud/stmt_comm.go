package mud

import (
	"bytes"
)

//******
// SAY
//******

// SayStatement allows you to converse with others in the same room
type SayStatement struct {
	mob      *Mob
	sentence string
	token    Token
}

// TODO: Finish writing options
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
	if s.sentence == "" {
		s.mob.myMessageToChannel("You open your mouth to speak, but nothing comes out.\n")
		return
	}

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

//******
// TELL
//******

// TellStatement allows you to converse with others in the same room
type TellStatement struct {
	mob      *Mob
	sentence string
	token    Token
	object   string // mob telling something to
}

// TODO: Finish writing options
func (s *TellStatement) String() string {
	var buf bytes.Buffer
	_, _ = buf.WriteString("TELL")

	return buf.String()
}

// parseTellStatement parses a look command and returns a Statement AST object.
// This function assumes the QUIT token has already been consumed.
func (p *Parser) parseTellStatement() (*TellStatement, error) {
	stmt := &TellStatement{}
	tok, _, lit := p.ScanIgnoreWhitespace()
	switch tok {
	case IDENT:
		stmt.object = lit
	case EOF:
		return stmt, nil
	}

	stmt.token, _, stmt.sentence = p.ScanSentence()

	return stmt, nil
}

func (s *TellStatement) execute() {
	if s.sentence == "" {
		s.mob.myMessageToChannel("You're too distracted and nothing comes out.\n")
		return
	}

	if s.mob.Ap < 5 {
		s.mob.myMessageToChannel("You are too tired to concentrate.\n")
		return
	} else {
		s.mob.Ap -= 5
	}

	if obj, ok := ServerInstance.mobs[s.object]; ok {
		yourDescriptor := s.mob.Name + " tells you: "
		myDescriptor := "You tell " + s.object + ": "

		myStr := myDescriptor + s.sentence + "\n"
		s.mob.myMessageToChannel(myStr)

		yourStr := yourDescriptor + s.sentence + "\n"
		obj.yourMessageToChannel(yourStr)
	} else {
		s.mob.myMessageToChannel("You close your eyes and concentrate, but " + s.object + " could not be found in this plane of existence.\n")
	}
}

func (s *TellStatement) setMob(mob *Mob) {
	s.mob = mob
}

//******
// SHOUT
//******

// ShoutStatement allows you to converse with others in the same room
type ShoutStatement struct {
	mob      *Mob
	sentence string
	token    Token
}

// TODO: Finish writing options
func (s *ShoutStatement) String() string {
	var buf bytes.Buffer
	_, _ = buf.WriteString("SHOUT")

	return buf.String()
}

// parseShoutStatement parses a look command and returns a Statement AST object.
// This function assumes the QUIT token has already been consumed.
func (p *Parser) parseShoutStatement() (*ShoutStatement, error) {
	stmt := &ShoutStatement{}
	stmt.token, _, stmt.sentence = p.ScanSentence()

	return stmt, nil
}

func (s *ShoutStatement) execute() {
	if s.sentence == "" {
		s.mob.myMessageToChannel("You're too distracted and nothing comes out.\n")
		return
	}

	if s.mob.Ap < 50 {
		s.mob.myMessageToChannel("You are too tired to concentrate on shouting to the world.\n")
		return
	} else {
		s.mob.Ap -= 50
	}

	yourDescriptor := s.mob.Name + " shouts: "
	myDescriptor := "You Shout: "

	myStr := myDescriptor + s.sentence + "\n"
	s.mob.myMessageToChannel(myStr)

	yourStr := yourDescriptor + s.sentence + "\n"
	for _, mob := range ServerInstance.mobs {
		if s.mob.Slug != mob.Slug {
			mob.yourMessageToChannel(yourStr)
		}
	}
}

func (s *ShoutStatement) setMob(mob *Mob) {
	s.mob = mob
}
