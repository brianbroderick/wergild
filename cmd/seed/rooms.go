package main

import "github.com/brianbroderick/wergild/internal/mud"

func (world *worldSeed) setRooms() {
	world.blacksmith()
	world.gates()
	world.generalStore()
	world.inn()
	world.magicShop()
	world.mapShop()
	world.openSquare()
	world.streets()
	world.tradingPost()
	world.wells()
}

func (world *worldSeed) blacksmith() {
	room := mud.Room{
		UID:     "_:blacksmith",
		Type:    "Room",
		Region:  mud.Region{UID: "_:region_blacksmith"},
		Terrain: mud.Terrain{UID: "_:t_building"},
		Slug:    "blacksmith",
		Name:    "Blacksmith",
		Desc:    "You are in the main room of the blacksmith. Large iron tools line the walls. A forge and an anvil are the room's prominent features. There is a doorway to the south that leads to the central square.",
		Env: []string{"Heat radiates from the forge.",
			"Hot embers glow in the forge.",
			"Orange and red flames flicker causing shadows to dance around the room."},
		Exits: []mud.Exit{
			mud.NewExit("_:open_squareNW", "s"),
		},
		Mobs: []mud.Mob{mud.Mob{UID: "_:erik"}},
	}

	world.Rooms = append(world.Rooms, room)
}

func (world *worldSeed) gates() {
	world.Rooms = append(world.Rooms, mud.Room{
		UID:     "_:iridium_gate",
		Type:    "Room",
		Region:  mud.Region{UID: "_:region_forwell"},
		Terrain: mud.Terrain{UID: "_:t_building"},
		Slug:    "iridium_gate",
		Name:    "Iridium Gate",
		Desc:    "Iridium Gate is an impressive portal between Forwell and the world beyond. The walls are painted silvery-white. The iron portcullis is closed preventing people from passing through.",
		CoorX:   -4,
		CoorY:   6,
		Exits:   []mud.Exit{
			// mud.NewExit("_:open_squareSE", "w"),
		},
	})

	world.Rooms = append(world.Rooms, mud.Room{
		UID:     "_:rhodium_gate",
		Type:    "Room",
		Region:  mud.Region{UID: "_:region_forwell"},
		Terrain: mud.Terrain{UID: "_:t_building"},
		Slug:    "rhodium_gate",
		Name:    "Rhodium Gate",
		Desc:    "Rhodium Gate is an impressive silvery-white walled portal between Forwell and the world beyond. An iron portcullis is closed preventing people from passing through.",
		CoorX:   1,
		CoorY:   6,
		Exits:   []mud.Exit{
			// mud.NewExit("_:open_squareSE", "w"),
		},
	})

	world.Rooms = append(world.Rooms, mud.Room{
		UID:     "_:palladium_gate",
		Type:    "Room",
		Region:  mud.Region{UID: "_:region_forwell"},
		Terrain: mud.Terrain{UID: "_:t_building"},
		Slug:    "palladium_gate",
		Name:    "Palladium Gate",
		Desc:    "Palladium Gate consists of gleaming silvery-white walls. When it's open, its a portal between Forwell and the world beyond. Unfortunately, the iron portcullis is closed preventing people from passing through.",
		CoorX:   6,
		CoorY:   0,
		Exits:   []mud.Exit{
			// mud.NewExit("_:open_squareSE", "w"),
		},
	})

	world.Rooms = append(world.Rooms, mud.Room{
		UID:     "_:tungston_gate",
		Type:    "Room",
		Region:  mud.Region{UID: "_:region_forwell"},
		Terrain: mud.Terrain{UID: "_:t_building"},
		Slug:    "tungston_gate",
		Name:    "Tungston Gate",
		Desc:    "Tungston Gate is an impressive portal between Forwell and the world beyond. Though the walls are naturally gray, they are covered in lustrous green, yellow, and violet paint. The iron portcullis is closed preventing people from passing through.",
		CoorX:   6,
		CoorY:   -6,
		Exits:   []mud.Exit{
			// mud.NewExit("_:open_squareSE", "w"),
		},
	})

	world.Rooms = append(world.Rooms, mud.Room{
		UID:     "_:ruthenium_gate",
		Type:    "Room",
		Region:  mud.Region{UID: "_:region_forwell"},
		Terrain: mud.Terrain{UID: "_:t_building"},
		Slug:    "ruthenium_gate",
		Name:    "Ruthenium Gate",
		Desc:    "Ruthenium Gate is an impressive portal between Forwell and the world beyond. The walls are painted silvery-white. The iron portcullis is closed preventing people from passing through.",
		CoorX:   5,
		CoorY:   -9,
		Exits:   []mud.Exit{
			// mud.NewExit("_:open_squareSE", "w"),
		},
	})

	world.Rooms = append(world.Rooms, mud.Room{
		UID:     "_:indium_gate",
		Type:    "Room",
		Region:  mud.Region{UID: "_:region_forwell"},
		Terrain: mud.Terrain{UID: "_:t_building"},
		Slug:    "indium_gate",
		Name:    "Indium Gate",
		Desc:    "Indium Gate is an impressive portal between Forwell and the world beyond. The walls are painted silvery-white. The iron portcullis is closed preventing people from passing through.",
		CoorX:   -6,
		CoorY:   -11,
		Exits:   []mud.Exit{
			// mud.NewExit("_:open_squareSE", "w"),
		},
	})

	world.Rooms = append(world.Rooms, mud.Room{
		UID:     "_:rhenium_gate",
		Type:    "Room",
		Region:  mud.Region{UID: "_:region_forwell"},
		Terrain: mud.Terrain{UID: "_:t_building"},
		Slug:    "rhenium_gate",
		Name:    "Rhenium Gate",
		Desc:    "Rhenium Gate is an impressive portal between Forwell and the world beyond. The walls are painted silvery-gray. The iron portcullis is closed preventing people from passing through.",
		CoorX:   0,
		CoorY:   -11,
		Exits:   []mud.Exit{
			// mud.NewExit("_:open_squareSE", "w"),
		},
	})
}

