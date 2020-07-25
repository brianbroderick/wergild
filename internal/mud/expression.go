package mud

import (
	"strings"

	"github.com/brianbroderick/logit"
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
	// admin
	Executor.Tokens[mql.QUIT] = func(x *Expr) {
		l := &QuitExpression{mob: x.mob}
		l.Execute()
	}

	// Interact with the Environment
	Executor.Tokens[mql.LOOK] = func(x *Expr) {
		l := &LookExpression{mob: x.mob, stmt: x.stmt.(*mql.LookStatement)}
		l.Execute()
	}

	Executor.Tokens[mql.LISTEN] = func(x *Expr) {
		l := &ListenExpression{mob: x.mob, stmt: x.stmt.(*mql.ListenStatement)}
		l.Execute()
	}

	Executor.Tokens[mql.SMELL] = func(x *Expr) {
		l := &SmellExpression{mob: x.mob, stmt: x.stmt.(*mql.SmellStatement)}
		l.Execute()
	}

	Executor.Tokens[mql.TOUCH] = func(x *Expr) {
		l := &TouchExpression{mob: x.mob, stmt: x.stmt.(*mql.TouchStatement)}
		l.Execute()
	}

	// Communicate
	Executor.Tokens[mql.SAY] = func(x *Expr) {
		l := &SayExpression{mob: x.mob, stmt: x.stmt.(*mql.SayStatement)}
		l.Execute()
	}

	Executor.Tokens[mql.TELL] = func(x *Expr) {
		l := &TellExpression{mob: x.mob, stmt: x.stmt.(*mql.TellStatement)}
		l.Execute()
	}

	Executor.Tokens[mql.SHOUT] = func(x *Expr) {
		l := &ShoutExpression{mob: x.mob, stmt: x.stmt.(*mql.ShoutStatement)}
		l.Execute()
	}

	Executor.Tokens[mql.EMOTE] = func(x *Expr) {
		l := &EmoteExpression{mob: x.mob, stmt: x.stmt.(*mql.EmoteStatement)}
		l.Execute()
	}

	// Directions
	Executor.Tokens[mql.DIRECTION] = func(x *Expr) {
		l := &DirectionExpression{mob: x.mob, stmt: x.stmt.(*mql.DirectionStatement)}
		l.Execute()
	}
}

func (mob *Mob) Do(message string) {
	stmt, err := mql.NewParser(strings.NewReader(message)).ParseStatement()
	if err != nil {
		return
	}

	s := Executor.Tokens[stmt.KeyTok()]
	if s != nil {
		s(&Expr{msg: message, stmt: stmt, mob: mob})
	} else {
		logit.Info("command not found, %v", s)
	}
}
