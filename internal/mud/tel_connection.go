package mud

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"time"
)

/**
 * Various user states, recorded in the connection struct
 */
const STATE_WELCOME = 0
const STATE_LOGIN_USERNAME = 1
const STATE_LOGIN_PASSWORD = 2
const STATE_CHARACTER_CREATION = 3
const STATE_PLAYING = 20
const STATE_STATUE = 21

const MAX_PASSWORD_FAILURES = 3

type Connection struct {
	conn             net.Conn
	timeConnected    time.Time
	state            int8
	username         string
	Player           *Player
	passwordFailures int
}

func (connection *Connection) Write(message string) {
	connection.conn.Write([]byte(message))
}

func (connection *Connection) sendMOTD() {
	connection.Write(ServerInstance.Motd)
}

func (connection *Connection) listen() {
	reader := bufio.NewReader(connection.conn)

	connection.sendMOTD()
	connection.Write("What is your name: ")
	connection.state = STATE_LOGIN_USERNAME

	for {
		message, err := reader.ReadString('\n')

		if err != nil {
			fmt.Printf("Connection error: %s\n", connection.username)

			connection.state = STATE_STATUE

			connection.conn.Close()
			ServerInstance.onClientConnectionClosed(connection, err)
			return
		}

		message = strings.TrimSpace(message)

		switch connection.state {

		// Player has just seen the MOTD, and is asked for username
		case STATE_LOGIN_USERNAME:
			connection.state = STATE_LOGIN_PASSWORD
			connection.username = message
			connection.Write("Password: ")

		// Player is being asked to authenticate
		case STATE_LOGIN_PASSWORD:
			exists, player := ServerInstance.login(connection.username, message)

			if exists {
				if player != nil {
					// auth succeeded, do a bit of housekeeping
					connection.Player = player
					player.setConnection(connection)
					ServerInstance.onPlayerAuthenticated(connection)
				} else {
					// auth fails, try again
					connection.Write("Sorry, that wasn't right. Try again: ")

					connection.passwordFailures++
					if connection.passwordFailures > MAX_PASSWORD_FAILURES {
						connection.Write("Pfft...  Goodbye.")
						connection.conn.Close()
					}

				}
			} else {
				connection.state = STATE_CHARACTER_CREATION
				connection.Write("STATE_CHARACTER_CREATION\n")
			}

		case STATE_PLAYING:
			ServerInstance.onMessageReceived(connection, message)

		default:
			ServerInstance.onMessageReceived(connection, message)
		}
	}
}
