package main

// FeelingStatement represents a command for looking at a room or object.
type FeelingStatement struct {
	player *Player
	ident  string // feeling. if not found, respond with something like "what?"
}

func (s *FeelingStatement) execute() {
	switch s.ident {
	case "laugh":
		s.player.connection.Write("You fall down laughing.\n")
	default:
		s.player.connection.Write("what? \n")
	}

}
