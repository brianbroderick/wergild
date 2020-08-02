package mud

import (
	"encoding/json"
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
		roomMap[room.UID].Desc = formatToWidth(roomMap[room.UID].Desc, 80)
		roomMap[room.UID].MobMap = make(map[string]*Mob)

		for j, mob := range roomMap[room.UID].Mobs {
			roomMap[room.UID].MobMap[mob.Slug] = &roomMap[room.UID].Mobs[j]
		}
	}

	return roomMap
}

func (room *Room) updateRoom(desc string) error {
	room.Desc = desc

	j, err := json.Marshal(room)
	if err != nil {
		return err
	}

	agora.MutateDgraph(j)
	return nil
}

func (room *Room) exitRoom(mob *Mob, direction string) {
	delete(room.MobMap, mob.Slug)
	str := mob.Name + " leaves " + strings.ToLower(direction) + ".\n"

	for _, mob := range room.MobMap {
		mob.yourMessageToChannel(str)
	}
}

func (room *Room) exitWorld(mob *Mob) {
	delete(room.MobMap, mob.Slug)
	str := mob.Name + " slowly fades away.\n"

	for _, mob := range room.MobMap {
		mob.yourMessageToChannel(str)
	}
}

func (room *Room) enterRoom(mob *Mob) {
	str := mob.Name + " arrives.\n"

	for _, mob := range room.MobMap {
		mob.yourMessageToChannel(str)
	}

	room.MobMap[mob.Slug] = mob
}

func (room *Room) showTo(mob *Mob) {
	str := room.Desc + room.listExits() + room.listContents() + room.listMobs(mob)
	mob.myMessageToChannel(str)
}

func (room *Room) showSmellTo(mob *Mob) {
	mob.myMessageToChannel(room.Smell)
}

func (room *Room) showListenTo(mob *Mob) {
	mob.myMessageToChannel(room.Listen)
}

func (room *Room) showEnv(mob *Mob) {
	if len(room.Env) == 0 {
		return
	}

	// Don't spam env messages. Make it happen 34% of the time for each pulse
	percent := rand.Intn(100)
	if percent > 34 {
		return
	}

	str := room.Env[rand.Intn(len(room.Env))] + "\n"
	mob.myMessageToChannel(str)
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
	if len(keys) > 0 {
		return "\n Obvious directions are:\n  " + strings.Join(keys, ", ") + "\n"
	}
	return "\n"
}

func (room *Room) listMobs(me *Mob) string {
	str := ""

	for _, mob := range room.MobMap {
		switch {
		case mob.Slug == me.Slug:
			// do nothing
		case mob.Title != "":
			str += mob.Name + ", " + mob.Title + ".\n"
		default:
			str += mob.Name + "\n"
		}
	}

	return str
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
			mobs {
				uid
				mobName
				mobTitle
				mobSlug
			}
		}
	}`

	var r DgraphResponse
	err := agora.ResolveQuery(&r, query)
	if err != nil {
		r.Errors = append(r.Errors, err)
	}

	rooms := r.Rooms
	if len(r.Rooms) == 0 {
		return []Room{}
	}
	// Populate ExitMap
	for i, room := range rooms {
		if rooms[i].ExitMap == nil {
			rooms[i].ExitMap = make(map[string]string)
		}

		for _, e := range room.Exits {
			if len(e.Dest) > 0 {
				rooms[i].ExitMap[e.Direction] = e.Dest[0].UID
			}
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
