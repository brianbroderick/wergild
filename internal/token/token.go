package token

import "strings"

type TokenType int

type Token struct {
	Type TokenType
	Lit  string
}

const (
	ILLEGAL TokenType = iota
	EOF
	WS // whitespace
	NIL

	literalBeg // Literals
	IDENT      // identity: add, foobar, x, y, my_var, ...
	INT        // 12345
	STRING     // "foobar"
	literalEnd

	// Operators
	ASSIGN   // =
	PLUS     // +
	MINUS    // -
	BANG     // !
	ASTERISK // *
	SLASH    // /

	LT // <
	GT // >

	EQ     // ==
	NOT_EQ // !=

	// Delimiters
	COMMA     // ,
	SEMICOLON // ;
	COLON     // :

	LPAREN   // (
	RPAREN   // )
	LBRACE   // {
	RBRACE   // }
	LBRACKET // [
	RBRACKET // ]

	// Keywords
	keywordBeg
	FUNCTION
	LET
	TRUE
	FALSE
	IF
	ELSE
	RETURN
	keywordEnd
)

// These are how a string is mapped to the token
var Tokens = [...]string{
	ILLEGAL: "ILLEGAL",
	EOF:     "EOF",
	WS:      "WS",
	NIL:     "NIL",

	IDENT:  "IDENT",
	INT:    "INTEGER",
	STRING: "STRING",

	ASSIGN:   "ASSIGN '='",
	PLUS:     "PLUS '+'",
	MINUS:    "MINUS '-'",
	BANG:     "BANG '!'",
	ASTERISK: "ASTERISK '*'",
	SLASH:    "SLASH '/'",

	LT: "LT '<'",
	GT: "GT '>'",

	EQ:     "EQ '=='",
	NOT_EQ: "NOT_EQ '!='",

	// Delimiters
	COMMA:     "COMMA ','",
	SEMICOLON: "SEMICOLON ';'",
	COLON:     "COLON ':'",

	LPAREN:   "LPAREN '('",
	RPAREN:   "RPAREN ')'",
	LBRACE:   "LBRACE '{'",
	RBRACE:   "RBRACE '}'",
	LBRACKET: "LBRACKET '['",
	RBRACKET: "RBRACKET ']'",

	// Keywords
	FUNCTION: "FN",
	LET:      "LET",
	TRUE:     "TRUE",
	FALSE:    "FALSE",
	IF:       "IF",
	ELSE:     "ELSE",
	RETURN:   "RETURN",
}

var keywords map[string]TokenType

func init() {
	keywords = make(map[string]TokenType)
	for tok := keywordBeg + 1; tok < keywordEnd; tok++ {
		keywords[strings.ToLower(Tokens[tok])] = tok
	}
}

// String returns the string representation of the token.
func (tok TokenType) String() string {
	if tok >= 0 && tok < TokenType(len(Tokens)) {
		return Tokens[tok]
	}
	return ""
}

// Lookup returns the token associated with a given string.
func Lookup(ident string) TokenType {
	if tok, ok := keywords[strings.ToLower(ident)]; ok {
		return tok
	}

	return IDENT
}
