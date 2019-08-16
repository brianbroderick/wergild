package main

import (
	"net"
	"time"
)

const stateWelcome = 0
const stateLoginUsername = 1
const stateLoginPassword = 2
const stateLoginMenu = 3
const stateCharacterCreation = 4
const statePlaying = 20

const menuStartGame = "1"
const menuChangePassword = "2"
const menuDeleteCharacter = "3"
const menuExit = "4"

type connection struct {
	conn          net.Conn
	timeConnected time.Time
	state         int8
	player        *player
}

func (connection *connection) write(message string) {
	connection.conn.Write([]byte(message))
}
