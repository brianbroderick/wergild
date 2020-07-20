package mql

import "strings"

// ql_token.go contains a list of valid tokens to be used by
// the query language as well as their string representations.

type Token int

const (
	ILLEGAL Token = iota
	EOF
	WS // whitespace
	NIL

	// Communication
	SENTENCE
	QUESTION
	EXCLAIM

	literalBeg
	// Literals
	IDENT     // main
	STRING    // "abc"
	NUMBER    // 12345.67
	INTEGER   // 12345
	BADSTRING // "abc
	BADESCAPE // \q
	literalEnd

	// Misc characters
	LPAREN      // (
	RPAREN      // )
	COMMA       // ,
	COLON       // :
	APOSTROPHE  // '
	DOUBLECOLON // ::
	SEMICOLON   // ;
	DOT         // .

	aliasBeg
	L    // LOOK
	N    // NORTH
	S    // SOUTH
	E    // EAST
	W    // WEST
	U    // UP
	D    // DOWN
	EXIT // QUIT
	aliasEnd

	keywordBeg
	// Sense Keywords
	LOOK
	AT
	ON
	IN
	SMELL
	LISTEN
	TO
	TOUCH
	SAY
	TELL
	SHOUT
	EMOTE

	// Movement Keywords
	NORTH
	SOUTH
	EAST
	WEST
	UP
	DOWN

	// Admin Keywords
	NICK // nickname i.e. alias mapping
	QUIT

	// Player Keywords
	SCORE
	keywordEnd
)

// These are how a string is mapped to the token
var tokens = [...]string{
	ILLEGAL: "ILLEGAL",
	EOF:     "EOF",
	WS:      "WS",
	NIL:     "NIL",

	// Communication
	SENTENCE: "SENTENCE",
	QUESTION: "QUESTION",
	EXCLAIM:  "EXCLAIM",

	IDENT:     "IDENT",
	NUMBER:    "NUMBER",
	STRING:    "STRING",
	BADSTRING: "BADSTRING",
	BADESCAPE: "BADESCAPE",

	LPAREN:      "(",
	RPAREN:      ")",
	COMMA:       ",",
	COLON:       ":",
	DOUBLECOLON: "::",
	SEMICOLON:   ";",
	DOT:         ".",
	APOSTROPHE:  "'",

	LOOK:   "LOOK",
	AT:     "AT",
	ON:     "ON",
	IN:     "IN",
	SMELL:  "SMELL",
	LISTEN: "LISTEN",
	TO:     "TO",
	TOUCH:  "TOUCH",
	SAY:    "SAY",
	TELL:   "TELL",
	SHOUT:  "SHOUT",
	EMOTE:  "EMOTE",

	// Movement Keywords:
	NORTH: "NORTH",
	SOUTH: "SOUTH",
	EAST:  "EAST",
	WEST:  "WEST",
	UP:    "UP",
	DOWN:  "DOWN",

	// Admin Keywords
	NICK: "NICK",
	QUIT: "QUIT",

	// Aliases
	N:    "N",
	S:    "S",
	E:    "E",
	W:    "W",
	U:    "U",
	D:    "D",
	L:    "L",
	EXIT: "EXIT",

	SCORE: "SCORE",
}

// These are how globalAlias tokens are resolved
var convertToTokens = [...]Token{
	N:    NORTH,
	S:    SOUTH,
	E:    EAST,
	W:    WEST,
	U:    UP,
	D:    DOWN,
	L:    LOOK,
	EXIT: QUIT,
}

var keywords map[string]Token
var globalAliases map[string]Token

func init() {
	keywords = make(map[string]Token)
	for tok := keywordBeg + 1; tok < keywordEnd; tok++ {
		keywords[strings.ToLower(tokens[tok])] = tok
	}

	globalAliases = make(map[string]Token)
	for tok := aliasBeg + 1; tok < aliasEnd; tok++ {
		globalAliases[strings.ToLower(tokens[tok])] = tok
	}
}

// String returns the string representation of the token.
func (tok Token) String() string {
	if tok >= 0 && tok < Token(len(tokens)) {
		return tokens[tok]
	}
	return ""
}

// tokstr returns a literal if provided, otherwise returns the token string.
func tokstr(tok Token, lit string) string {
	if lit != "" {
		return lit
	}
	return tok.String()
}

// Lookup returns the token associated with a given string.
func Lookup(ident string) Token {
	if tok, ok := keywords[strings.ToLower(ident)]; ok {
		return tok
	}

	if tok, ok := globalAliases[strings.ToLower(ident)]; ok {
		return convertToTokens[tok]
	}

	return IDENT
}

// Pos specifies the line and character position of a token.
// The Char and Line are both zero-based indexes.
type Pos struct {
	Line int
	Char int
}
