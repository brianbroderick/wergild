package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Ensure the parser can parse strings into Statement ASTs.
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
			obj: &LookStatement{token: AT, ident: "chair"},
			p:   `LOOK AT chair`,
		},

		{
			s:   `look`,
			obj: &LookStatement{token: EOF, ident: ""},
			p:   `LOOK`,
		},

		{
			s:   `look north`,
			obj: &LookStatement{token: NORTH, ident: ""},
			p:   `LOOK NORTH`,
		},

		{
			s:   `look up`,
			obj: &LookStatement{token: UP, ident: ""},
			p:   `LOOK UP`,
		},

		// check alias
		{
			s:   `l n`,
			obj: &LookStatement{token: NORTH, ident: ""},
			p:   `LOOK NORTH`,
		},

		{
			s:   `quit`,
			obj: &QuitStatement{},
			p:   `QUIT`,
		},

		{
			s:   `north`,
			obj: &DirectionStatement{token: NORTH},
			p:   `NORTH`,
		},

		{
			s:   `south`,
			obj: &DirectionStatement{token: SOUTH},
			p:   `SOUTH`,
		},

		// // Errors
		// {s: `foo`, err: `found "foo", expected SELECT`},
		// {s: `SELECT !`, err: `found "!", expected field`},
		// {s: `SELECT field xxx`, err: `found "xxx", expected FROM`},
		// {s: `SELECT field FROM *`, err: `found "*", expected table name`},
	}

	for _, tt := range tests {
		obj, err := NewParser(strings.NewReader(tt.s)).ParseStatement()
		assert.NoError(t, err)
		assert.Equal(t, tt.obj, obj)
		assert.Equal(t, tt.p, tt.obj.String())
	}
}
