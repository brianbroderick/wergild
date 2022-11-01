package lexer

import (
	"io"
	"strings"

	"github.com/brianbroderick/wergild/internal/token"
)

type Lexer struct {
	r       io.RuneScanner
	lastPos Pos
	pos     Pos
	ch      rune
	eof     bool // true if reader has ever seen eof.
}

// Pos specifies the line and character position of a token.
// The Char and Line are both zero-based indexes.
type Pos struct {
	Line int
	Char int
}

// eof is a marker code to signify that the reader can't read any more.
const eof = rune(0)
const eol = '\n'

func New(input string) *Lexer {
	l := &Lexer{r: strings.NewReader(input)}
	return l
}

func (l *Lexer) Scan() (tok token.Token, pos Pos) {
	l.skipWhitespace()

	switch l.ch {
	case '+':
		tok = newToken(token.PLUS, l.ch)
	}

	return tok, l.pos
}

// func (l *Lexer) Scan() (tok token.Token, pos Pos) {
// 	// l.scanWhitespace()

// 	// ch0, pos := l.read()
// }

func newToken(tokenType token.TokenType, ch rune) token.Token {
	return token.Token{Type: tokenType, Lit: string(ch)}
}

func (l *Lexer) read() {
	var err error
	l.ch, _, err = l.r.ReadRune()
	if err != nil {
		l.ch = eof
	}

	l.lastPos.Char = l.pos.Char
	l.lastPos.Line = l.pos.Line

	// Update position
	// Only count EOF once.
	if l.ch == eol {
		l.pos.Line++
		l.pos.Char = 0
	} else if !l.eof {
		l.pos.Char++
	}

	if l.ch == eof {
		l.eof = true
	}
}

func (l *Lexer) unread() {
	l.r.UnreadRune()
	l.pos.Char = l.lastPos.Char
	l.pos.Line = l.lastPos.Line
}

func (l *Lexer) peek() rune {
	ch, _, err := l.r.ReadRune()
	if err != nil {
		ch = eof
	}
	l.r.UnreadRune()
	return ch
}

func (l *Lexer) skipWhitespace() {
	for {
		l.read()
		if l.ch == eof {
			break
		} else if !isWhitespace(l.ch) {
			l.unread()
			break
		}
	}
}

// func (l *Lexer) scanIdent() {
// 	// Save the starting position of the identifier
// 	pos := l.pos

// 	for isIdentChar(l.ch) {

// 	}
// }

// isWhitespace returns true if the rune is a space, tab, or newline.
func isWhitespace(ch rune) bool { return ch == ' ' || ch == '\t' || ch == eol || ch == '\r' }

// isLetter returns true if the rune is a letter.
func isLetter(ch rune) bool { return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') }

// isDigit returns true if the rune is a digit.
func isDigit(ch rune) bool { return (ch >= '0' && ch <= '9') }

// isIdentChar returns true if the rune can be used in an unquoted identifier.
func isIdentChar(ch rune) bool { return isLetter(ch) || isDigit(ch) || ch == '_' }
