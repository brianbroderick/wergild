package mud

import "fmt"

type World struct {
	roomList map[string]*Room
}

func BuildWorld() *World {
	if WorldInstance == nil {
		WorldInstance = &World{
			roomList: make(map[string]*Room),
		}
	}

	// 3. Load in the rooms
	fmt.Println("[CONFIG] Loading rooms")
	WorldInstance.roomList = loadRooms()

	return WorldInstance
}

func (world *World) getRoom(roomId string) *Room {
	return world.roomList[roomId]
}
