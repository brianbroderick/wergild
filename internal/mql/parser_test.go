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

func TestParser_ParseStatement(t *testing.T) {
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