func (world *worldSeed) generalStore() {
	room := mud.Room{
		UID:     "_:general_store",
		Type:    "Room",
		Region:  mud.Region{UID: "_:region_general_store"},
		Terrain: mud.Terrain{UID: "_:t_shop"},
		Slug:    "general_store",
		Name:    "General Store",
		Desc:    "General Store",
		Exits: []mud.Exit{
			mud.NewExit("_:open_squareSE", "w"),
		},
	}

	world.Rooms = append(world.Rooms, room)
}

func (world *worldSeed) inn() {
	world.Rooms = append(world.Rooms, mud.Room{
		UID:     "_:northRoom",
		Type:    "Room",
		Region:  mud.Region{UID: "_:region_inn"},
		Terrain: mud.Terrain{UID: "_:t_building"},
		Slug:    "forwell_inn_north_room",
		Name:    "North Room",
		Desc:    "You are at a small room north of the inn's common room. A large firepit is the dominating feature here, casting warmth and powerful shadows across the tables and chairs arranged around the room.",
		Smell:   "The smell of burning charcoal and wood permeate the room.",
		Listen:  "The sounds of people laughing and chatting can be heard.",
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
		},
		Exits: []mud.Exit{
			mud.NewExit("_:inn_common", "s"),
		},
	})

	world.Rooms = append(world.Rooms, mud.Room{
		UID:     "_:inn_common",
		Type:    "Room",
		Region:  mud.Region{UID: "_:region_inn"},
		Terrain: mud.Terrain{UID: "_:t_building"},
		Slug:    "forwell_inn_common_room",
		Name:    "Common Room",
		Desc:    "You're in the common room of the inn where many start their adventures. There are a number of chairs and tables scattered around the room. There is a large desk at the north end of the room, over which hangs an ornate clock. A doorway leads south into the world of Wergild and the adventure it has to offer.",
		Smell:   "The smell of rustic wood enters your nose.",
		Listen:  "The sounds of people chatting and discussing their travels can be heard.",
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
			mud.NewExit("_:northRoom", "n"),
			mud.NewExit("_:open_squareN", "s"),
		},
	})
}

func (world *worldSeed) magicShop() {
	room := mud.Room{
		UID:     "_:magic_shop",
		Type:    "Room",
		Region:  mud.Region{UID: "_:region_magic_shop"},
		Terrain: mud.Terrain{UID: "_:t_shop"},
		Slug:    "magic_shop",
		Name:    "Magic Shop",
		Desc:    "Magic Shop",
		Exits: []mud.Exit{
			mud.NewExit("_:open_squareE", "w"),
		},
	}

	world.Rooms = append(world.Rooms, room)
}

