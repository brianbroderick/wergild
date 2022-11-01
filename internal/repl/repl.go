package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/brianbroderick/wergild/internal/lexer"
	"github.com/brianbroderick/wergild/internal/token"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprintf(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)

		for tok, _ := l.Scan(); tok.Type != token.EOF; tok, _ = l.Scan() {
			fmt.Fprintf(out, "%+v\n", tok)
		}
	}
}
