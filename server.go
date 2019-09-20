package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"strings"
	"time"
)

type Server struct {
	connectionList []*Connection
	playerList     []*Connection
	ticker         *time.Ticker
	Motd           string
	Menu           string
	roomList       map[int]*Room
}

func newDescriptor(connection net.Conn) {
	ServerInstance.AddConnection(connection)
}

func (server *Server) AddConnection(connection net.Conn) *Connection {
	newConnection := Connection{conn: connection, timeConnected: time.Now()}
	server.connectionList = append(server.connectionList, &newConnection)
	go newConnection.listen()
	fmt.Printf("[CONN] There are %d connected users.\n", server.ConnectionCount())
	return &newConnection
}

func GetServer() *Server {
	if ServerInstance == nil {
		ServerInstance = &Server{
			connectionList: make([]*Connection, 0),
			playerList:     make([]*Connection, 0),
		}

		pwd, _ := os.Getwd()

		// fmt.Printf("[CONFIG] Current working directory set to %s\n", pwd)

		// 1. Pull in the welcome screen
		fmt.Println("[CONFIG] Pulling MOTD")
		motdBytes, _ := ioutil.ReadFile(pwd + "/resources/MOTD")
		ServerInstance.Motd = string(motdBytes)

		// 2. Prepare the command hashes
		fmt.Println("[CONFIG] Preparing commands")
		prepareCommands()

		// 3. Load in the rooms
		fmt.Println("[CONFIG] Loading rooms")
		ServerInstance.roomList = loadRooms()
	}

	return ServerInstance
}

func (server *Server) onClientConnectionClosed(connection *Connection, err error) {
	// delete the connection

	for i, conn := range server.connectionList {
		if conn == connection {
			server.connectionList[i] = server.connectionList[len(server.connectionList)-1]
			server.connectionList[len(server.connectionList)-1] = nil
			server.connectionList = server.connectionList[:len(server.connectionList)-1]
			break
		}
	}

	fmt.Printf("[DISC] There are %d connected users.\n", server.ConnectionCount())
}

func (server *Server) onMessageReceived(connection *Connection, message string) {

	if len(message) == 0 {
		connection.Player.sendPrompt()
		return
	}
	words := strings.Fields(message)
	input, arguments := words[0], words[1:]

	fmt.Printf("%v %v\n", input, arguments)

	// connection.Player.do(input, arguments)
	connection.Player.sendPrompt()
}

func (server *Server) Start() {
	server.ticker = time.NewTicker(time.Millisecond * 3000)

	go func() {
		for range server.ticker.C {
			for _, c := range server.playerList {
				fmt.Printf("[TICK] Running update tick on player (%s) at state [%d]\n", c.username, c.state)
				if c.Player != nil {
					c.Player.pulseUpdate()
				}
			}
		}
	}()

}

func (server *Server) ConnectionCount() int {
	return len(server.connectionList)
}

func (server *Server) login(username string, password string) (bool, *Player) {

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
	return username == "azkul"
}

/**
 * Authenticate a user via database, and fetch the player
 * @TODO implement the actual check!
 */
func authenticate(username string, password string) *Player {
	if username == "azkul" && password == "123" {
		player := &Player{Name: username, CurrentRoom: 1, hitPointsMax: 100, hitPoints: 50}
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
func (server *Server) onPlayerAuthenticated(connection *Connection) {
	fmt.Printf("[AUTH] Player authenticated (%s)\n", connection.Player.Name)
	server.playerList = append(server.playerList, connection)

	connection.state = STATE_PLAYING
	connection.Write("Welcome. Death Awaits.\n")
}

func (server *Server) getRoom(roomId int) *Room {
	return server.roomList[roomId]
}
