package main

func schemaString() string {
	return `
	type Room {
		region: string
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
	}

	region: string @index(exact) .
	coorX: int @index(int) .
	coorY: int @index(int) .
	coorZ: int @index(int) .
	# locationHash: hash of region and coordinates
	locationHash: string @index(exact) @upsert . 
	exits: [uid] . # no reverse - can have one way exits
	pointsOfInterest: [uid] .
	terrain: [uid] .
	direction: string .
		
	type PointOfInterest {
		keywords: [string]
		poiName: string
		poiDesc: string
		poiListen: string
		poiSmell: string
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
	}
		
	age: int @index(int) .
	level: int @index(int) .
	exp: int @index(int) .
	coins: int @index(int) .
	bankCoins: int @index(int) .
	insertedCreatureAt: dateTime @index(hour) .

	# item: uid @reverse @count .
	# class: uid @reverse @count .
	# guild: uid @reverse @count .
	# race: uid @reverse @count .

	`
}
