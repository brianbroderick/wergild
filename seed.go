package main

import (
	"github.com/brianbroderick/agora"
)

func schemaString() string {
	return `
	type Region {
		regionName: string
		rooms: [Room]
	}

	regionName: string @index(exact) . # world, building, dungeon, etc
	rooms: [uid] @reverse .

	type Room {
		coorX: int
		coorY: int
		coorZ: int
		locationHash: string
		roomName: string
		roomDesc: string
		listen: string
		smell: string
		lightLevel: int
		pointsOfInterest: [PointOfInterest]
		exits: [Room]
		terrain: Terrain
		items: [Item]
	}

	coorX: int @index(int) .
	coorY: int @index(int) .
	coorZ: int @index(int) .
	# locationHash: hash of region and coordinates
	locationHash: string @index(exact) @upsert . 
	exits: [uid] . # no reverse - can have one way exits
	pointsOfInterest: [uid] .
	terrain: uid .
	items: [uid] .
	direction: string .
		
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
		items: [Item]
	}

	itemName: string @index(exact) .

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
	// loadSeed()
}

// func loadSeed() {
// 	// getSampleLens has a string arg because it can also
// 	// be used in a test resolver
// 	for _, p := range getSampleLens("test") {
// 		j, err := json.Marshal(p)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		agora.MutateDgraph(j)
// 	}
// }
