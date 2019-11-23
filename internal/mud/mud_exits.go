package mud

func NewExit(uid string, dir string) Exit {
	tok := Lookup(dir)

	return Exit{
		Dest:      []Room{Room{UID: uid}},
		Direction: tok.String(),
		Portal:    "OPEN",
		Type:      "Exit",
	}
}
