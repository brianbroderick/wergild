package ast

import (
	"testing"

	"github.com/brianbroderick/wergild/internal/token"
)

func TestString(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&LetStatement{
				Token: token.Token{Type: token.LET, Lit: "let"},
				Name: &Identifier{
					Token: token.Token{Type: token.IDENT, Lit: "myVar"},
					Value: "myVar",
				},
				Value: &Identifier{
					Token: token.Token{Type: token.IDENT, Lit: "anotherVar"},
					Value: "anotherVar",
				},
			},
		},
	}

	if program.String() != "let myVar = anotherVar;" {
		t.Errorf("program.String() wrong. got=%q", program.String())
	}
}
