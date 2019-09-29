package main

// Parser represents a parser.
// type Parser struct {
// 	s   *Scanner
// 	buf struct {
// 		tok Token  // last read token
// 		lit string // last read literal
// 		n   int    // buffer size (max=1)
// 	}
// }

// // NewParser returns a new instance of Parser.
// func NewParser(r io.Reader) *Parser {
// 	return &Parser{s: NewScanner(r)}
// }

// // Parse parses a Loo
// func (p *Parser) Parse() (string, error) {
// 	// First token should be a "LOOK" keyword.
// 	if tok, lit := p.scanIgnoreWhitespace(); tok != LOOK {
// 		return "", fmt.Errorf("found %q, expected LOOK", lit)
// 	}

// 	tok, lit := p.scanIgnoreWhitespace()
// 	switch tok {
// 	case EOF:
// 		return "", nil
// 	case UP:
// 		return "You look up to the sky.", nil
// 	case DOWN:
// 		return "You look down towards the ground.", nil
// 	}

// 	if tok != AT {
// 		return "", fmt.Errorf("found %q, expected AT", lit)
// 	}

// 	tok, lit = p.scanIgnoreWhitespace()
// 	if tok != IDENT {
// 		return "", fmt.Errorf("found %q, expected field", lit)
// 	}
// 	obj := lit

// 	// Return the successfully parsed statement.
// 	return obj, nil
// }

// // scan returns the next token from the underlying scanner.
// // If a token has been unscanned then read that instead.
// func (p *Parser) scan() (tok Token, lit string) {
// 	// If we have a token on the buffer, then return it.
// 	if p.buf.n != 0 {
// 		p.buf.n = 0
// 		return p.buf.tok, p.buf.lit
// 	}

// 	// Otherwise read the next token from the scanner.
// 	tok, lit = p.s.Scan()

// 	// Save it to the buffer in case we unscan later.
// 	p.buf.tok, p.buf.lit = tok, lit

// 	return
// }

// // scanIgnoreWhitespace scans the next non-whitespace token.
// func (p *Parser) scanIgnoreWhitespace() (tok Token, lit string) {
// 	tok, lit = p.scan()
// 	if tok == WS {
// 		tok, lit = p.scan()
// 	}
// 	return
// }

// // unscan pushes the previously read token back onto the buffer.
// func (p *Parser) unscan() { p.buf.n = 1 }
