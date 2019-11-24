package main

import "github.com/brianbroderick/wergild/internal/mud"

func (world *worldSeed) setRooms() {
	world.commonBuildings()
	world.gates()
	world.inn()
	world.shops()
	world.centralCourtyard()
	world.streets()
	world.wells()
}

func (world *worldSeed) commonBuildings() {
	world.Rooms = append(world.Rooms, newRoom(quickRoom{
		name: "Blacksmith",
		desc: "You are in the main room of the blacksmith. Large iron tools line the walls. A forge and an anvil are the room's prominent features. There is a doorway to the south that leads to the central square.",
		env: []string{"Heat radiates from the forge.",
			"Hot embers glow in the forge.",
			"Orange and red flames flicker causing shadows to dance around the room."},
		region: "_:region_blacksmith",
		exits: []mud.Exit{
			mud.NewExit("_:central_courtyard_nw", "s"),
		},
		mobs: []mud.Mob{mud.Mob{UID: "_:erik"}},
	}))
}

func (world *worldSeed) gates() {
	world.Rooms = append(world.Rooms, newRoom(quickRoom{
		name: "Iridium Gate",
		desc: "Iridium Gate is an impressive portal between Forwell and the world beyond. The walls are painted silvery-white. The iron portcullis is closed preventing people from passing through.",
		x:    -4,
		y:    6,
		exits: []mud.Exit{
			mud.NewExit("_:iridium_and_rhodium", "s"),
		},
	}))

	world.Rooms = append(world.Rooms, newRoom(quickRoom{
		name: "Rhodium Gate",
		desc: "Rhodium Gate is an impressive silvery-white walled portal between Forwell and the world beyond. An iron portcullis is closed preventing people from passing through.",
		x:    1,
		y:    6,
	}))

	world.Rooms = append(world.Rooms, newRoom(quickRoom{
		name: "Palladium Gate",
		desc: "Palladium Gate consists of gleaming silvery-white walls. When it's open, its a portal between Forwell and the world beyond. Unfortunately, the iron portcullis is closed preventing people from passing through.",
		x:    6,
		y:    0,
	}))

	world.Rooms = append(world.Rooms, newRoom(quickRoom{
		name: "Tungston Gate",
		desc: "Tungston Gate is an impressive portal between Forwell and the world beyond. Though the walls are naturally gray, they are covered in lustrous green, yellow, and violet paint. The iron portcullis is closed preventing people from passing through.",
		x:    6,
		y:    -6,
	}))

	world.Rooms = append(world.Rooms, newRoom(quickRoom{
		name: "Ruthenium Gate",
		desc: "Ruthenium Gate is an impressive portal between Forwell and the world beyond. The walls are painted silvery-white. The iron portcullis is closed preventing people from passing through.",
		x:    5,
		y:    -9,
		exits: []mud.Exit{
			mud.NewExit("_:ruthenium_street_1", "s"),
		},
	}))

	world.Rooms = append(world.Rooms, newRoom(quickRoom{
		name: "Rhenium Gate",
		desc: "Rhenium Gate is an impressive portal between Forwell and the world beyond. The walls are painted silvery-gray. The iron portcullis is closed preventing people from passing through.",
		x:    0,
		y:    -11,
		exits: []mud.Exit{
			mud.NewExit("_:central_street_10", "n"),
		},
	}))

	world.Rooms = append(world.Rooms, newRoom(quickRoom{
		name: "Indium Gate",
		desc: "Indium Gate is an impressive portal between Forwell and the world beyond. The walls are painted silvery-white. The iron portcullis is closed preventing people from passing through.",
		x:    -6,
		y:    -11,
		exits: []mud.Exit{
			mud.NewExit("_:indium_street_16", "n"),
		},
	}))

	world.Rooms = append(world.Rooms, newRoom(quickRoom{
		name: "Platinum Gate",
		desc: "Platinum Gate is an impressive portal between Forwell and the world beyond. The walls are painted silvery-white though the iron portcullis is closed preventing people from passing through.",
		x:    -7,
		y:    -4,
		exits: []mud.Exit{
			mud.NewExit("_:indium_and_platinum", "e"),
		},
	}))

	world.Rooms = append(world.Rooms, newRoom(quickRoom{
		name: "Scandium Gate",
		desc: "Scandium Gate is an impressive portal between Forwell and the world beyond. The walls are painted silvery-white though the iron portcullis is closed preventing people from passing through.",
		x:    -7,
		y:    1,
		exits: []mud.Exit{
			mud.NewExit("_:indium_and_scandium", "e"),
		},
	}))
}

