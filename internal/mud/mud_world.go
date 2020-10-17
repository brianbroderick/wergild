package mud

import (
	"io/ioutil"

	"github.com/brianbroderick/logit"
)

var (
	ServerInstance *Server
	WorldInstance  *World
)

type World struct {
	roomList map[string]*Room
}

func init() {
	ServerInstance = &Server{
		mobs:  make(map[string]*Mob),
		users: make(map[string]*Mob),
	}
	WorldInstance = &World{
		roomList: make(map[string]*Room),
	}

	WorldInstance.BuildWorld()
	ServerInstance.Start()
}

func (world *World) BuildWorld() {
	if world == nil {
		world = &World{
			roomList: make(map[string]*Room),
		}
	}

	// 3. Load in the rooms
	logit.Info("[CONFIG] Loading rooms")
	world.roomList = loadRooms()
}

func (world *World) getRoom(roomId string) *Room {
	return world.roomList[roomId]
}

func motd() string {
	motdBytes, err := ioutil.ReadFile("welcome/1.txt")
	if err != nil {
		return ""
	} else {
		return string(motdBytes) + "\n"
	}
}