func (world *worldSeed) mapShop() {
	room := mud.Room{
		UID:     "_:map_shop",
		Type:    "Room",
		Region:  mud.Region{UID: "_:region_map_shop"},
		Terrain: mud.Terrain{UID: "_:t_shop"},
		Slug:    "map_shop",
		Name:    "Map Shop",
		Desc:    "Map Shop",
		Exits: []mud.Exit{
			mud.NewExit("_:open_squareNE", "s"),
		},
	}

	world.Rooms = append(world.Rooms, room)
}

func (world *worldSeed) openSquare() {
	world.Rooms = append(world.Rooms, mud.Room{
		UID:     "_:open_squareNW",
		Type:    "Room",
		Region:  mud.Region{UID: "_:region_forwell"},
		Terrain: mud.Terrain{UID: "_:t_street"},
		Slug:    "forwell_open_square_nw",
		Name:    "Open Square, Northwest",
		Desc:    "Open Square, Northwest",
		CoorX:   -1,
		CoorY:   1,
		Exits: []mud.Exit{
			mud.NewExit("_:blacksmith", "n"),
			mud.NewExit("_:open_squareN", "e"),
			mud.NewExit("_:open_squareW", "s"),
			// mud.NewExit("_:well3", "w"),
		},
	})

	world.Rooms = append(world.Rooms, mud.Room{
		UID:     "_:open_squareN",
		Type:    "Room",
		Region:  mud.Region{UID: "_:region_forwell"},
		Terrain: mud.Terrain{UID: "_:t_street"},
		Slug:    "forwell_open_square_n",
		Name:    "Open Square, North",
		Desc:    "Open Square, North",
		CoorX:   0,
		CoorY:   1,
		Exits: []mud.Exit{
			mud.NewExit("_:inn_common", "n"),
			mud.NewExit("_:open_squareNE", "e"),
			mud.NewExit("_:open_squareC", "s"),
			mud.NewExit("_:open_squareNW", "w"),
		},
	})

	world.Rooms = append(world.Rooms, mud.Room{
		UID:     "_:open_squareNE",
		Type:    "Room",
		Region:  mud.Region{UID: "_:region_forwell"},
		Terrain: mud.Terrain{UID: "_:t_street"},
		Slug:    "forwell_open_square_ne",
		Name:    "Open Square, Northeast",
		Desc:    "Open Square, Northeast",
		CoorX:   1,
		CoorY:   1,
		Exits: []mud.Exit{
			mud.NewExit("_:map_shop", "n"),
			mud.NewExit("_:trading_post", "e"),
			mud.NewExit("_:open_squareE", "s"),
			mud.NewExit("_:open_squareN", "w"),
		},
	})

	world.Rooms = append(world.Rooms, mud.Room{
		UID:     "_:open_squareW",
		Type:    "Room",
		Region:  mud.Region{UID: "_:region_forwell"},
		Terrain: mud.Terrain{UID: "_:t_street"},
		Slug:    "forwell_open_square_w",
		Name:    "Open Square, West",
		Desc:    "Open Square, West",
		CoorX:   -1,
		CoorY:   0,
		Exits: []mud.Exit{
			mud.NewExit("_:open_squareNW", "n"),
			mud.NewExit("_:open_squareC", "e"),
			mud.NewExit("_:open_squareSW", "s"),
		},
	})

	world.Rooms = append(world.Rooms, mud.Room{
		UID:     "_:open_squareC",
		Type:    "Room",
		Region:  mud.Region{UID: "_:region_forwell"},
		Terrain: mud.Terrain{UID: "_:t_street"},
		Slug:    "forwell_open_square_c",
		Name:    "Open Square, Central",
		Desc:    "Open Square, Central",
		CoorX:   0,
		CoorY:   0,
		Exits: []mud.Exit{
			mud.NewExit("_:open_squareN", "n"),
			mud.NewExit("_:open_squareE", "e"),
			mud.NewExit("_:open_squareS", "s"),
			mud.NewExit("_:open_squareW", "w"),
		},
	})

	world.Rooms = append(world.Rooms, mud.Room{
		UID:     "_:open_squareE",
		Type:    "Room",
		Region:  mud.Region{UID: "_:region_forwell"},
		Terrain: mud.Terrain{UID: "_:t_street"},
		Slug:    "forwell_open_square_e",
		Name:    "Open Square, East",
		Desc:    "Open Square, East",
		CoorX:   1,
		CoorY:   0,
		Exits: []mud.Exit{
			mud.NewExit("_:open_squareNE", "n"),
			mud.NewExit("_:magic_shop", "e"),
			mud.NewExit("_:open_squareSE", "s"),
			mud.NewExit("_:open_squareC", "w"),
		},
	})

	world.Rooms = append(world.Rooms, mud.Room{
		UID:     "_:open_squareSW",
		Type:    "Room",
		Region:  mud.Region{UID: "_:region_forwell"},
		Terrain: mud.Terrain{UID: "_:t_street"},
		Slug:    "forwell_open_square_sw",
		Name:    "Open Square, Southwest",
		Desc:    "Open Square, Southwest",
		CoorX:   -1,
		CoorY:   -1,
		Exits: []mud.Exit{
			mud.NewExit("_:open_squareW", "n"),
			mud.NewExit("_:open_squareS", "e"),
		},
	})

	world.Rooms = append(world.Rooms, mud.Room{
		UID:     "_:open_squareS",
		Type:    "Room",
		Region:  mud.Region{UID: "_:region_forwell"},
		Terrain: mud.Terrain{UID: "_:t_street"},
		Slug:    "forwell_open_square_s",
		Name:    "Open Square, South",
		Desc:    "Open Square, South",
		CoorX:   0,
		CoorY:   -1,
		Exits: []mud.Exit{
			mud.NewExit("_:open_squareC", "n"),
			mud.NewExit("_:open_squareSE", "e"),
			// mud.NewExit("_:south_road", "s"), // name this road
			mud.NewExit("_:open_squareSW", "w"),
		},
	})

	world.Rooms = append(world.Rooms, mud.Room{
		UID:     "_:open_squareSE",
		Type:    "Room",
		Region:  mud.Region{UID: "_:region_forwell"},
		Terrain: mud.Terrain{UID: "_:t_street"},
		Slug:    "forwell_open_square_se",
		Name:    "Open Square, Southeast",
		Desc:    "Open Square, Southeast",
		CoorX:   1,
		CoorY:   -1,
		Exits: []mud.Exit{
			mud.NewExit("_:open_squareE", "n"),
			mud.NewExit("_:general_store", "e"),
			mud.NewExit("_:open_squareS", "w"),
		},
	})
}

