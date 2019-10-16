package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/brianbroderick/agora"
)

/**
 * Load rooms from database
 * @TODO Actually load them from database...
 */

func loadRooms() map[string]*Room {
	roomMap := make(map[string]*Room)

	rooms := queryAllRooms()
	for _, room := range rooms {
		roomMap[room.UID] = &room
	}

	return roomMap
}

// func loadRooms() map[int]*Room {
// 	rooms := make(map[int]*Room, 5)
// 	rooms[1] = &Room{
// 		UID:  "1",
// 		Name: "Ancient Inn common room",
// 		Desc: "You stand in the common room of the Ancient Inn of Tantallon. There are a number of chairs and tables scattered around the room, and there are two booths where people can go for private conversation. There is a large desk at the north end of the room, over which hangs an ornate clock. A doorway leads south into the world of Ancient Anguish and the adventure it has to offer.",
// 		Items: []Item{
// 			{
// 				Name: "Oak Chairs",
// 				Desc: "The chairs are sturdy, and made from oak. They complement the other furnishings nicely.",
// 			},
// 		},
// 		// Exits: map[string]*Exit{"NORTH": &Exit{1, 2}},
// 	}

// 	rooms[2] = &Room{
// 		UID:  "2",
// 		Name: "Ancient Inn north room",
// 		Desc: "You are at a small room north of the inn's common room. A large firepit is the dominating feature here, casting warmth and powerful shadows across the tables and chairs arranged around the room. A large window to the northwest displays the forest outside.",
// 		Items: []Item{
// 			{
// 				Name: "a warm firepit",
// 				Desc: "The firepit is set halfway into the northern wall, spreading warmth throughout the inn.",
// 			},
// 		},
// 		// Exits: map[string]*Exit{"SOUTH": &Exit{2, 1}},
// 	}

// 	return rooms
// }

func (room *Room) showTo(player *Player) {
	//room.title + "\n\n" +
	str := room.Desc + room.listExits() + room.listContents()
	player.connection.Write(str)
}

func (room *Room) listContents() string {
	str := ""
	for _, item := range room.Items {
		str += item.Desc + "\n"
	}
	return str
}

func (room *Room) listExits() string {
	return "listExits"
	// keys := make([]string, len(room.Exits))

	// i := 0
	// for k := range room.Exits {
	// 	keys[i] = k
	// 	i++
	// }
	// return "\n Obvious directions are:\n  " + strings.Join(keys, ", ") + "\n"
}

func queryAllRooms() []Room {
	query := `{room(func: type(Room)) {
			uid roomName roomDesc items {
				itemName itemDesc 
			}
			}}`

	return resolveQuery(query).Rooms
}

func queryRoomUID(slug string) (string, error) {
	query := `query Room($slug: string){
		        room(first:1, func: eq(roomSlug, $slug)) {
				  uid 
				}
			  }`

	variables := make(map[string]string)
	variables["$slug"] = slug

	r := resolveQueryWithVars(query, variables)
	if r.Errors != nil {
		return "", r.Errors[0]
	}

	rooms := r.Rooms
	if len(rooms) > 0 {
		return rooms[0].UID, nil
	}

	return "", fmt.Errorf("Room not found")
}

func queryRoom(slug string) []Room {
	// slug := "ancient_inn_common_room"
	query := `query Room($slug: string){
		room(first:1, func: eq(roomSlug, $slug)) {
			uid
			roomName 
			roomDesc
			items {
				itemName
		        itemDesc 
			}
	}`

	// rooms[1] = &Room{
	// 		UID:  "1",
	// 		Name: "Ancient Inn common room",
	// 		Desc: "You stand in the common room of the Ancient Inn of Tantallon. There are a number of chairs and tables scattered around the room, and there are two booths where people can go for private conversation. There is a large desk at the north end of the room, over which hangs an ornate clock. A doorway leads south into the world of Ancient Anguish and the adventure it has to offer.",
	// 		Items: []Item{
	// 			{
	// 				Name: "Oak Chairs",
	// 				Desc: "The chairs are sturdy, and made from oak. They complement the other furnishings nicely.",
	// 			},
	// 		},
	// 		// Exits: map[string]*Exit{"NORTH": &Exit{1, 2}},

	variables := make(map[string]string)
	variables["$slug"] = slug

	r := resolveQueryWithVars(query, variables)
	return r.Rooms
}

func resolveQuery(query string) DgraphResponse {
	j := agora.QueryDgraph(query)

	var r DgraphResponse
	err := json.Unmarshal(j, &r)
	if err != nil {
		log.Fatal(err)
	}

	return r
}

func resolveQueryWithVars(query string, variables map[string]string) DgraphResponse {
	j := agora.QueryDgraphWithVars(query, variables)

	var r DgraphResponse
	err := json.Unmarshal(j, &r)
	if err != nil {
		r.Errors = append(r.Errors, err)
	}

	return r
}
