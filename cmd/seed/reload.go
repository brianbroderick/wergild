package main

import (
	"encoding/json"
	"log"

	"github.com/brianbroderick/agora"
	"github.com/brianbroderick/wergild/internal/login"
	"github.com/brianbroderick/wergild/internal/mud"
)

func reloadData() {
	agora.DropAll()
	agora.SetSchema(schemaString())
	loadSeed()
}

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
		roomEnv: string
		roomListen: string
		roomSmell: string
		lightLevel: int
		pointsOfInterest: [PointOfInterest]
		exits: [Exit]
		terrain: Terrain
		items: [Item]
		mobs: [Mob]
	}

	region: uid @reverse .
	coorX: int @index(int) .
	coorY: int @index(int) .
	coorZ: int @index(int) .
	# locationHash: hash of region and coordinates
	locationHash: string @index(exact) @upsert . 
	roomSlug: string @index(exact) @upsert . 
	roomEnv: [string] .
	exits: [uid] . # no reverse - can have one way exits
	pointsOfInterest: [uid] @reverse .
	mobs: [uid] @reverse .
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
		poiTouch: string
		search: [Item]
	}	

	poiName: string @index(fulltext) .

	type Terrain {
		terrainType: string
	}

	type User {
		userName: string
		pass: password 
		email: string
	}

	userName: string @index(exact) @upsert .
	pass: password .
	mob: uid @reverse .

	type Mob {
		user: User
		mobName: string
		mobDesc: string
		mobSlug: string
		mobTitle: string # prince, duke
		mobRank: string # the bold, the brave, etc
		age: int 
		lang: string
		gender: string
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
		insertedMobAt: dateTime 
		items: [Item]
	}
	
	user: uid @reverse .
	mobSlug: string @index(exact) @upsert . 
	age: int @index(int) .
	level: int @index(int) .
	exp: int @index(int) .
	coins: int @index(int) .
	bankCoins: int @index(int) .
	insertedMobAt: dateTime @index(hour) .

	type Item {
		itemHash: string
		itemName: string
		itemDesc: string 
		itemListen: string
		itemSmell: string
		itemTouch: string
		coinValue: int
		weight: int
		items: [Item]
	}

	itemName: string @index(fulltext) .
	itemHash: string @index(exact) @upsert . 

	type mobClass {
		className: string
	}

	className: string @index(exact) @count .

	type mobRace {
		race: string
	}

	race: string @index(exact) @count .

	type Guild {
		guildName: string
	}

	guildName: string @index(exact) @count .

	`
}

type worldSeed struct {
	Regions []mud.Region `json:"regions,omitempty"`
	Rooms   []mud.Room   `json:"rooms,omitempty"`
	Mobs    []mud.Mob    `json:"mobs,omitempty"`
	Users   []login.User `json:"users,omitempty"`
}

func combineStructs() worldSeed {
	seed := worldSeed{}
	seed.Regions = getRegions()
	seed.Rooms = getRooms()
	seed.Mobs = getMobs()
	return seed
}

func loadSeed() {
	j, err := json.Marshal(combineStructs())
	if err != nil {
		log.Fatal(err)
	}

	agora.MutateDgraph(j)
}

func getRegions() []mud.Region {
	var forwell = mud.Region{
		UID:        "_:region_forwell",
		RegionName: "Forwell",
		Type:       "Region",
	}

	var inn = mud.Region{
		UID:        "_:region_inn",
		RegionName: "Forwell Inn",
		Type:       "Region",
	}

	var blacksmithRegion = mud.Region{
		UID:        "_:region_blacksmith",
		RegionName: "Blacksmith",
		Type:       "Region",
	}

	return []mud.Region{forwell, inn, blacksmithRegion}
}

func getRooms() []mud.Room {
	var blacksmith = mud.Room{
		UID:    "_:blacksmith",
		Type:   "Room",
		Region: mud.Region{UID: "_:region_blacksmith"},
		Slug:   "blacksmith",
		Name:   "Blacksmith",
		Desc:   "You are in the main room of the blacksmith. Large iron tools line the walls. A forge and an anvil are the room's prominent features. There is a doorway to the north that leads to Main street.",
		Env: []string{"Heat radiates from the forge.",
			"Hot embers glow in the forge.",
			"Orange and red flames flicker causing shadows to dance around the room."},
		Exits: []mud.Exit{
			mud.Exit{
				Dest:      []mud.Room{mud.Room{UID: "_:mainStreet"}},
				Direction: "NORTH",
				Portal:    "OPEN",
				Type:      "Exit",
			},
		},
		Mobs: []mud.Mob{mud.Mob{UID: "_:william"}},
	}

	var mainStreet = mud.Room{
		UID:    "_:mainStreet",
		Type:   "Room",
		Region: mud.Region{UID: "_:region_forwell"},
		Slug:   "forwell_main_street",
		Name:   "Main Street",
		Desc:   "You are standing on main street directly south of the Forwell Inn and directly north of the blacksmith.",
		Exits: []mud.Exit{
			mud.Exit{
				Dest:      []mud.Room{mud.Room{UID: "_:common"}},
				Direction: "NORTH",
				Portal:    "OPEN",
				Type:      "Exit",
			},
			mud.Exit{
				Dest:      []mud.Room{mud.Room{UID: "_:blacksmith"}},
				Direction: "SOUTH",
				Portal:    "OPEN",
				Type:      "Exit",
			},
		},
	}

	var northRoom = mud.Room{
		UID:    "_:northRoom",
		Type:   "Room",
		Region: mud.Region{UID: "_:region_inn"},
		Slug:   "forwell_inn_north_room",
		Name:   "North Room",
		Desc:   "You are at a small room north of the inn's common room. A large firepit is the dominating feature here, casting warmth and powerful shadows across the tables and chairs arranged around the room. A large window to the northwest displays the forest outside.",
		Smell:  "The smell of burning charcoal and wood permeate the room.",
		Listen: "The sounds of people laughing and chatting can be heard.",
		Env: []string{"The logs burn softly in the pit.",
			"The fire crackles and sends up a few sparks.",
			"Warm embers glow in the firepit.",
			"Orange and red flames flicker causing shadows to dance around the room."},
		PointsOfInterest: []mud.PointOfInterest{
			{
				Name:   "large firepit",
				Desc:   "The firepit is set halfway into the northern wall, spreading warmth throughout the inn.",
				Touch:  "The firepit is much too hot to touch.",
				Listen: "The sounds of the crackling fire can be heard.",
				Smell:  "The aroma of burning charcoal and wood drift to your nose.",
				Type:   "PointOfInterest",
			},
			{
				Name:  "walnut chairs",
				Desc:  "The chairs are ornate, and made from walnut. They are neatly arranged with tables to give the room a tidy look.",
				Touch: "The chair is smooth and feels well made.",
				Type:  "PointOfInterest",
			},
			{
				Name:  "walnut tables",
				Desc:  "The walnut tables are well crafted with maple inlays.",
				Touch: "The table is smooth and feels well made.",
				Type:  "PointOfInterest",
			},
			{
				Name:   "large northwest window",
				Desc:   "Peering through the window reveals a beautiful forest outside.",
				Listen: "Birds can be faintly heard out the window.",
				Touch:  "The window is smooth. You better not let the innkeeper know you left prints on the window!",
				Type:   "PointOfInterest",
			},
		},
		Exits: []mud.Exit{
			mud.Exit{
				Dest:      []mud.Room{mud.Room{UID: "_:common"}},
				Direction: "SOUTH",
				Portal:    "OPEN",
				Type:      "Exit",
			}},
	}

	var commonRoom = mud.Room{
		UID:    "_:common",
		Type:   "Room",
		Region: mud.Region{UID: "_:region_inn"},
		Slug:   "forwell_inn_common_room",
		Name:   "Common Room",
		Desc:   "You stand in the common room of the Ancient Inn of Forwell. There are a number of chairs and tables scattered around the room. There is a large desk at the north end of the room, over which hangs an ornate clock. A doorway leads south into the world of Wergild and the adventure it has to offer.",
		Smell:  "The smell of rustic wood enters your nose.",
		Listen: "The sounds of people chatting and discussing their travels can be heard.",
		PointsOfInterest: []mud.PointOfInterest{
			{
				Name:  "oak chairs",
				Desc:  "The chairs are sturdy, and made from oak. They complement the other furnishings nicely.",
				Touch: "The chair is smooth to the touch.",
				Type:  "PointOfInterest",
			},
			{
				Name:  "antique oak tables",
				Desc:  "The antique oak tables are sturdy and well used. If they could talk, oh the stories they might tell.",
				Touch: "The table is smooth to the touch.",
				Type:  "PointOfInterest",
			},
			{
				Name:  "large desk",
				Desc:  "The desk is old, yet still maintains its luster. Various papers and writing implements can be seen.",
				Touch: "The desk is rough from much use.",
				Type:  "PointOfInterest",
			},
			{
				Name:  "writing implements",
				Desc:  "On the desk are various pens used to maintain records of the Inn's activity.",
				Smell: "The ink has a sort of sweet smell. You wonder what it's made of.",
				Touch: "The pen is smooth to the touch.",
				Type:  "PointOfInterest",
			},
			{
				Name:  "papers",
				Desc:  "The papers on the desk look to be a log of people currently staying at the Inn.",
				Smell: "The papers have a distinct smell of old parchment.",
				Touch: "The papers are textured like parchment.",
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
				Listen: "Constant ticking sounds can be heard from the clock.",
				Touch:  "You probably shouldn't touch the clock. The innkeeper might get upset.",
				Type:   "PointOfInterest",
			},
		},
		Exits: []mud.Exit{
			mud.Exit{
				Dest:      []mud.Room{mud.Room{UID: "_:northRoom"}},
				Direction: "NORTH",
				Portal:    "OPEN",
				Type:      "Exit",
			},
			mud.Exit{
				Dest:      []mud.Room{mud.Room{UID: "_:mainStreet"}},
				Direction: "SOUTH",
				Portal:    "OPEN",
				Type:      "Exit",
			},
		},
	}

	return []mud.Room{blacksmith, mainStreet, northRoom, commonRoom}
}

func getMobs() []mud.Mob {
	william := mud.Mob{
		UID:   "_:william",
		Name:  "William",
		Slug:  "william",
		Title: "the apprentice smith",
		Desc:  "This tall lad already shows the stature of his uncle in sheer height. Bulging muscles have started to fill out his gangly frame, and his skin has taken on the dark hue of blacksmiths everywhere. His dark hair is clipped into a short burr, showing a faded scar near his temple. He has kind, merry brown eyes and you sense that he may have a soft spot for those new to this world.",
		Type:  "Mob",
	}

	azkul := mud.Mob{
		UID:   "_:azkul",
		Name:  "Azkul",
		Slug:  "azkul",
		Title: "the utter novice",
		Type:  "Mob",
		Ap:    50,
		Hp:    50,
		ApMax: 50,
		HpMax: 50,
		User: login.User{
			UID:  "_:azkulUser",
			Name: "azkul",
			Pass: "123456",
			Type: "User",
		},
	}

	return []mud.Mob{william, azkul}
}
