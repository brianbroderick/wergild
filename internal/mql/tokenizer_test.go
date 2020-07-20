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

func TestTokenizer(t *testing.T) {
	type token = struct {
		tok Token
		lit string
	}

	var tests = []struct {
		s      string
		tokens []token
	}{
		{
			s: "look north",
			tokens: []token{
				{tok: LOOK},
				{tok: WS, lit: " "},
				{tok: NORTH}},
		},
		{
			s: "look at treasure",
			tokens: []token{
				{tok: LOOK},
				{tok: WS, lit: " "},
				{tok: AT},
				{tok: WS, lit: " "},
				{tok: IDENT, lit: "treasure"},
			},
		},
		{
			s: "say hi",
			tokens: []token{
				{tok: SAY},
				{tok: WS, lit: " "},
				{tok: IDENT, lit: "hi"}},
		},
	}

	for _, tt := range tests {
		tokenizer := newTokenizer(strings.NewReader(tt.s))
		tok := ILLEGAL
		var lit string
		var s []token
		for tok != EOF {
			tok, _, lit = tokenizer.Scan()
			if tok != EOF {
				s = append(s, token{tok: tok, lit: lit})
			}
		}

		assert.Equal(t, tt.tokens, s)
	}

}

func TestLookNorth(t *testing.T) {
	str := "look north"
	tokenizer := newTokenizer(strings.NewReader(str))

	tok, _, lit := tokenizer.Scan()
	assert.Equal(t, LOOK, tok)
	assert.Equal(t, "", lit)

	tok, _, lit = tokenizer.Scan()
	assert.Equal(t, WS, tok)
	assert.Equal(t, " ", lit)

	tok, _, lit = tokenizer.Scan()
	assert.Equal(t, NORTH, tok)
	assert.Equal(t, "", lit)
}

func TestSayHi(t *testing.T) {
	str := "say hi"
	tokenizer := newTokenizer(strings.NewReader(str))

	tok, _, lit := tokenizer.Scan()
	assert.Equal(t, SAY, tok)
	assert.Equal(t, "", lit)

	tok, _, lit = tokenizer.Scan()
	assert.Equal(t, WS, tok)
	assert.Equal(t, " ", lit)

	tok, _, lit = tokenizer.Scan()
	assert.Equal(t, IDENT, tok)
	assert.Equal(t, "hi", lit)
}
