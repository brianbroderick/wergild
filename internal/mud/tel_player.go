package mud

import (
	"fmt"
	"strings"
)

type Player struct {
	Name        string
	CurrentRoom string
	connection  *Connection
	mob         *Mob
	hp          int // hit points - hitting zero means death
	hpMax       int // max hit points
	ap          int // action points. - hitting zero means you are too fatigued to do the action
	apMax       int // max action points
}

func (player *Player) setConnection(connection *Connection) {
	player.connection = connection
}

func (player *Player) getCurrentRoom() string {
	return player.CurrentRoom
}

func (player *Player) sendPrompt() {
	str := fmt.Sprintf("%d:%d> ",
		player.hp,
		player.ap)
	player.connection.Write(str)
}

func (player *Player) pulseUpdate() {
	currentRoom := WorldInstance.getRoom(player.CurrentRoom)
	currentRoom.showEnv(player)

	// fmt.Printf("This is for player %s\n", player.Name)
	//if player.hitPoints < player.hitPointsMax {
	//player.hitPoints = min(player.hitPoints+player.regenHP(), player.hitPointsMax)
	// player.sendPrompt()
	//}
}

func (player *Player) do(message string) {
	stmt, err := NewParser(strings.NewReader(message)).ParseStatement()
	if err != nil {
		player.connection.Write(fmt.Sprint(err) + "\n")
		return
	}

	stmt.setPlayer(player)
	stmt.execute()
}
