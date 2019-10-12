package main

import (
	"fmt"
	"strings"
)

type Player struct {
	Name        string
	CurrentRoom int
	connection  *Connection
	//inventory    []*Item
	//race         *Race
	hitPoints       int
	hitPointsMax    int
	actionPoints    int
	actionPointsMax int
}

func (player *Player) setConnection(connection *Connection) {
	player.connection = connection
}

func (player *Player) getCurrentRoom() int {
	return player.CurrentRoom
}

func (player *Player) sendPrompt() {
	str := fmt.Sprintf("%d:%d> ",
		player.hitPoints,
		player.actionPoints)
	player.connection.Write(str)
}

func (player *Player) pulseUpdate() {
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
