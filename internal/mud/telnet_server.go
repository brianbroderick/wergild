package mud

import (
	"fmt"
	"net"
	"time"

	"github.com/brianbroderick/wergild/internal/login"
)

type Server struct {
	ticker *time.Ticker
	mobs   map[string]*Mob
	users  map[string]*Mob
}

func (server *Server) AddConnection(connection net.Conn) {
	newConnection := Connection{conn: connection, timeConnected: time.Now(), user: login.User{}}
	go newConnection.listen()

	fmt.Println("[CONN] Received")
}

func (server *Server) onPlayerAuthenticated(connection *Connection) {
	roomUID, err := queryRoomUID("forwell_inn_common_room")
	if err != nil {
		connection.hasError(err)
		return
	}
	name := connection.user.Name

	if _, ok := server.mobs[name]; ok {
		server.mobs[name].conn.Write("You've signed in elsewhere, so we're logging you off here.\n")
		server.mobs[name].conn.conn.Close()
		connection.Write("We've closed your other connection.\n")
	} else {
		connection.Write("The world darkens...\n")
	}

	m, err := queryMob(name)
	m.CurrentRoom = roomUID
	m.conn = connection
	m.cmd = make(chan string, 1)
	m.me = make(chan string, 1)
	m.you = make(chan string, 1)

	server.mobs[name] = &m
	server.users[name] = &m
	connection.state = STATE_PLAYING
	fmt.Printf("[AUTH] There are %d connected sessions.\n", server.usersCount())

	server.onCommandReceived(connection, "look")
	WorldInstance.roomList[roomUID].enterRoom(&m)
}

func (server *Server) onCommandReceived(conn *Connection, message string) {
	select {
	case server.mobs[conn.user.Name].cmd <- message:
		server.Command(conn.user.Name)
	default:
		server.mobs[conn.user.Name].conn.Write("Slow down. Your last command was not processed.\n")
	}
}

func (server *Server) Command(k string) {
	select {
	case msg := <-server.mobs[k].cmd:
		if len(msg) == 0 {
			server.mobs[k].sendPrompt()
			return
		}

		server.mobs[k].do(msg)
		server.mobs[k].sendPrompt()

	default:
		server.mobs[k].conn.Write("Slow down. Your last command was not processed.\n")
	}
}

func (server *Server) onClientConnectionClosed(connnection *Connection, err error) {
	// delete the connection & remove from maps
	connnection.conn.Close()
	delete(server.mobs, connnection.user.Name)
	delete(server.users, connnection.user.Name)
	fmt.Printf("[DISC] There are %d connected sessions.\n", server.usersCount())
}

func (server *Server) usersCount() int {
	return len(server.users)
}

func (server *Server) Start() {
	fmt.Println("[SERVER] Ticker Started")
	server.ticker = time.NewTicker(time.Millisecond * 8000)

	go func() {
		for range server.ticker.C {
			for _, m := range server.users {
				m.pulseUpdate()
			}
		}
	}()
}
