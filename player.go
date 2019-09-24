package main

import "fmt"

type Player struct {
	Name        string
	CurrentRoom int
	connection  *Connection
	//inventory    []*Item
	//race         *Race
	hitPoints    int
	hitPointsMax int
}

func (player *Player) setConnection(connection *Connection) {
	player.connection = connection
}

func (player *Player) getCurrentRoom() int {
	return player.CurrentRoom
}

func (player *Player) sendPrompt() {
	// str := fmt.Sprintf("\n%s%s< %dh/%dH %s%s >\n<>%s\n",
	// 	FG_GREEN,
	// 	MOD_FAINT,
	// 	player.hitPoints,
	// 	player.hitPointsMax,
	// 	FG_GREEN,
	// 	MOD_FAINT,
	// 	MOD_CLEAR,
	// )

	str := fmt.Sprintf("%d:%d> ",
		player.hitPoints,
		player.hitPointsMax)
	player.connection.Write(str)
}

func (player *Player) pulseUpdate() {
	// fmt.Printf("This is for player %s\n", player.Name)
	//if player.hitPoints < player.hitPointsMax {
	//player.hitPoints = min(player.hitPoints+player.regenHP(), player.hitPointsMax)
	// player.sendPrompt()
	//}
}

func (player *Player) do(verb string, arguments []string) {
	command, err := getCommand(verb)
	if err != nil {
		player.connection.Write(fmt.Sprint(err) + "\n")
		return
	}

	command.closure(player, arguments)
}
