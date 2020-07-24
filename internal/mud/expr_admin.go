package mud

// QuitExpression represents a command for exiting the game.
type QuitExpression struct {
	mob *Mob
}

// func (s *QuitStatement) String() string {
// 	var buf bytes.Buffer
// 	_, _ = buf.WriteString("QUIT")

// 	return buf.String()
// }

// // parseQuitStatement parses a look command and returns a Statement AST object.
// // This function assumes the QUIT token has already been consumed.
// func (p *Parser) parseQuitStatement() (*QuitStatement, error) {
// 	return &QuitStatement{}, nil
// }

func (s *QuitExpression) Execute() {
	WorldInstance.getRoom(s.mob.CurrentRoom).exitWorld(s.mob)
	s.mob.conn.Write("You slowly fade away.\n")
	s.mob.conn.conn.Close()
	ServerInstance.onClientConnectionClosed(s.mob.conn, nil)
}
