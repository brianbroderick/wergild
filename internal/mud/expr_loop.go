package mud

import (
	"github.com/brianbroderick/wergild/internal/mql"
)

// LoopExpression represents a command for repeating another command X times
type LoopExpression struct {
	mob  *Mob
	stmt *mql.LoopStatement
}

// TODO: modify to use Expression
func (s *LoopExpression) Execute() {
	for i := 0; i < s.stmt.I; i++ {
		cmd := Executor.Tokens[s.stmt.NStmt.KeyTok()]
		cmd(&Expr{msg: "loop", stmt: s.stmt.NStmt, mob: s.mob})
	}
}