func (world *worldSeed) inn() {
	world.Rooms = append(world.Rooms, newRoom(quickRoom{
		name:   "Forwell Inn North Room",
		desc:   "You are at a small room north of the inn's common room. A large firepit is the dominating feature here, casting warmth and powerful shadows across the tables and chairs arranged around the room.",
		region: "_:region_inn",
		env: []string{"The logs burn softly in the pit.",
			"The fire crackles and sends up a few sparks.",
			"Warm embers glow in the firepit.",
			"Orange and red flames flicker causing shadows to dance around the room."},
		smell:  "The smell of burning charcoal and wood permeate the room.",
		listen: "The sounds of people laughing and chatting can be heard.",
		poi: []mud.PointOfInterest{
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
		},
		exits: []mud.Exit{
			mud.NewExit("_:forwell_inn_common_room", "s"),
		},
	}))

	world.Rooms = append(world.Rooms, newRoom(quickRoom{
		name:   "Forwell Inn Common Room",
		desc:   "You're in the common room of the inn where many start their adventures. There are a number of chairs and tables scattered around the room. There is a large desk at the north end of the room, over which hangs an ornate clock. A doorway leads south into the world of Wergild and the adventure it has to offer.",
		smell:  "The smell of rustic wood enters your nose.",
		listen: "The sounds of people chatting and discussing their travels can be heard.",
		poi: []mud.PointOfInterest{
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
		exits: []mud.Exit{
			mud.NewExit("_:forwell_inn_north_room", "n"),
			mud.NewExit("_:central_courtyard_n", "s"),
		},
	}))
}

func (world *worldSeed) shops() {
	world.Rooms = append(world.Rooms, newRoom(quickRoom{
		region:  "_:region_general_store",
		terrain: "_:t_shop",
		name:    "General Store",
		exits: []mud.Exit{
			mud.NewExit("_:central_courtyard_se", "w"),
		},
	}))

	world.Rooms = append(world.Rooms, newRoom(quickRoom{
		region:  "_:region_magic_shop",
		terrain: "_:t_shop",
		name:    "Magic Shop",
		exits: []mud.Exit{
			mud.NewExit("_:central_courtyard_e", "w"),
		},
	}))

	world.Rooms = append(world.Rooms, newRoom(quickRoom{
		region:  "_:region_map_shop",
		terrain: "_:t_shop",
		name:    "Map Shop",
		exits: []mud.Exit{
			mud.NewExit("_:central_courtyardNE", "s"),
		},
	}))

	world.Rooms = append(world.Rooms, newRoom(quickRoom{
		name:    "Trading Post",
		terrain: "_:t_shop",
		region:  "_:region_trading_post",
		exits: []mud.Exit{
			mud.NewExit("_:central_courtyard_ne", "w"),
		},
	}))
}

func (world *worldSeed) wells() {
	world.Rooms = append(world.Rooms, newRoom(quickRoom{
		name:    "Well 2",
		x:       0,
		y:       -6,
		terrain: "_:t_poi",
		exits: []mud.Exit{
			mud.NewExit("_:central_street_5", "n"),
			mud.NewExit("_:central_street_7", "s"),
		},
	}))

	world.Rooms = append(world.Rooms, newRoom(quickRoom{
		name:    "Well 3",
		x:       -2,
		y:       1,
		terrain: "_:t_poi",
		exits: []mud.Exit{
			mud.NewExit("_:central_courtyard_nw", "e"),
			mud.NewExit("_:scandium_street_4", "w"),
		},
	}))

	world.Rooms = append(world.Rooms, newRoom(quickRoom{
		name:    "Well 4",
		x:       -2,
		y:       -5,
		terrain: "_:t_poi",
	}))
}

type quickRoom struct {
	name    string
	desc    string
	region  string
	terrain string
	x       int
	y       int
	exits   []mud.Exit
	env     []string
	mobs    []mud.Mob
	poi     []mud.PointOfInterest
	smell   string
	listen  string
}

func newRoom(quick quickRoom) mud.Room {
	if quick.desc == "" {
		quick.desc = quick.name
	}

	if quick.region == "" {
		quick.region = "_:region_forwell"
	}

	if quick.terrain == "" {
		quick.terrain = "_:t_building"
	}

	slug := mud.ToSlug(quick.name)

	return mud.Room{
		UID:              "_:" + slug,
		Type:             "Room",
		Region:           mud.Region{UID: quick.region},
		Terrain:          mud.Terrain{UID: quick.terrain},
		Name:             quick.name,
		Slug:             slug,
		Desc:             quick.desc,
		CoorX:            quick.x,
		CoorY:            quick.y,
		Exits:            quick.exits,
		PointsOfInterest: quick.poi,
		Smell:            quick.smell,
		Listen:           quick.listen,
		Mobs:             quick.mobs,
	}
}
