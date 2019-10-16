package main

import (
	"encoding/json"
	"log"

	"github.com/brianbroderick/agora"
)

func schemaString() string {
	return `
	type Region {
		regionName: string
	}

	regionName: string @index(exact) . # world, building, dungeon, etc

	type Room {
		region: Region
		coorX: int
		coorY: int
		coorZ: int
		locationHash: string
		roomName: string
		roomSlug: string
		roomDesc: string
		listen: string
		smell: string
		lightLevel: int
		pointsOfInterest: [PointOfInterest]
		exits: [Exit]
		terrain: Terrain
		items: [Item]
	}

	region: uid @reverse .
	coorX: int @index(int) .
	coorY: int @index(int) .
	coorZ: int @index(int) .
	# locationHash: hash of region and coordinates
	locationHash: string @index(exact) @upsert . 
	roomSlug: string @index(exact) @upsert . 
	exits: [uid] . # no reverse - can have one way exits
	pointsOfInterest: [uid] .
	terrain: uid .
	items: [uid] .

	type Exit {
		dest:      Room   
		direction: string 
		portal:    string # is there a door?
	}

	direction: string @index(exact) .
		
	type PointOfInterest {
		keywords: [string] # full text?
		poiName: string
		poiDesc: string
		poiListen: string
		poiSmell: string
		search: [Item]
	}	

	keywords: string @index(exact) .

	type Terrain {
		terrainType: string
	}

	type Creature {
		creatureName: string
		creatureDesc: string
		age: int 
		lang: string
		level: int 
		exp: int 
		coins: int 
		bankCoins: int 
		hp: int
		hpMax: int
		ap: int
		apMax: int
		wimpy: int
		wimpyDir: string
		encumb: int
		sober: int
		thirst: int
		hunger: int
		poison: int
		defend: string
		aim: string
		attack: string
		str: int
		agl: int
		intl: int
		tgh: int
		per: int
		strMod: int
		aglMod: int
		intlMod: int
		tghMod: int
		perMod: int
		insertedCreatureAt: dateTime 
		items: [Item]
	}
		
	age: int @index(int) .
	level: int @index(int) .
	exp: int @index(int) .
	coins: int @index(int) .
	bankCoins: int @index(int) .
	insertedCreatureAt: dateTime @index(hour) .

	type Item {
		itemName: string
		itemDesc: string 
		coinValue: int
		weight: int
		items: [Item]
	}

	itemName: string @index(fulltext) .

	type CreatureClass {
		className: string
	}

	className: string @index(exact) @count .

	type CreatureRace {
		race: string
	}

	race: string @index(exact) @count .

	type Guild {
		guildName: string
	}

	guildName: string @index(exact) @count .

	`
}

func reloadData() {
	agora.DropAll()
	agora.SetSchema(schemaString())
	loadSeed()
}

func loadSeed() {
	// getSampleLens has a string arg because it can also
	// be used in a test resolver
	for _, p := range getSampleRooms() {
		j, err := json.Marshal(p)
		if err != nil {
			log.Fatal(err)
		}
		agora.MutateDgraph(j)
	}
}

func getSampleRooms() []Room {
	var region = Region{
		UID:        "_:region",
		RegionName: "Ancient Inn",
		Type:       "Region",
	}

	var northRoom = Room{
		UID:    "_:northRoom",
		Type:   "Room",
		Region: region,
		Slug:   "ancient_inn_north_room",
		Name:   "North Room",
		Desc:   "You are at a small room north of the inn's common room. A large firepit is the dominating feature here, casting warmth and powerful shadows across the tables and chairs arranged around the room. A large window to the northwest displays the forest outside.",
		Items: []Item{
			{
				Name: "a warm firepit",
				Desc: "The firepit is set halfway into the northern wall, spreading warmth throughout the inn.",
			},
		},
		Exits: []Exit{
			Exit{
				Dest:      Room{UID: "_:common"},
				Direction: "SOUTH",
				Portal:    "OPEN",
			}},
	}

	var commonRoom = Room{
		UID:    "_:common",
		Type:   "Room",
		Region: region,
		Slug:   "ancient_inn_common_room",
		Name:   "Common Room",
		Desc:   "You stand in the common room of the Ancient Inn of Tantallon. There are a number of chairs and tables scattered around the room, and there are two booths where people can go for private conversation. There is a large desk at the north end of the room, over which hangs an ornate clock. A doorway leads south into the world of Ancient Anguish and the adventure it has to offer.",
		Items: []Item{
			{
				Name: "oak chairs",
				Desc: "The chairs are sturdy, and made from oak. They complement the other furnishings nicely.",
			},
		},
		Exits: []Exit{
			Exit{
				Dest:      northRoom,
				Direction: "NORTH",
				Portal:    "OPEN",
			}},
	}

	return []Room{commonRoom}
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
