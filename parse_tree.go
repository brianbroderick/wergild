package main

import "fmt"

var Language = &ParseTree{}

type ParseTree struct {
	Handlers map[Token]func(*Parser) (Statement, error)
	Tokens   map[Token]*ParseTree
	Keys     []string
}

// With passes the current parse tree to a function to allow nested functions.
func (t *ParseTree) With(fn func(*ParseTree)) {
	fn(t)
}

// Group groups together a set of related handlers with a common token prefix.
func (t *ParseTree) Group(tokens ...Token) *ParseTree {
	for _, tok := range tokens {
		// Look for the parse tree for this token.
		if subtree := t.Tokens[tok]; subtree != nil {
			t = subtree
			continue
		}

		// No subtree exists yet. Verify that we don't have a conflicting
		// statement.
		if _, conflict := t.Handlers[tok]; conflict {
			panic(fmt.Sprintf("conflict for token %s", tok))
		}

		// Create the new parse tree and register it inside of this one for
		// later reference.
		newT := &ParseTree{}
		if t.Tokens == nil {
			t.Tokens = make(map[Token]*ParseTree)
		}
		t.Tokens[tok] = newT
		t.Keys = append(t.Keys, tok.String())
		t = newT
	}
	return t
}

// Handle registers a handler to be invoked when seeing the given token.
func (t *ParseTree) Handle(tok Token, fn func(*Parser) (Statement, error)) {
	// Verify that there is no conflict for this token in this parse tree.
	if _, conflict := t.Tokens[tok]; conflict {
		panic(fmt.Sprintf("conflict for token %s", tok))
	}

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
		if subtree := t.Tokens[tok]; subtree != nil {
			t = subtree
			continue
		}

		if stmt := t.Handlers[tok]; stmt != nil {
			return stmt(p)
		}

		// There were no registered handlers. Return the valid tokens in the order they were added.
		return nil, newParseError(tokstr(tok, lit), t.Keys, pos)
	}
}

func (t *ParseTree) Clone() *ParseTree {
	newT := &ParseTree{}
	if t.Handlers != nil {
		newT.Handlers = make(map[Token]func(*Parser) (Statement, error), len(t.Handlers))
		for tok, handler := range t.Handlers {
			newT.Handlers[tok] = handler
		}
	}

	if t.Tokens != nil {
		newT.Tokens = make(map[Token]*ParseTree, len(t.Tokens))
		for tok, subtree := range t.Tokens {
			newT.Tokens[tok] = subtree.Clone()
		}
	}
	return newT
}

// This is where the handlers are set up
func init() {

}
