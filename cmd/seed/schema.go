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
	regionName: string @index(exact) . 

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

	roomName: string .
	roomDesc: string .
	roomListen: string .
	roomSmell: string .
	lightLevel: int .

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

	dest: uid  .
	direction: string @index(exact) .
	portal: string .
		
	type PointOfInterest {
		poiName: string
		poiDesc: string
		poiListen: string
		poiSmell: string
		poiTouch: string
		search: [Item]
	}	

	poiName: string @index(fulltext) .
	poiDesc: string .
	poiListen: string .
	poiSmell: string .
	poiTouch: string .
	search: [uid] .

	type Terrain {
		terrainName: string
	}

	terrainName: string .

	type User {
		userName: string
		pass: password 
		email: string
	}

	userName: string @index(exact) @upsert .
	pass: password .
	mob: uid @reverse .
	email: string .

	type Mob {
		user: User
		mobName: string
		mobDesc: string
		mobSlug: string
		mobTitle: string # prince, duke
		mobRank: string # the bold, the brave, etc
		level: int 
		exp: int 
		coins: int 
		hp: int
		hpMax: int
		ap: int
		apMax: int
		str: int
		agl: int
		intl: int
		tgh: int
		per: int
		insertedMobAt: dateTime 
		items: [Item]
		conversations: [Conversation]
	}

	mobName: string .
	mobDesc: string .
	mobSlug: string .
	mobTitle: string .
	mobRank: string .
	hp: int .
	hpMax: int .
	ap: int .
	apMax: int .
	str: int .
	agl: int .
	intl: int .
	tgh: int .
	per: int .
	conversations: [uid] .
	
	user: uid @reverse .
	mobSlug: string @index(exact) @upsert . 
	level: int @index(int) .
	exp: int @index(int) .
	coins: int @index(int) .
	bankCoins: int @index(int) .
	insertedMobAt: dateTime @index(hour) .

	type Conversation {
		trigger: string 
		say:     string 
	}

	trigger: string .
	say: string .

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
	itemDesc: string .
	itemListen: string .
	itemSmell: string .
	itemTouch: string .
	coinValue: int .
	weight: int .

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
