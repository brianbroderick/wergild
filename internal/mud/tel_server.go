package mud

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"time"
)

type Server struct {
	connectionList []*Connection
	playerList     []*Connection
	ticker         *time.Ticker
	Motd           string
	// roomList       map[string]*Room
}

func NewDescriptor(connection net.Conn) {
	ServerInstance.AddConnection(connection)
}

func GetServer() *Server {
	if ServerInstance == nil {
		ServerInstance = &Server{
			connectionList: make([]*Connection, 0),
			playerList:     make([]*Connection, 0),
		}

		pwd, _ := os.Getwd()

		// 1. Pull in the welcome screen
		fmt.Println("[CONFIG] Pulling MOTD")
		motdBytes, _ := ioutil.ReadFile(pwd + "/resources/MOTD")
		ServerInstance.Motd = string(motdBytes)

		// 2. Prepare the command hashes
		// fmt.Println("[CONFIG] Preparing commands")
		// prepareCommands()
	}

	return ServerInstance
}

func (server *Server) onMessageReceived(connection *Connection, message string) {

	if len(message) == 0 {
		connection.Player.sendPrompt()
		return
	}

	fmt.Printf("%s\n", message)

	// // message = applyNick(message)
	// words := strings.Fields(message)
	// input, arguments := words[0], words[1:]

	// fmt.Printf("%v %v\n", input, arguments)

	connection.Player.do(message)
	connection.Player.sendPrompt()

}

func (server *Server) Start() {
	server.ticker = time.NewTicker(time.Millisecond * 3000)

	go func() {
		for range server.ticker.C {
			for _, c := range server.playerList {
				// fmt.Printf("[TICK] Running update tick on player (%s) at state [%d]\n", c.username, c.state)
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
		roomUID, err := queryRoomUID("forwell_inn_common_room")

		fmt.Printf("ROOMUID: %v \n", roomUID)
		if err != nil {
			fmt.Printf("ROOMUID err: %v \n", err)
			return nil
		}

		player := &Player{Name: username, CurrentRoom: roomUID, hp: 58, hpMax: 50, ap: 50, apMax: 58}
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

	found := false
	for i, conn := range server.playerList {
		if conn.Player.Name == connection.Player.Name {
			found = true
			server.playerList[i] = connection
			break
		}
	}
	if !found {
		server.playerList = append(server.playerList, connection)
	}
	connection.state = STATE_PLAYING

	connection.Write("The world darkens...\n")
	connection.Player.do("look")
	connection.Player.sendPrompt()
}

func (server *Server) AddConnection(connection net.Conn) *Connection {
	newConnection := Connection{conn: connection, timeConnected: time.Now()}
	server.connectionList = append(server.connectionList, &newConnection)
	go newConnection.listen()
	fmt.Printf("[CONN] There are %d connectioned sessions.\n", server.ConnectionCount())
	return &newConnection
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

	fmt.Printf("[DISC] There are %d connected sessions.\n", server.ConnectionCount())
}

func (server *Server) playerExited(connection *Connection, err error) {
	// delete the connection
	player := connection.username

	for i, conn := range server.playerList {
		if conn == connection {
			server.playerList[i] = server.playerList[len(server.playerList)-1]
			server.playerList[len(server.playerList)-1] = nil
			server.playerList = server.playerList[:len(server.playerList)-1]
			break
		}
	}

	fmt.Printf("[LOGOUT] %s logged out.\n", player)
}
