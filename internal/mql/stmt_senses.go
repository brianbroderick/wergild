package mql

import (
	"bytes"
)

//*******
// LOOK
//*******

// LookStatement represents a command for looking at a room or object.
type LookStatement struct {
	Ident string // object that is being looked at, ex: `chair` in `look at chair`
	Token Token  // usually a direction. i.e. `north` in `look north`
}

// parseLookStatement parses a look command and returns a Statement AST object.
// This function assumes the LOOK token has already been consumed.
func (p *Parser) parseLookStatement() (*LookStatement, error) {
	stmt := &LookStatement{}
	var err error

	tok, _, _ := p.ScanIgnoreWhitespace()

	directions := [6]Token{NORTH, SOUTH, EAST, WEST, UP, DOWN}
	for _, direction := range directions {
		if tok == direction {
			stmt.Token = tok
			return stmt, nil
		}
	}

	switch tok {
	case AT, IN:
		stmt.Token = tok
	case EOF:
		stmt.Token = EOF
		return stmt, nil
	default:
		stmt.Token = NIL
		return stmt, nil
	}

	stmt.Ident, err = p.ParseIdent()
	if err != nil {
		return nil, err
	}

	return stmt, nil
}

func (s *LookStatement) KeyTok() Token {
	return LOOK
}

func (s *LookStatement) String() string {
	var buf bytes.Buffer
	_, _ = buf.WriteString("LOOK")

	directions := [6]Token{NORTH, SOUTH, EAST, WEST, UP, DOWN}
	for _, direction := range directions {
		if s.Token == direction {
			_, _ = buf.WriteString(" " + tokens[s.Token])
		}
	}

	switch s.Token {
	case AT, ON, IN:
		_, _ = buf.WriteString(" " + tokens[s.Token] + " ")
		_, _ = buf.WriteString(s.Ident)
	case NIL:
		_, _ = buf.WriteString(s.Ident)
	}

	return buf.String()
}

//*******
// LISTEN
//*******

// ListenStatement represents a command for listening to a room or object.
type ListenStatement struct {
	Ident string // object that is being listened to, ex: `clock` in `listen to clock`
	Token Token  // usually a direction. i.e. `north` in `listen north`
}

// parseListenStatement parses a listen command and returns a Statement AST object.
// This function assumes the LISTEN token has already been consumed.
func (p *Parser) parseListenStatement() (*ListenStatement, error) {
	stmt := &ListenStatement{}
	var err error

	tok, _, _ := p.ScanIgnoreWhitespace()

	directions := [6]Token{NORTH, SOUTH, EAST, WEST, UP, DOWN}
	for _, direction := range directions {
		if tok == direction {
			stmt.Token = tok
			return stmt, nil
		}
	}

	switch tok {
	case TO, IN:
		stmt.Token = tok
	case EOF:
		stmt.Token = EOF
		return stmt, nil
	default:
		stmt.Token = NIL
		return stmt, nil
	}

	stmt.Ident, err = p.ParseIdent()
	if err != nil {
		return nil, err
	}

	return stmt, nil
}

func (s *ListenStatement) KeyTok() Token {
	return LISTEN
}

func (s *ListenStatement) String() string {
	var buf bytes.Buffer
	_, _ = buf.WriteString("LISTEN")

	directions := [6]Token{NORTH, SOUTH, EAST, WEST, UP, DOWN}
	for _, direction := range directions {
		if s.Token == direction {
			_, _ = buf.WriteString(" " + tokens[s.Token])
		}
	}

	switch s.Token {
	case TO, IN:
		_, _ = buf.WriteString(" " + tokens[s.Token] + " ")
		_, _ = buf.WriteString(s.Ident)
	case NIL:
		_, _ = buf.WriteString(s.Ident)
	}

	return buf.String()
}

//*******
// SMELL
//*******

// SmellStatement represents a command for smelling a room or object.
type SmellStatement struct {
	Ident string // object that is being smelled, ex: `chair` in `smell chair`
	Token Token  // usually not used, but you might try to `smell north`
}

// parseSmellStatement parses a look command and returns a Statement AST object.
// This function assumes the Smell token has already been consumed.
func (p *Parser) parseSmellStatement() (*SmellStatement, error) {
	stmt := &SmellStatement{}
	var err error

	tok, _, lit := p.ScanIgnoreWhitespace()

	directions := [6]Token{NORTH, SOUTH, EAST, WEST, UP, DOWN}
	for _, direction := range directions {
		if tok == direction {
			stmt.Token = tok
			return stmt, nil
		}
	}

	switch tok {
	case IN:
		stmt.Token = tok
	case EOF:
		stmt.Token = EOF
		return stmt, nil
	default:
		stmt.Token = IDENT
		stmt.Ident = lit
		return stmt, nil
	}

	stmt.Ident, err = p.ParseIdent()
	if err != nil {
		return nil, err
	}

	return stmt, nil
}

func (s *SmellStatement) KeyTok() Token {
	return SMELL
}

func (s *SmellStatement) String() string {
	var buf bytes.Buffer
	_, _ = buf.WriteString("SMELL")

	directions := [6]Token{NORTH, SOUTH, EAST, WEST, UP, DOWN}
	for _, direction := range directions {
		if s.Token == direction {
			_, _ = buf.WriteString(" " + tokens[s.Token])
		}
	}

	switch s.Token {
	case IN:
		_, _ = buf.WriteString(" " + tokens[s.Token] + " ")
		_, _ = buf.WriteString(s.Ident)
	case IDENT:
		_, _ = buf.WriteString(" " + s.Ident)
	}

	return buf.String()
}

//*******
// TOUCH
//*******

// TouchStatement represents a command for touching an object.
type TouchStatement struct {
	Ident string // object that is being touched, ex: `chair` in `touch chair`
	Token Token  // usually not used, but you might try to `touch north`
}

func (s *TouchStatement) KeyTok() Token {
	return TOUCH
}

func (s *TouchStatement) String() string {
	var buf bytes.Buffer
	_, _ = buf.WriteString("TOUCH")

	directions := [6]Token{NORTH, SOUTH, EAST, WEST, UP, DOWN}
	for _, direction := range directions {
		if s.Token == direction {
			_, _ = buf.WriteString(" " + tokens[s.Token])
		}
	}

	switch s.Token {
	case IN:
		_, _ = buf.WriteString(" " + tokens[s.Token] + " ")
		_, _ = buf.WriteString(s.Ident)
	case IDENT:
		_, _ = buf.WriteString(" " + s.Ident)
	}

	return buf.String()
}

// parseTouchStatement parses a look command and returns a Statement AST object.
// This function assumes the Touch token has already been consumed.
func (p *Parser) parseTouchStatement() (*TouchStatement, error) {
	stmt := &TouchStatement{}
	var err error

	tok, _, lit := p.ScanIgnoreWhitespace()

	directions := [6]Token{NORTH, SOUTH, EAST, WEST, UP, DOWN}
	for _, direction := range directions {
		if tok == direction {
			stmt.Token = tok
			return stmt, nil
		}
	}

	switch tok {
	case IN:
		stmt.Token = tok
	case EOF:
		stmt.Token = EOF
		return stmt, nil
	default:
		stmt.Token = IDENT
		stmt.Ident = lit
		return stmt, nil
	}

	stmt.Ident, err = p.ParseIdent()
	if err != nil {
		return nil, err
	}

	return stmt, nil
}
