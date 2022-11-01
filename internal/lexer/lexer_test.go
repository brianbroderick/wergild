package lexer

import (
	"testing"

	"github.com/brianbroderick/wergild/internal/token"
	"github.com/stretchr/testify/assert"
)

func TestCodeScan(t *testing.T) {
	input := `
	let five = 5;
	let ten = 10;
	
	let add = fn(x, y) {
		x + y;
	};
	
	let result = add(five, ten);
	
	if (5 < 10) {
		return true;
	} else {
		return false;
	}	
	`

	tests := []token.Token{
		{Type: token.LET, Lit: "let"},
		{Type: token.IDENT, Lit: "five"},
		{Type: token.ASSIGN, Lit: "="},
		{Type: token.INT, Lit: "5"},
		{Type: token.SEMICOLON, Lit: ";"},
		{Type: token.LET, Lit: "let"},
		{Type: token.IDENT, Lit: "ten"},
		{Type: token.ASSIGN, Lit: "="},
		{Type: token.INT, Lit: "10"},
		{Type: token.SEMICOLON, Lit: ";"},
		{Type: token.LET, Lit: "let"},
		{Type: token.IDENT, Lit: "add"},
		{Type: token.ASSIGN, Lit: "="},
		{Type: token.FUNCTION, Lit: "fn"},
		{Type: token.LPAREN, Lit: "("},
		{Type: token.IDENT, Lit: "x"},
		{Type: token.COMMA, Lit: ","},
		{Type: token.IDENT, Lit: "y"},
		{Type: token.RPAREN, Lit: ")"},
		{Type: token.LBRACE, Lit: "{"},
		{Type: token.IDENT, Lit: "x"},
		{Type: token.PLUS, Lit: "+"},
		{Type: token.IDENT, Lit: "y"},
		{Type: token.SEMICOLON, Lit: ";"},
		{Type: token.RBRACE, Lit: "}"},
		{Type: token.SEMICOLON, Lit: ";"},
		{Type: token.LET, Lit: "let"},
		{Type: token.IDENT, Lit: "result"},
		{Type: token.ASSIGN, Lit: "="},
		{Type: token.IDENT, Lit: "add"},
		{Type: token.LPAREN, Lit: "("},
		{Type: token.IDENT, Lit: "five"},
		{Type: token.COMMA, Lit: ","},
		{Type: token.IDENT, Lit: "ten"},
		{Type: token.RPAREN, Lit: ")"},
		{Type: token.SEMICOLON, Lit: ";"},
		{Type: token.IF, Lit: "if"},
		{Type: token.LPAREN, Lit: "("},
		{Type: token.INT, Lit: "5"},
		{Type: token.LT, Lit: "<"},
		{Type: token.INT, Lit: "10"},
		{Type: token.RPAREN, Lit: ")"},
		{Type: token.LBRACE, Lit: "{"},
		{Type: token.RETURN, Lit: "return"},
		{Type: token.TRUE, Lit: "true"},
		{Type: token.SEMICOLON, Lit: ";"},
		{Type: token.RBRACE, Lit: "}"},
		{Type: token.ELSE, Lit: "else"},
		{Type: token.LBRACE, Lit: "{"},
		{Type: token.RETURN, Lit: "return"},
		{Type: token.FALSE, Lit: "false"},
		{Type: token.SEMICOLON, Lit: ";"},
		{Type: token.RBRACE, Lit: "}"},
		{Type: token.EOF, Lit: ""},
	}

	l := New(input)

	for _, tt := range tests {
		tok, _ := l.Scan()

		assert.Equal(t, token.Tokens[tt.Type], token.Tokens[tok.Type])
		assert.Equal(t, tt.Lit, tok.Lit)
	}
}

func TestSimpleScan(t *testing.T) {
	input := ` + - * < >
	; , { } ( ) = == ! != / foo 12345 `

	tests := []token.Token{
		{Type: token.PLUS, Lit: "+"},
		{Type: token.MINUS, Lit: "-"},
		{Type: token.ASTERISK, Lit: "*"},
		{Type: token.LT, Lit: "<"},
		{Type: token.GT, Lit: ">"},
		{Type: token.SEMICOLON, Lit: ";"},
		{Type: token.COMMA, Lit: ","},
		{Type: token.LBRACE, Lit: "{"},
		{Type: token.RBRACE, Lit: "}"},
		{Type: token.LPAREN, Lit: "("},
		{Type: token.RPAREN, Lit: ")"},
		{Type: token.ASSIGN, Lit: "="},
		{Type: token.EQ, Lit: "=="},
		{Type: token.BANG, Lit: "!"},
		{Type: token.NOT_EQ, Lit: "!="},
		{Type: token.SLASH, Lit: "/"},
		{Type: token.IDENT, Lit: "foo"},
		{Type: token.INT, Lit: "12345"},
		{Type: token.EOF, Lit: ""},
	}

	l := New(input)

	for _, tt := range tests {
		tok, _ := l.Scan()

		assert.Equal(t, token.Tokens[tt.Type], token.Tokens[tok.Type])
		assert.Equal(t, tt.Lit, tok.Lit)
	}
}
