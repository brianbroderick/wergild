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
	pointsOfInterest: [uid] @reverse .
	terrain: uid .
	items: [uid] .

	type Exit {
		dest:      Room   
		direction: string 
		portal:    string # is there a door?
	}

	direction: string @index(exact) .
		
	type PointOfInterest {
		poiName: string
		poiDesc: string
		poiListen: string
		poiSmell: string
		search: [Item]
	}	

	poiName: string @index(fulltext) .

	type Terrain {
		terrainType: string
	}

	type Creature {
		creatureName: string
		creatureDesc: string
		creatureSlug: string
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
	
	creatureSlug: string @index(exact) @upsert . 
	age: int @index(int) .
	level: int @index(int) .
	exp: int @index(int) .
	coins: int @index(int) .
	bankCoins: int @index(int) .
	insertedCreatureAt: dateTime @index(hour) .

	type Item {
		itemName: string
		itemDesc: string 
		itemListen: string
		itemSmell: string
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
		// fmt.Printf("%s \n", j)
		agora.MutateDgraph(j)
	}
}

func getSampleRooms() []Room {
	var region = Region{
		UID:        "_:region",
		RegionName: "Forwell Inn",
		Type:       "Region",
	}

	var northRoom = Room{
		UID:    "_:northRoom",
		Type:   "Room",
		Region: region,
		Slug:   "forwell_inn_north_room",
		Name:   "North Room",
		Desc:   "You are at a small room north of the inn's common room. A large firepit is the dominating feature here, casting warmth and powerful shadows across the tables and chairs arranged around the room. A large window to the northwest displays the forest outside.",
		Smell:  "The smell of charcoal and wood permeate the room.",
		Listen: "The sounds of people laughing and chatting can be heard.",
		PointsOfInterest: []PointOfInterest{
			{
				Name: "large firepit",
				Desc: "The firepit is set halfway into the northern wall, spreading warmth throughout the inn.",
				Type: "PointOfInterest",
			},
			{
				Name: "walnut chairs",
				Desc: "The chairs are ornate, and made from walnut. They are neatly arranged with tables to give the room a tidy look.",
				Type: "PointOfInterest",
			},
			{
				Name: "walnut tables",
				Desc: "The walnut tables are well crafted with maple inlays.",
				Type: "PointOfInterest",
			},
			{
				Name: "large northwest window",
				Desc: "Peering through the window reveals a beautiful forest outside.",
				Type: "PointOfInterest",
			},
		},
		Exits: []Exit{
			Exit{
				Dest:      []Room{Room{UID: "_:common"}},
				Direction: "SOUTH",
				Portal:    "OPEN",
			}},
	}

	var commonRoom = Room{
		UID:    "_:common",
		Type:   "Room",
		Region: region,
		Slug:   "forwell_inn_common_room",
		Name:   "Common Room",
		Desc:   "You stand in the common room of the Ancient Inn of Forwell. There are a number of chairs and tables scattered around the room. There is a large desk at the north end of the room, over which hangs an ornate clock. A doorway leads south into the world of Wergild and the adventure it has to offer.",
		Smell:  "The smell of rustic wood enters your nose.",
		Listen: "The sounds of people chatting and discussing their travels can be heard.",
		PointsOfInterest: []PointOfInterest{
			{
				Name: "oak chairs",
				Desc: "The chairs are sturdy, and made from oak. They complement the other furnishings nicely.",
				Type: "PointOfInterest",
			},
			{
				Name: "antique oak tables",
				Desc: "The antique oak tables are sturdy and well used. If they could talk, oh the stories they might tell.",
				Type: "PointOfInterest",
			},
			{
				Name: "large desk",
				Desc: "The desk is old, yet still maintains its luster. Various papers and writing implements can be seen.",
				Type: "PointOfInterest",
			},
			{
				Name:  "writing implements",
				Desc:  "On the desk are various pens used to maintain records of the Inn's activity.",
				Smell: "The ink has a sort of sweet smell. You wonder what it's made of.",
				Type:  "PointOfInterest",
			},
			{
				Name:  "papers",
				Desc:  "The papers on the desk look to be a log of people currently staying at the Inn.",
				Smell: "The papers have a distinct smell of old parchment.",
				Type:  "PointOfInterest",
			},
			{
				Name: "log",
				Desc: "The log indicates that someone by the name of Viktor recently stayed here.",
				Type: "PointOfInterest",
			},
			{
				Name:   "ornate clock",
				Desc:   "The clock is well crafted and reads 9:17 PM",
				Listen: "Constant ticking sounds can be heard from the clock",
				Type:   "PointOfInterest",
			},
		},
		Exits: []Exit{
			Exit{
				Dest:      []Room{northRoom},
				Direction: "NORTH",
				Portal:    "OPEN",
			}},
	}

	return []Room{commonRoom}
}
