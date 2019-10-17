package main

import (
	"fmt"
	"strings"
)

/**
 * Load rooms from database
 */
func loadRooms() map[string]*Room {
	roomMap := make(map[string]*Room)

	rooms := queryAllRooms()

	for i, room := range rooms {
		roomMap[room.UID] = &rooms[i]
	}

	return roomMap
}

func (room *Room) showTo(player *Player) {
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
	keys := make([]string, len(room.Exits))

	i := 0
	for _, k := range room.Exits {
		keys[i] = strings.ToLower(k.Direction)
		i++
	}
	return "\n Obvious directions are:\n  " + strings.Join(keys, ", ") + "\n"
}

func queryAllRooms() []Room {
	query := `{
		room(func: type(Room)) {
			uid
			roomName 
			roomDesc
			exits {
				direction
				dest {
					uid
				}
			}
		}
	}`

	rooms := resolveQuery(query).Rooms

	// Populate ExitMap
	for i, room := range rooms {
		if rooms[i].ExitMap == nil {
			rooms[i].ExitMap = make(map[string]string)
		}

		for _, e := range room.Exits {
			rooms[i].ExitMap[e.Direction] = e.Dest[0].UID
		}
	}
	return rooms
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

	variables := make(map[string]string)
	variables["$slug"] = slug

	r := resolveQueryWithVars(query, variables)
	return r.Rooms
}

func queryPointOfInterest(uid string, q string) ([]PointOfInterest, error) {
	query := `query POI($uid: string, $q: string){
				room(func: uid($uid)) {
				  pointsOfInterest @filter(alloftext(poiName, $q)) {
					poiDesc
				  }
				}
			  }`

	variables := make(map[string]string)
	variables["$uid"] = uid
	variables["$q"] = q

	r := resolveQueryWithVars(query, variables)
	if r.Errors != nil {
		return nil, r.Errors[0]
	}
	if len(r.Rooms) > 0 {
		return r.Rooms[0].PointsOfInterest, nil
	}

	return nil, nil
}
