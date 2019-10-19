package main

import "bytes"

// QuitStatement represents a command for exiting the game.
type QuitStatement struct {
	player *Player
}

func (s *QuitStatement) String() string {
	var buf bytes.Buffer
	_, _ = buf.WriteString("QUIT")

	return buf.String()
}

// parseQuitStatement parses a look command and returns a Statement AST object.
// This function assumes the QUIT token has already been consumed.
func (p *Parser) parseQuitStatement() (*QuitStatement, error) {
	return &QuitStatement{}, nil
}

func (s *QuitStatement) execute() {
	s.player.connection.Write("You slowly fade away.\n")
	s.player.connection.conn.Close()
	ServerInstance.playerExited(s.player.connection, nil)
	ServerInstance.onClientConnectionClosed(s.player.connection, nil)
}

func (s *QuitStatement) setPlayer(player *Player) {
	s.player = player
}
