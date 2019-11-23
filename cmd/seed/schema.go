package main

func schemaString() string {
	return `
	type Region {
		regionName: string
		partOf: Region
		coorX: int
		coorY: int
		coorZ: int
	}

	partOf: uid @reverse .
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
		terrainName: string
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
		insertedMobAt: dateTime 
		items: [Item]
		conversations: [Conversation]
	}
	
	user: uid @reverse .
	mobSlug: string @index(exact) @upsert . 
	age: int @index(int) .
	level: int @index(int) .
	exp: int @index(int) .
	coins: int @index(int) .
	bankCoins: int @index(int) .
	insertedMobAt: dateTime @index(hour) .

	type Conversation {
		trigger: string 
		say:     string 
	}

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
