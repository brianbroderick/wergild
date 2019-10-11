package main

import "strings"

// Token represents a lexical token.
type Token int

const (
	// Special tokens
	ILLEGAL Token = iota
	EOF
	WS
	NIL
	COMMENT

	literalBeg
	// Literals
	IDENT       // main
	BOUNDPARAM  // $param
	NUMBER      // 12345.67
	INTEGER     // 12345
	DURATIONVAL // 13h
	STRING      // "abc"
	BADSTRING   // "abc
	BADESCAPE   // \q
	TRUE        // true
	FALSE       // false
	REGEX       // Regular expressions
	BADREGEX    // `.*
	literalEnd

	operatorBeg
	// ADD and the following are InfluxQL Operators
	ADD         // +
	SUB         // -
	MUL         // *
	DIV         // /
	MOD         // %
	BITWISE_AND // &
	BITWISE_OR  // |
	BITWISE_XOR // ^

	AND // AND
	OR  // OR

	EQ       // =
	NEQ      // !=
	EQREGEX  // =~
	NEQREGEX // !~
	LT       // <
	LTE      // <=
	GT       // >
	GTE      // >=
	operatorEnd

	// Misc characters
	LPAREN      // (
	RPAREN      // )
	COMMA       // ,
	COLON       // :
	DOUBLECOLON // ::
	SEMICOLON   // ;
	DOT         // .

	aliasBeg
	L // LOOK
	N // NORTH
	S // SOUTH
	E // EAST
	W // WEST
	U // UP
	D // DOWN
	aliasEnd

	keywordBeg
	// Sense Keywords
	LOOK

	AT
	ON
	IN

	// Movement Keywords
	NORTH
	SOUTH
	EAST
	WEST
	UP
	DOWN

	// Admin Keywords
	NICK // nickname i.e. alias mapping
	keywordEnd
)

// These are how a string is mapped to the token
var tokens = [...]string{
	ILLEGAL: "ILLEGAL",
	EOF:     "EOF",
	WS:      "WS",
	NIL:     "NIL",

	IDENT:       "IDENT",
	NUMBER:      "NUMBER",
	DURATIONVAL: "DURATIONVAL",
	STRING:      "STRING",
	BADSTRING:   "BADSTRING",
	BADESCAPE:   "BADESCAPE",
	TRUE:        "TRUE",
	FALSE:       "FALSE",
	REGEX:       "REGEX",

	ADD:         "+",
	SUB:         "-",
	MUL:         "*",
	DIV:         "/",
	MOD:         "%",
	BITWISE_AND: "&",
	BITWISE_OR:  "|",
	BITWISE_XOR: "^",

	AND: "AND",
	OR:  "OR",

	EQ:       "=",
	NEQ:      "!=",
	EQREGEX:  "=~",
	NEQREGEX: "!~",
	LT:       "<",
	LTE:      "<=",
	GT:       ">",
	GTE:      ">=",

	LPAREN:      "(",
	RPAREN:      ")",
	COMMA:       ",",
	COLON:       ":",
	DOUBLECOLON: "::",
	SEMICOLON:   ";",
	DOT:         ".",

	LOOK: "LOOK",
	AT:   "AT",
	ON:   "ON",
	IN:   "IN",

	// Movement Keywords:
	NORTH: "NORTH",
	SOUTH: "SOUTH",
	EAST:  "EAST",
	WEST:  "WEST",
	UP:    "UP",
	DOWN:  "DOWN",

	// Admin Keywords
	NICK: "NICK",

	// Aliases
	N: "N",
	S: "S",
	E: "E",
	W: "W",
	U: "U",
	D: "D",
	L: "L",
}

// These are how globalAlias tokens are resolved
var convertToTokens = [...]Token{
	N: NORTH,
	S: SOUTH,
	E: EAST,
	W: WEST,
	U: UP,
	D: DOWN,
	L: LOOK,
}

var keywords map[string]Token
var globalAliases map[string]Token

func init() {
	keywords = make(map[string]Token)
	for tok := keywordBeg + 1; tok < keywordEnd; tok++ {
		keywords[strings.ToLower(tokens[tok])] = tok
	}
	for _, tok := range []Token{AND, OR} {
		keywords[strings.ToLower(tokens[tok])] = tok
	}
	keywords["true"] = TRUE
	keywords["false"] = FALSE

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

// Precedence returns the operator precedence of the binary operator token.
func (tok Token) Precedence() int {
	switch tok {
	case OR:
		return 1
	case AND:
		return 2
	case EQ, NEQ, EQREGEX, NEQREGEX, LT, LTE, GT, GTE:
		return 3
	case ADD, SUB, BITWISE_OR, BITWISE_XOR:
		return 4
	case MUL, DIV, MOD, BITWISE_AND:
		return 5
	}
	return 0
}

// isOperator returns true for operator tokens.
func (tok Token) isOperator() bool { return tok > operatorBeg && tok < operatorEnd }

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
