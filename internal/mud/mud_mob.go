package mud

import (
	"fmt"
	"strings"
)

// probably not needed?
func (mob *Mob) getCurrentRoom() string {
	return mob.CurrentRoom
}

func (mob *Mob) sendPrompt() {
	if mob == nil || mob.conn == nil {
		return
	}

	str := fmt.Sprintf("%d:%d> ",
		mob.Hp,
		mob.Ap)

	mob.conn.Write(str)
}

func (mob *Mob) do(message string) {
	stmt, err := NewParser(strings.NewReader(message)).ParseStatement()
	if err != nil {
		mob.conn.Write(fmt.Sprint(err) + "\n")
		return
	}

	stmt.setMob(mob)
	stmt.execute()
}
