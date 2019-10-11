package main

// LoopStatement represents a command for repeating another command X times
type LoopStatement struct {
	player *Player
	i      int
	msg    string // msg to parse and run i times
}

func (s *LoopStatement) execute() {
	s.player.connection.Write("Loops are not implemented yet.\n")
}
