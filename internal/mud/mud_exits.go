package mud

import "github.com/brianbroderick/wergild/internal/mql"

func NewExit(uid string, dir string) Exit {
	tok := mql.Lookup(dir)

	return Exit{
		Dest:      []Room{{UID: uid}},
		Direction: tok.String(),
		Portal:    "OPEN",
		Type:      "Exit",
	}
}
