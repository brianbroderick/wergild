package mql

import (
	"bytes"
	"strings"
)

//******
// QUIT
//******

// QuitStatement represents a command for exiting the game.
type QuitStatement struct{}

func (s *QuitStatement) String() string {
	var buf bytes.Buffer
	_, _ = buf.WriteString("QUIT")

	return buf.String()
}

func (s *QuitStatement) KeyTok() Token {
	return QUIT
}

// parseQuitStatement parses a look command and returns a Statement AST object.
// This function assumes the QUIT token has already been consumed.
func (p *Parser) parseQuitStatement() (*QuitStatement, error) {
	return &QuitStatement{}, nil
}

//********
// IMAGINE
//********

// ImagineStatement is the admin command to create an object such as a room
type ImagineStatement struct {
	Object    Token
	Direction Token
	Name      string
}

func (s *ImagineStatement) String() string {
	var buf bytes.Buffer
	_, _ = buf.WriteString("IMAGINE")

	return buf.String()
}

func (s *ImagineStatement) KeyTok() Token {
	return IMAGINE
}

// parseImagineStatement parses an imagine command and returns a Statement AST object.
// This function assumes the IMAGINE token has already been consumed.
func (p *Parser) parseImagineStatement() (*ImagineStatement, error) {
	stmt := &ImagineStatement{}

	stmt.Object, _, _ = p.ScanIgnoreWhitespace()
	if stmt.Object == ROOM {
		dir, pos, lit := p.ScanIgnoreWhitespace()
		if val, ok := directionKeywords[strings.ToLower(Tokens[dir])]; ok {
			stmt.Direction = val
		} else {
			return nil, newParseError(tokstr(dir, lit), []string{"direction"}, pos)
		}

		name, err := p.ParseIdent()
		if err != nil {
			return nil, err
		}
		stmt.Name = name
	}

	return stmt, nil
}
