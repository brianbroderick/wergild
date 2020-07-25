package mud

import (
	"github.com/brianbroderick/wergild/internal/mql"
)

//******
// SAY
//******

// SayExpression allows you to converse with others in the same room
type SayExpression struct {
	mob  *Mob
	stmt *mql.SayStatement
	// sentence string
	// token    mql.Token
	// object   string
}

func (s *SayExpression) Execute() {
	if s.stmt.Sentence == "" {
		s.mob.myMessageToChannel("You open your mouth to speak, but nothing comes out.\n")
		return
	}

	room := WorldInstance.getRoom(s.mob.CurrentRoom)

	toObj := ""
	toYou := ""
	yourDescriptor := ""
	myDescriptor := ""
	switch s.stmt.Token {
	case mql.SENTENCE:
		yourDescriptor = s.mob.Name + " says"
		myDescriptor = "You say"
		if s.stmt.Object != "" {
			toObj = " to " + s.stmt.Object
			toYou = " to you"
		}
	case mql.EXCLAIM:
		yourDescriptor = s.mob.Name + " exclaims"
		myDescriptor = "You exclaim"
		if s.stmt.Object != "" {
			toObj = " to " + s.stmt.Object
			toYou = " to you"
		}
	case mql.QUESTION:
		yourDescriptor = s.mob.Name + " asks"
		myDescriptor = "You ask"
		if s.stmt.Object != "" {
			toObj = " " + s.stmt.Object
			toYou = " you"
		}
	}

	myStr := myDescriptor + toObj + ": " + s.stmt.Sentence + "\n"
	s.mob.myMessageToChannel(myStr)

	yourStr := ""

	for _, mob := range room.MobMap {
		if s.mob.Slug == mob.Slug {
			continue
		}

		if mob.Slug == s.stmt.Object {
			yourStr = yourDescriptor + toYou + ": " + s.stmt.Sentence + "\n"
		} else {
			yourStr = yourDescriptor + toObj + ": " + s.stmt.Sentence + "\n"
		}
		mob.yourMessageToChannel(yourStr)
	}
}

//******
// TELL
//******

// TellExpression allows you to converse with others in the same room
type TellExpression struct {
	mob  *Mob
	stmt *mql.TellStatement
	// sentence string
	// token    mql.Token
	// object   string // mob telling something to
}

func (s *TellExpression) Execute() {
	if s.stmt.Sentence == "" {
		s.mob.myMessageToChannel("You're too distracted and nothing comes out.\n")
		return
	}

	if s.mob.Ap < 5 {
		s.mob.myMessageToChannel("You are too tired to concentrate.\n")
		return
	} else {
		s.mob.Ap -= 5
	}

	if obj, ok := ServerInstance.mobs[s.stmt.Object]; ok {
		yourDescriptor := s.mob.Name + " tells you: "
		myDescriptor := "You tell " + s.stmt.Object + ": "

		myStr := myDescriptor + s.stmt.Sentence + "\n"
		s.mob.myMessageToChannel(myStr)

		yourStr := yourDescriptor + s.stmt.Sentence + "\n"
		obj.yourMessageToChannel(yourStr)
	} else {
		s.mob.myMessageToChannel("You close your eyes and concentrate, but " + s.stmt.Object + " could not be found in this plane of existence.\n")
	}
}

//******
// SHOUT
//******

// ShoutExpression allows you to converse with others in the same room
type ShoutExpression struct {
	mob  *Mob
	stmt *mql.ShoutStatement
	// sentence string
	// token    mql.Token
}

func (s *ShoutExpression) Execute() {
	if s.stmt.Sentence == "" {
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

	myStr := myDescriptor + s.stmt.Sentence + "\n"
	s.mob.myMessageToChannel(myStr)

	yourStr := yourDescriptor + s.stmt.Sentence + "\n"
	for _, mob := range ServerInstance.mobs {
		if s.mob.Slug != mob.Slug {
			mob.yourMessageToChannel(yourStr)
		}
	}
}

//******
// EMOTE
//******

// EmoteExpression allows you to converse with others in the same room
type EmoteExpression struct {
	mob  *Mob
	stmt *mql.EmoteStatement
	// sentence string
	// token    mql.Token
}

func (s *EmoteExpression) Execute() {
	if s.stmt.Sentence == "" {
		s.mob.myMessageToChannel("You're too distracted and nothing happens.\n")
		return
	}

	yourDescriptor := s.mob.Name + " "
	myDescriptor := "You "

	myStr := myDescriptor + s.stmt.Sentence + "\n"
	s.mob.myMessageToChannel(myStr)

	room := WorldInstance.getRoom(s.mob.CurrentRoom)

	youStr := yourDescriptor + s.stmt.Sentence + "\n"

	for _, mob := range room.MobMap {
		if s.mob.Slug == mob.Slug {
			continue
		}
		mob.yourMessageToChannel(youStr)
	}
}
