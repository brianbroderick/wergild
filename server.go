package main

import (
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

type server struct {
	connectionList []*connection
	playerList     []*connection
	// roomList       map[int]*Room
	ticker *time.Ticker
	motd   string
	menu   string
}

var serverInstance *server

func getServer() *server {

	if serverInstance == nil {
		serverInstance = &Server{
			connectionList: make([]*Connection, 0),
			playerList:     make([]*Connection, 0),
		}

		pwd, _ := os.Getwd()

		fmt.Printf("[CONFIG] Current working directory set to %s\n", pwd)

		// 1. Pull in the welcome screen
		fmt.Println("[CONFIG] Pulling MOTD")
		// motdBytes, _ := ioutil.ReadFile(pwd + "/resources/MOTD")
		// serverInstance.motd = string(motdBytes)
		serverInstance.motd = "Wergild"

		// 2. Pull in the menu
		fmt.Println("[CONFIG] Pulling Menu")
		// menuBytes, _ := ioutil.ReadFile(pwd + "/resources/Menu")
		// serverInstance.menu = string(menuBytes)
		serverInstance.menu = `[1] Enter the realms
		[2] Change your password
		[3] Delete your character
		[4] Exit`

		// // 3. Prepare the command hashes
		// fmt.Println("[CONFIG] Preparing commands")
		// prepareCommands()

		// // 4. Load in the rooms
		// fmt.Println("[CONFIG] Loading rooms")
		// ServerInstance.roomList = loadRooms()
	}

	return serverInstance
}

func (server *server) addConnection(conn net.Conn) *connection {
	newConnection := connection{conn: conn, timeConnected: time.Now(), state: stateWelcome}
	server.connectionList = append(server.connectionList, &newConnection)
	go newConnection.listen()
	fmt.Printf("[CONN] There are %d connected users.\n", server.connectionCount())
	return &newConnection
}

func (server *server) onClientConnectionClosed(conn *connection, err error) {
	// delete the connection

	for i, c := range server.connectionList {
		if c == conn {
			server.connectionList[i] = server.connectionList[len(server.connectionList)-1]
			server.connectionList[len(server.connectionList)-1] = nil
			server.connectionList = server.connectionList[:len(server.connectionList)-1]
			break
		}
	}

	fmt.Printf("[DISC] There are %d connected users.\n", server.connectionCount())
}

/**
 * User log-in process
 */
func (server *server) login(username string, password string) (bool, *player) {

	if !userExists(username) {
		return false, nil
	}

	player := authenticate(username, password)

	if player != nil {
		return true, player
	}

	return true, nil
}

/**
 * Check the database to see if the player exists
 * @TODO perform an actual database scan
 */
func userExists(username string) bool {
	return username == "brian"
}

/**
 * Authenticate a user via database, and fetch the player
 * @TODO implement the actual check!
 */
func authenticate(username string, password string) *player {
	if username == "brian" && password == "123" {
		// player := &player{Name: username, CurrentRoom: 1, hitPointsMax: 100, hitPoints: 50, vitalityMax: 250, vitality: 250, race: getRace("demon")}
		player := &player{name: username}
		// player.inventory = []*Item{
		// 	{Name: "A Dark Sword", Description: "A test object to test object loading"},
		// }

		return player

	}
	return nil
}

/**
 * Server-side trigger when authentication occurs in the comm handler
 */
func (server *server) onPlayerAuthenticated(conn *connection) {
	fmt.Printf("[AUTH] Player authenticated (%s)\n", conn.player.name)
	server.playerList = append(server.playerList, conn)

	conn.state = stateLoginMenu
	conn.write("Welcome. Death Awaits.\n")
	conn.sendMenu()
}

func (server *server) start() {
	server.ticker = time.NewTicker(time.Millisecond * 3000)

	go func() {
		for range server.ticker.C {
			for _, c := range server.playerList {
				// fmt.Printf("[TICK] Running update tick on player (%s) at state [%d]\n", c.username, c.state)
				fmt.Printf("[TICK] Running update tick on player at state [%d]\n", c.state)
				if c.player != nil {
					c.player.pulseUpdate()
				}
			}
		}
	}()

}

/**
 * How many connections are active?
 */
func (server *server) connectionCount() int {
	return len(server.connectionList)
}

/**
 * A message has been received on a given connection descriptor
 */
func (server *server) onMessageReceived(conn *connection, message string) {

	if len(message) == 0 {
		conn.player.sendPrompt()
		return
	}
	words := strings.Fields(message)
	input, arguments := words[0], words[1:]

	conn.player.do(input, arguments)
	conn.player.sendPrompt()
}

// func (server *server) getRoom(roomId int) *room {
// 	return server.roomList[roomId]
// }
