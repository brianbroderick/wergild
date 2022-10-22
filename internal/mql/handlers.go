package mql

import (
	"fmt"
)

var Language = &ParseTree{}

type ParseTree struct {
	Handlers map[Token]func(*Parser) (Statement, error)
	Keys     []string
}

func init() {
	// Iterate INTEGER times over rest of the command
	Language.Handle(INTEGER, func(p *Parser) (Statement, error) {
		return p.parseLoopStatement()
	})

	// Interact with the Environment
	Language.Handle(LOOK, func(p *Parser) (Statement, error) {
		return p.parseLookStatement()
	})

	Language.Handle(LISTEN, func(p *Parser) (Statement, error) {
		return p.parseListenStatement()
	})

	Language.Handle(SMELL, func(p *Parser) (Statement, error) {
		return p.parseSmellStatement()
	})

	Language.Handle(TOUCH, func(p *Parser) (Statement, error) {
		return p.parseTouchStatement()
	})

	// Communicate
	Language.Handle(SAY, func(p *Parser) (Statement, error) {
		return p.parseSayStatement()
	})

	Language.Handle(TELL, func(p *Parser) (Statement, error) {
		return p.parseTellStatement()
	})

	Language.Handle(SHOUT, func(p *Parser) (Statement, error) {
		return p.parseShoutStatement()
	})

	Language.Handle(EMOTE, func(p *Parser) (Statement, error) {
		return p.parseEmoteStatement()
	})

	// Directions

	Language.Handle(NORTH, func(p *Parser) (Statement, error) {
		return p.parseDirectionStatement()
	})

	Language.Handle(SOUTH, func(p *Parser) (Statement, error) {
		return p.parseDirectionStatement()
	})

	Language.Handle(EAST, func(p *Parser) (Statement, error) {
		return p.parseDirectionStatement()
	})

	Language.Handle(WEST, func(p *Parser) (Statement, error) {
		return p.parseDirectionStatement()
	})

	Language.Handle(UP, func(p *Parser) (Statement, error) {
		return p.parseDirectionStatement()
	})

	Language.Handle(DOWN, func(p *Parser) (Statement, error) {
		return p.parseDirectionStatement()
	})

	// Admin

	Language.Handle(SCORE, func(p *Parser) (Statement, error) {
		return p.parseScoreStatement()
	})

	Language.Handle(QUIT, func(p *Parser) (Statement, error) {
		return p.parseQuitStatement()
	})

	Language.Handle(IMAGINE, func(p *Parser) (Statement, error) {
		return p.parseImagineStatement()
	})

	Language.Handle(HELP, func(p *Parser) (Statement, error) {
		return p.parseHelpStatement()
	})

	// Catch all - usually represents feelings
	Language.Handle(IDENT, func(p *Parser) (Statement, error) {
		return p.parseFeelingStatement()
	})
}

// Handle registers a handler to be invoked when seeing the given token.
func (t *ParseTree) Handle(tok Token, fn func(*Parser) (Statement, error)) {
	// Verify that there is no conflict for this token in this parse tree.
	if _, conflict := t.Handlers[tok]; conflict {
		panic(fmt.Sprintf("conflict for token %s", tok))
	}

	if t.Handlers == nil {
		t.Handlers = make(map[Token]func(*Parser) (Statement, error))
	}
	t.Handlers[tok] = fn
	t.Keys = append(t.Keys, tok.String())
}

// Parse parses a statement using the language defined in the parse tree.
func (t *ParseTree) Parse(p *Parser) (Statement, error) {
	for {
		tok, pos, lit := p.ScanIgnoreWhitespace()

		if stmt := t.Handlers[tok]; stmt != nil {
			return stmt(p)
		}

		// There were no registered handlers. Return the valid tokens in the order they were added.
		return nil, newParseError(tokstr(tok, lit), t.Keys, pos)
	}
}

// Statement represents a single command
type Statement interface {
	String() string
	KeyTok() Token
}
