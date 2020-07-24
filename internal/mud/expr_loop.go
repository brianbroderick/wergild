package mud

import "github.com/brianbroderick/wergild/internal/mql"

// LoopExpression represents a command for repeating another command X times
type LoopExpression struct {
	mob   *Mob
	i     int
	nStmt mql.Statement // msg to parse and run i times
}

// TODO: modify to use Expression
func (s *LoopExpression) Execute() {

	// for i := 0; i < s.i; i++ {
	// 	s.nStmt.execute()
	// }
}
