package mql

import (
	"bytes"
	"strings"

	"github.com/brianbroderick/logit"
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
	Location  string
	Smell     string
	Listen    string
	Mob       string
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
	switch stmt.Object {
	case ROOM:
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
	case LOCATION:
		_, _, lit := p.ScanSentence()
		stmt.Location = lit
	case SMELL:
		_, _, lit := p.ScanSentence()
		stmt.Smell = lit
	case LISTEN:
		_, _, lit := p.ScanSentence()
		stmt.Listen = lit
	case MOB:
		name, err := p.ParseIdent()
		if err != nil {
			return nil, err
		}
		stmt.Mob = name
		logit.Info("STMT %s", stmt.Mob)
	}

	return stmt, nil
}

//********
// HELP
//********

// HelpStatement is the admin command to display help
type HelpStatement struct {
	Token Token
	Ident string
}

func (s *HelpStatement) String() string {
	var buf bytes.Buffer
	_, _ = buf.WriteString("HELP")

	return buf.String()
}

func (s *HelpStatement) KeyTok() Token {
	return HELP
}

// parseHelpStatement parses a help command and returns a Statement AST object.
// This function assumes the HELP token has already been consumed.
func (p *Parser) parseHelpStatement() (*HelpStatement, error) {
	stmt := &HelpStatement{}
	var err error

	tok, _, _ := p.ScanIgnoreWhitespace()

	switch tok {
	case EOF:
		stmt.Token = EOF
		return stmt, nil
	default:
		stmt.Token = HELP
		p.Unscan()
	}

	stmt.Ident, err = p.ParseIdent()
	if err != nil {
		return nil, err
	}

	return stmt, nil
}
