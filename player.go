package main

import (
	"bufio"
	"fmt"
	"strings"
)

type player struct {
	connection  *connection
	currentRoom int
	name        string
}

const (
	posProne    = 2
	posSitting  = 1
	posStanding = 0
)

func (player *player) setConnection(conn *connection) {
	player.connection = conn
}

func (player *player) getCurrentRoom() int {
	return player.currentRoom
}

func (player *player) sendPrompt() {
	str := "Prompt:"
	player.connection.write(str)
}

func (connection *connection) sendMOTD() {
	connection.write(serverInstance.motd)
	connection.write("What is your name, mortal? ")
}

func (connection *connection) sendMenu() {
	connection.write(serverInstance.menu)
}

func (connection *connection) listen() {
	reader := bufio.NewReader(connection.conn)

	for {
		message, err := reader.ReadString('\n')

		if err != nil {
			connection.conn.Close()
			serverInstance.onClientConnectionClosed(connection, err)
			return
		}

		message = strings.TrimSpace(message)

		serverInstance.onMessageReceived(connection, message)
	}

}

func (player *player) do(verb string, arguments []string) {
	command, err := getCommand(verb)
	if err != nil {
		player.connection.write(fmt.Sprint(err))
		return
	}

	command.closure(player, arguments)
}
