package mud

import (
	"strings"

	"github.com/brianbroderick/wergild/internal/mql"
)

var Executor = &ExpressionTree{
	Tokens: make(map[mql.Token]func(*Expr)),
}

type ExpressionTree struct {
	Tokens map[mql.Token]func(*Expr)
}

type Expr struct {
	mob  *Mob
	msg  string
	stmt mql.Statement
}

func init() {
	Executor.Tokens[mql.LOOK] = func(x *Expr) {
		l := &LookExpression{mob: x.mob, stmt: x.stmt.(*mql.LookStatement)}
		l.Execute()
	}
}

func (mob *Mob) Do(message string) {
	stmt, err := mql.NewParser(strings.NewReader(message)).ParseStatement()
	if err != nil {
		return
	}

	s := Executor.Tokens[stmt.KeyTok()]
	s(&Expr{msg: message, stmt: stmt, mob: mob})
}

// func (mob *Mob) do(message string) {
// 	stmt, err := mql.NewParser(strings.NewReader(message)).ParseStatement()
// 	if err != nil {
// 		mob.conn.Write(fmt.Sprint(err) + "\n")
// 		return
// 	}

// 	stmt.Execute()
// }
