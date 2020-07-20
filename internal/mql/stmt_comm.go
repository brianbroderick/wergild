package mql

import (
	"bytes"
)

//******
// SAY
//******

// SayStatement allows you to converse with others in the same room
type SayStatement struct {
	Sentence string
	Token    Token
	Object   string
}

// TODO: Finish writing options
func (s *SayStatement) String() string {
	var buf bytes.Buffer
	_, _ = buf.WriteString("SAY")

	return buf.String()
}

func (s *SayStatement) KeyTok() Token {
	return SAY
}

// parseSayStatement parses a look command and returns a Statement AST object.
// This function assumes the SAY token has already been consumed.
func (p *Parser) parseSayStatement() (*SayStatement, error) {
	stmt := &SayStatement{}

	tok, _, lit := p.ScanIgnoreWhitespace()
	if tok == TO {
		obj, err := p.ParseIdent()
		if err != nil {
			return nil, err
		}
		stmt.Object = obj
	}

	stmt.Token, _, stmt.Sentence = p.ScanSentence()

	// Add first token back
	if tok != TO {
		if len(stmt.Sentence) > 0 && stmt.Sentence[0] == '\'' {
			stmt.Sentence = lit + stmt.Sentence
		} else if len(stmt.Sentence) == 0 {
			stmt.Sentence = lit
		} else {
			stmt.Sentence = lit + " " + stmt.Sentence
		}
	}

	return stmt, nil
}

//******
// TELL
//******

// TellStatement allows you to converse with others in the same room
type TellStatement struct {
	Sentence string
	Token    Token
	Object   string // mob telling something to
}

// TODO: Finish writing options
func (s *TellStatement) String() string {
	var buf bytes.Buffer
	_, _ = buf.WriteString("TELL")

	return buf.String()
}

func (s *TellStatement) KeyTok() Token {
	return TELL
}

// parseTellStatement parses a look command and returns a Statement AST object.
// This function assumes the QUIT token has already been consumed.
func (p *Parser) parseTellStatement() (*TellStatement, error) {
	stmt := &TellStatement{}
	tok, _, lit := p.ScanIgnoreWhitespace()
	switch tok {
	case IDENT:
		stmt.Object = lit
	case EOF:
		return stmt, nil
	}

	stmt.Token, _, stmt.Sentence = p.ScanSentence()

	return stmt, nil
}

//******
// SHOUT
//******

// ShoutStatement allows you to converse with others in the same room
type ShoutStatement struct {
	Sentence string
	Token    Token
}

// TODO: Finish writing options
func (s *ShoutStatement) String() string {
	var buf bytes.Buffer
	_, _ = buf.WriteString("SHOUT")

	return buf.String()
}

func (s *ShoutStatement) KeyTok() Token {
	return SHOUT
}

// parseShoutStatement parses a look command and returns a Statement AST object.
// This function assumes the QUIT token has already been consumed.
func (p *Parser) parseShoutStatement() (*ShoutStatement, error) {
	stmt := &ShoutStatement{}
	stmt.Token, _, stmt.Sentence = p.ScanSentence()

	return stmt, nil
}

//******
// EMOTE
//******

// EmoteStatement allows you to converse with others in the same room
type EmoteStatement struct {
	Sentence string
	Token    Token
}

// TODO: Finish writing options
func (s *EmoteStatement) String() string {
	var buf bytes.Buffer
	_, _ = buf.WriteString("EMOTE")

	return buf.String()
}

func (s *EmoteStatement) KeyTok() Token {
	return EMOTE
}

// parseEmoteStatement parses a look command and returns a Statement AST object.
// This function assumes the QUIT token has already been consumed.
func (p *Parser) parseEmoteStatement() (*EmoteStatement, error) {
	stmt := &EmoteStatement{}
	stmt.Token, _, stmt.Sentence = p.ScanSentence()

	return stmt, nil
}
