package mud

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/brianbroderick/agora"
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

func (room *Room) showSmellTo(player *Player) {
	player.connection.Write(room.Smell)
}

func (room *Room) showListenTo(player *Player) {
	player.connection.Write(room.Listen)
}

func (room *Room) showEnv(player *Player) {
	if len(room.Env) == 0 {
		return
	}

	// Don't spam env messages. Make it happen 25% of the time for each pulse
	percent := rand.Intn(100)
	if percent > 25 {
		return
	}

	player.connection.Write(room.Env[rand.Intn(len(room.Env))] + "\n")
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
			roomSmell
			roomListen
			roomEnv
			exits {
				direction
				dest {
					uid
				}
			}
		}
	}`

	var r DgraphResponse
	err := agora.ResolveQuery(&r, query)
	if err != nil {
		r.Errors = append(r.Errors, err)
	}

	rooms := r.Rooms
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

	var r DgraphResponse
	err := agora.ResolveQueryWithVars(&r, query, variables)
	if err != nil {
		return "", err
	}

	rooms := r.Rooms
	if len(rooms) > 0 {
		return rooms[0].UID, nil
	}
	return "", fmt.Errorf("Room not found")
}

func queryRoom(slug string) ([]Room, error) {
	query := `query Room($slug: string){
		room(first:1, func: eq(roomSlug, $slug)) {
			uid
			roomName 
			roomDesc
			roomSmell
			roomListen
			items {
				itemName
		        itemDesc 
			}
	}`

	variables := make(map[string]string)
	variables["$slug"] = slug

	var r DgraphResponse
	err := agora.ResolveQueryWithVars(&r, query, variables)
	if err != nil {
		return nil, err
	}

	return r.Rooms, nil
}

func queryPointOfInterest(uid string, q string) ([]PointOfInterest, error) {
	query := `query POI($uid: string, $q: string){
				room(func: uid($uid)) {
				  pointsOfInterest @filter(alloftext(poiName, $q)) {
					poiDesc
				  }
				}
			  }`

	return poiQueries(query, uid, q)
}

func queryPointOfInterestListen(uid string, q string) ([]PointOfInterest, error) {
	query := `query POI($uid: string, $q: string){
				room(func: uid($uid)) {
				  pointsOfInterest @filter(alloftext(poiName, $q)) {
					poiListen
				  }
				}
			  }`

	return poiQueries(query, uid, q)
}

func queryPointOfInterestSmell(uid string, q string) ([]PointOfInterest, error) {
	query := `query POI($uid: string, $q: string){
				room(func: uid($uid)) {
				  pointsOfInterest @filter(alloftext(poiName, $q)) {
					poiSmell
				  }
				}
			  }`

	return poiQueries(query, uid, q)
}

func queryPointOfInterestTouch(uid string, q string) ([]PointOfInterest, error) {
	query := `query POI($uid: string, $q: string){
				room(func: uid($uid)) {
				  pointsOfInterest @filter(alloftext(poiName, $q)) {
					poiTouch
				  }
				}
			  }`

	return poiQueries(query, uid, q)
}

func poiQueries(query string, uid string, q string) ([]PointOfInterest, error) {
	variables := make(map[string]string)
	variables["$uid"] = uid
	variables["$q"] = q

	var r DgraphResponse
	err := agora.ResolveQueryWithVars(&r, query, variables)
	if err != nil {
		return nil, err
	}

	if len(r.Rooms) > 0 {
		return r.Rooms[0].PointsOfInterest, nil
	}

	return nil, nil
}
