package main

import (
	"fmt"

	"github.com/brianbroderick/wergild/internal/evaluator"
	"github.com/brianbroderick/wergild/internal/lexer"
	"github.com/brianbroderick/wergild/internal/object"
	"github.com/brianbroderick/wergild/internal/parser"
)

// This is a benchmark to see how fast the interpreter is.
// Running this command gives the following output on a
// MacBook Pro 2021, M1 Max, 64 GB Ram (Monterey 12.6)
// time ./fib
// 9227465
// ./fib  11.33s user 0.33s system 105% cpu 11.081 total
// Here are some benchmarks for some JS interpreters running in Go:
// http://cyan.ly/blog/embedding-javascript-in-go-survey

func main() {
	input := `let fib = fn(x) {
	  if (x == 0) {
			0
		}	else {
			if (x == 1) {
				1
			} else {
				fib(x - 1) + fib(x - 2);
			}
		}
	};
	
	fib(35);`

	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()
	env := object.NewEnvironment()

	evaluated := evaluator.Eval(program, env)
	if evaluated != nil {
		fmt.Println(evaluated.Inspect())
	}
}
