package main

import (
	"os"
)

func init() {
	os.Setenv("PLATFORM_ENV", "test")
}

// // Look for a lens that doesn't exist
// func TestLook(t *testing.T) {
// 	var tests = []struct {
// 		s   string
// 		tok Token
// 		lit string
// 	}{
// 		// Special tokens (EOF, ILLEGAL, WS)
// 		{s: ``, tok: EOF},
// 		{s: `#`, tok: ILLEGAL, lit: `#`},
// 		{s: ` `, tok: WS, lit: " "},
// 		{s: "\t", tok: WS, lit: "\t"},
// 		{s: "\n", tok: WS, lit: "\n"},

// 		// Misc characters
// 		{s: `*`, tok: ASTERISK, lit: "*"},

// 		// Identifiers
// 		{s: `fireplace`, tok: IDENT, lit: `fireplace`},
// 		{s: `chair`, tok: IDENT, lit: `chair`},

// 		// Keywords
// 		{s: `look at chair`, tok: LOOK, lit: "look"},
// 	}

// 	for i, tt := range tests {
// 		s := NewScanner(strings.NewReader(tt.s))
// 		tok, lit := s.Scan()
// 		if tt.tok != tok {
// 			t.Errorf("%d. %q token mismatch: exp=%q got=%q <%q>", i, tt.s, tt.tok, tok, lit)
// 		} else if tt.lit != lit {
// 			t.Errorf("%d. %q literal mismatch: exp=%q got=%q", i, tt.s, tt.lit, lit)
// 		}
// 	}

// 	// str := "look at chair"

// 	// s := NewScanner(strings.NewReader(str))
// 	// tok, lit := s.Scan()

// 	//assert.Equal(t, LOOK, tok)
// 	// assert.Equal(t, "look", lit)
// }
