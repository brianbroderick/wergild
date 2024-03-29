package mql

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func init() {
	os.Setenv("PLATFORM_ENV", "test")
}

func TestParserParseStatement(t *testing.T) {
	var tests = []struct {
		s   string
		obj Statement
		p   string
		err string
	}{
		// Single field statement
		{
			s:   `look at chair`,
			obj: &LookStatement{Token: AT, Ident: "chair"},
			p:   `LOOK AT chair`,
		},

		{
			s:   `look at "bob's chair"`,
			obj: &LookStatement{Token: AT, Ident: "bob's chair"},
			p:   `LOOK AT bob's chair`,
		},

		{
			s:   `look`,
			obj: &LookStatement{Token: EOF, Ident: ""},
			p:   `LOOK`,
		},

		{
			s:   `look north`,
			obj: &LookStatement{Token: NORTH, Ident: ""},
			p:   `LOOK NORTH`,
		},

		{
			s:   `look up`,
			obj: &LookStatement{Token: UP, Ident: ""},
			p:   `LOOK UP`,
		},

		{
			s:   `smell`,
			obj: &SmellStatement{Token: EOF},
			p:   `SMELL`,
		},

		{
			s:   `smell north`,
			obj: &SmellStatement{Token: NORTH},
			p:   `SMELL NORTH`,
		},

		{
			s:   `smell dog`,
			obj: &SmellStatement{Token: IDENT, Ident: "dog"},
			p:   `SMELL dog`,
		},

		{
			s:   `listen`,
			obj: &ListenStatement{Token: EOF},
			p:   `LISTEN`,
		},

		{
			s:   `listen south`,
			obj: &ListenStatement{Token: SOUTH},
			p:   `LISTEN SOUTH`,
		},

		{
			s:   `listen to bob`,
			obj: &ListenStatement{Token: TO, Ident: "bob"},
			p:   `LISTEN TO bob`,
		},

		{
			s:   `touch`,
			obj: &TouchStatement{Token: EOF, Ident: ""},
			p:   `TOUCH`,
		},

		{
			s:   `touch bob`,
			obj: &TouchStatement{Token: IDENT, Ident: "bob"},
			p:   `TOUCH bob`,
		},

		{
			s:   `touch in dresser`,
			obj: &TouchStatement{Token: IN, Ident: "dresser"},
			p:   `TOUCH IN dresser`,
		},

		{
			s:   `touch west`,
			obj: &TouchStatement{Token: WEST},
			p:   `TOUCH WEST`,
		},

		{
			s:   "say hi there",
			obj: &SayStatement{Token: SENTENCE, Sentence: "hi there"},
			p:   "SAY",
		},

		{
			s:   "say hi",
			obj: &SayStatement{Token: SENTENCE, Sentence: "hi"},
			p:   "SAY",
		},

		{
			s:   "say to brian hi there",
			obj: &SayStatement{Token: SENTENCE, Sentence: "hi there", Object: "brian"},
			p:   "SAY",
		},

		{
			s:   "say How are you?",
			obj: &SayStatement{Token: QUESTION, Sentence: "How are you?"},
			p:   "SAY",
		},

		{
			s:   "say I am doing fine!",
			obj: &SayStatement{Token: EXCLAIM, Sentence: "I am doing fine!"},
			p:   "SAY",
		},

		{
			s:   "say I'm doing fine!",
			obj: &SayStatement{Token: EXCLAIM, Sentence: "I'm doing fine!"},
			p:   "SAY",
		},

		{
			s:   "tell brian hi there",
			obj: &TellStatement{Token: SENTENCE, Object: "brian", Sentence: "hi there"},
			p:   "TELL",
		},

		{
			s:   "tell brian watch out!",
			obj: &TellStatement{Token: EXCLAIM, Object: "brian", Sentence: "watch out!"},
			p:   "TELL",
		},

		{
			s:   "tell brian going my way?",
			obj: &TellStatement{Token: QUESTION, Object: "brian", Sentence: "going my way?"},
			p:   "TELL",
		},

		{
			s:   "tell brian hi",
			obj: &TellStatement{Token: SENTENCE, Object: "brian", Sentence: "hi"},
			p:   "TELL",
		},

		{
			s:   "tell brian hi!",
			obj: &TellStatement{Token: EXCLAIM, Object: "brian", Sentence: "hi!"},
			p:   "TELL",
		},

		{
			s:   "shout hello everyone",
			obj: &ShoutStatement{Token: SENTENCE, Sentence: "hello everyone"},
			p:   "SHOUT",
		},

		{
			s:   "shout The world is full of darkness!",
			obj: &ShoutStatement{Token: EXCLAIM, Sentence: "The world is full of darkness!"},
			p:   "SHOUT",
		},

		{
			s:   "shout Where is everyone?",
			obj: &ShoutStatement{Token: QUESTION, Sentence: "Where is everyone?"},
			p:   "SHOUT",
		},

		{
			s:   "emote stared at the roses.",
			obj: &EmoteStatement{Token: SENTENCE, Sentence: "stared at the roses."},
			p:   "EMOTE",
		},

		{
			s:   "emote wonders, where are you?",
			obj: &EmoteStatement{Token: QUESTION, Sentence: "wonders, where are you?"},
			p:   "EMOTE",
		},

		{
			s:   "emote thinks, I exist!",
			obj: &EmoteStatement{Token: EXCLAIM, Sentence: "thinks, I exist!"},
			p:   "EMOTE",
		},

		{
			s:   "north",
			obj: &DirectionStatement{Token: NORTH},
			p:   "NORTH",
		},

		{
			s:   "laugh",
			obj: &FeelingStatement{Ident: "laugh"},
			p:   "laugh",
		},

		{
			s:   "laugh at steve",
			obj: &FeelingStatement{Ident: "laugh", Object: "steve"},
			p:   "laugh",
		},

		{
			s:   "score",
			obj: &ScoreStatement{},
			p:   "SCORE",
		},

		{
			s:   "quit",
			obj: &QuitStatement{},
			p:   "QUIT",
		},

		{
			s:   "imagine room north forest",
			obj: &ImagineStatement{Object: ROOM, Direction: NORTH, Name: "forest"},
			p:   "IMAGINE",
		},

		// // check alias
		// {
		// 	s:   `l n`,
		// 	obj: &LookStatement{Token: NORTH, Ident: ""},
		// 	p:   `LOOK NORTH`,
		// },
	}

	for _, tt := range tests {
		obj, err := NewParser(strings.NewReader(tt.s)).ParseStatement()
		assert.NoError(t, err)
		assert.Equal(t, tt.obj, obj)
		assert.Equal(t, tt.p, tt.obj.String())
	}
}

func TestParserLoop(t *testing.T) {
	str := "10 look"
	p := "10"

	obj, err := NewParser(strings.NewReader(str)).ParseStatement()
	assert.NoError(t, err)
	assert.Equal(t, p, obj.String())
}