func (world *worldSeed) streets() {
}

func (world *worldSeed) tradingPost() {
	room := mud.Room{
		UID:     "_:trading_post",
		Type:    "Room",
		Region:  mud.Region{UID: "_:region_trading_post"},
		Terrain: mud.Terrain{UID: "_:t_shop"},
		Slug:    "trading_post",
		Name:    "Trading Post",
		Desc:    "Trading Post",
		Exits: []mud.Exit{
			mud.NewExit("_:open_squareNE", "w"),
		},
	}

	world.Rooms = append(world.Rooms, room)
}

func (world *worldSeed) wells() {
	world.Rooms = append(world.Rooms, mud.Room{
		UID:     "_:well2",
		Type:    "Room",
		Region:  mud.Region{UID: "_:region_forwell"},
		Terrain: mud.Terrain{UID: "_:t_poi"},
		Slug:    "well2",
		Name:    "Well 2",
		Desc:    "Well 2",
		CoorX:   0,
		CoorY:   -6,
		Exits:   []mud.Exit{
			// mud.NewExit("_:open_squareNW", "s"),
		},
		Mobs: []mud.Mob{mud.Mob{UID: "_:william"}},
	})

	world.Rooms = append(world.Rooms, mud.Room{
		UID:     "_:well3",
		Type:    "Room",
		Region:  mud.Region{UID: "_:region_forwell"},
		Terrain: mud.Terrain{UID: "_:t_poi"},
		Slug:    "well3",
		Name:    "Well 3",
		Desc:    "Well 3",
		CoorX:   -2,
		CoorY:   1,
		Exits:   []mud.Exit{
			// mud.NewExit("_:open_squareNW", "s"),
		},
	})

	world.Rooms = append(world.Rooms, mud.Room{
		UID:     "_:well4",
		Type:    "Room",
		Region:  mud.Region{UID: "_:region_forwell"},
		Terrain: mud.Terrain{UID: "_:t_poi"},
		Slug:    "well4",
		Name:    "Well 4",
		Desc:    "Well 4",
		CoorX:   -2,
		CoorY:   -5,
		Exits:   []mud.Exit{
			// mud.NewExit("_:open_squareNW", "s"),
		},
	})
}
