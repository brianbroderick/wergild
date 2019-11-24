package main

import "github.com/brianbroderick/wergild/internal/mud"

func (world *worldSeed) streets() {
	// Central Street
	world.Rooms = append(world.Rooms, newStreet(quickRoom{
		name: "Central Street 2",
		x:    0,
		y:    -2,
		exits: []mud.Exit{
			mud.NewExit("_:central_courtyard_s", "n"),
			mud.NewExit("_:central_street_3", "s"),
		},
	}))

	world.Rooms = append(world.Rooms, newStreet(quickRoom{
		name: "Central Street 3",
		x:    0,
		y:    -3,
		exits: []mud.Exit{
			mud.NewExit("_:central_street_2", "n"),
			mud.NewExit("_:central_street_4", "s"),
		},
	}))

	world.Rooms = append(world.Rooms, newStreet(quickRoom{
		name: "Central Street 4",
		x:    0,
		y:    -4,
		exits: []mud.Exit{
			mud.NewExit("_:central_street_3", "n"),
			mud.NewExit("_:central_street_5", "s"),
		},
	}))

	world.Rooms = append(world.Rooms, newStreet(quickRoom{
		name: "Central Street 5",
		x:    0,
		y:    -5,
		exits: []mud.Exit{
			mud.NewExit("_:central_street_4", "n"),
			mud.NewExit("_:well_2", "s"),
		},
	}))

	world.Rooms = append(world.Rooms, newStreet(quickRoom{
		name: "Central Street 7",
		x:    0,
		y:    -7,
		exits: []mud.Exit{
			mud.NewExit("_:well_2", "n"),
			mud.NewExit("_:central_street_8", "s"),
		},
	}))

	world.Rooms = append(world.Rooms, newStreet(quickRoom{
		name: "Central Street 8",
		x:    0,
		y:    -8,
		exits: []mud.Exit{
			mud.NewExit("_:central_street_7", "n"),
			mud.NewExit("_:central_street_9", "s"),
		},
	}))

	world.Rooms = append(world.Rooms, newStreet(quickRoom{
		name: "Central Street 9",
		x:    0,
		y:    -9,
		exits: []mud.Exit{
			mud.NewExit("_:central_street_8", "n"),
			mud.NewExit("_:central_street_10", "s"),
		},
	}))

	world.Rooms = append(world.Rooms, newStreet(quickRoom{
		name: "Central Street 10",
		x:    0,
		y:    -10,
		exits: []mud.Exit{
			mud.NewExit("_:central_street_9", "n"),
			mud.NewExit("_:rhenium_gate", "s"),
		},
	}))

	// Scandium Street

	world.Rooms = append(world.Rooms, newStreet(quickRoom{
		name: "Scandium Street 1",
		x:    -6,
		y:    1,
		exits: []mud.Exit{
			mud.NewExit("_:scandium_street_2", "e"),
			mud.NewExit("_:scandium_gate", "w"),
		},
	}))

	world.Rooms = append(world.Rooms, newStreet(quickRoom{
		name: "Scandium Street 2",
		x:    -5,
		y:    1,
		exits: []mud.Exit{
			mud.NewExit("_:iridium_and_scandium", "e"),
			mud.NewExit("_:scandium_street_1", "w"),
		},
	}))

	world.Rooms = append(world.Rooms, newStreet(quickRoom{
		name: "Iridium and Scandium",
		x:    -4,
		y:    1,
		exits: []mud.Exit{
			mud.NewExit("_:iridium_street_4", "n"),
			mud.NewExit("_:scandium_street_4", "e"),
			mud.NewExit("_:iridium_street_6", "s"),
			mud.NewExit("_:scandium_street_2", "w"),
		},
	}))

	world.Rooms = append(world.Rooms, newStreet(quickRoom{
		name: "Scandium Street 4",
		x:    -4,
		y:    1,
		exits: []mud.Exit{
			mud.NewExit("_:well_3", "e"),
			mud.NewExit("_:iridium_and_scandium", "w"),
		},
	}))

	// Iridium Street

	world.Rooms = append(world.Rooms, newStreet(quickRoom{
		name: "Iridium Street 1",
		x:    -4,
		y:    5,
		exits: []mud.Exit{
			mud.NewExit("_:iridium_gate", "n"),
			mud.NewExit("_:iridium_street_2", "s"),
		},
	}))

	world.Rooms = append(world.Rooms, newStreet(quickRoom{
		name: "Iridium Street 2",
		x:    -4,
		y:    4,
		exits: []mud.Exit{
			mud.NewExit("_:iridium_street_1", "n"),
			mud.NewExit("_:iridium_street_3", "s"),
		},
	}))

	world.Rooms = append(world.Rooms, newStreet(quickRoom{
		name: "Iridium Street 3",
		x:    -4,
		y:    3,
		exits: []mud.Exit{
			mud.NewExit("_:iridium_street_2", "n"),
			mud.NewExit("_:iridium_street_4", "s"),
		},
	}))

	world.Rooms = append(world.Rooms, newStreet(quickRoom{
		name: "Iridium Street 4",
		x:    -4,
		y:    2,
		exits: []mud.Exit{
			mud.NewExit("_:iridium_street_3", "n"),
			mud.NewExit("_:iridium_and_scandium", "s"),
		},
	}))

	world.Rooms = append(world.Rooms, newStreet(quickRoom{
		name: "Iridium Street 6",
		x:    -4,
		y:    0,
		exits: []mud.Exit{
			mud.NewExit("_:iridium_and_scandium", "n"),
			mud.NewExit("_:iridium_street_7", "s"),
		},
	}))

	world.Rooms = append(world.Rooms, newStreet(quickRoom{
		name: "Iridium Street 7",
		x:    -4,
		y:    -1,
		exits: []mud.Exit{
			mud.NewExit("_:iridium_street_6", "n"),
			mud.NewExit("_:iridium_street_8", "s"),
		},
	}))

	world.Rooms = append(world.Rooms, newStreet(quickRoom{
		name: "Iridium Street 8",
		x:    -4,
		y:    -1,
		exits: []mud.Exit{
			mud.NewExit("_:iridium_street_7", "n"),
			mud.NewExit("_:iridium_street_9", "e"),
		},
	}))

	world.Rooms = append(world.Rooms, newStreet(quickRoom{
		name: "Iridium Street 9",
		x:    -4,
		y:    -1,
		exits: []mud.Exit{
			mud.NewExit("_:iridium_street_8", "w"),
			mud.NewExit("_:iridium_street_10", "e"),
		},
	}))

	world.Rooms = append(world.Rooms, newStreet(quickRoom{
		name: "Iridium Street 10",
		x:    -4,
		y:    -1,
		exits: []mud.Exit{
			mud.NewExit("_:iridium_street_11", "s"),
			mud.NewExit("_:iridium_street_9", "w"),
		},
	}))

	world.Rooms = append(world.Rooms, newStreet(quickRoom{
		name: "Iridium Street 11",
		x:    -4,
		y:    -1,
		exits: []mud.Exit{
			mud.NewExit("_:iridium_street_10", "n"),
			mud.NewExit("_:iridium_street_12", "s"),
		},
	}))

	world.Rooms = append(world.Rooms, newStreet(quickRoom{
		name: "Iridium Street 12",
		x:    -4,
		y:    -1,
		exits: []mud.Exit{
			mud.NewExit("_:iridium_street_11", "n"),
			mud.NewExit("_:iridium_and_osmium", "s"),
		},
	}))

	world.Rooms = append(world.Rooms, newStreet(quickRoom{
		name: "Iridium and Osmium",
		x:    -4,
		y:    -1,
		exits: []mud.Exit{
			mud.NewExit("_:iridium_street_12", "n"),
			mud.NewExit("_:iridium_street_14", "s"),
		},
	}))

	world.Rooms = append(world.Rooms, newStreet(quickRoom{
		name: "Iridium Street 14",
		x:    -4,
		y:    -1,
		exits: []mud.Exit{
			mud.NewExit("_:iridium_and_osmium", "n"),
			mud.NewExit("_:iridium_street_15", "s"),
		},
	}))

	world.Rooms = append(world.Rooms, newStreet(quickRoom{
		name: "Iridium Street 15",
		x:    -4,
		y:    -1,
		exits: []mud.Exit{
			mud.NewExit("_:iridium_street_14", "n"),
			// goes sw and se once those are options
		},
	}))
}

func (world *worldSeed) centralCourtyard() {
	world.Rooms = append(world.Rooms, newStreet(quickRoom{
		name: "Central Courtyard NW",
		x:    -1,
		y:    1,
		exits: []mud.Exit{
			mud.NewExit("_:blacksmith", "n"),
			mud.NewExit("_:central_courtyard_n", "e"),
			mud.NewExit("_:central_courtyard_w", "s"),
			mud.NewExit("_:well_3", "w"),
		},
	}))

	world.Rooms = append(world.Rooms, newStreet(quickRoom{
		name: "Central Courtyard N",
		x:    0,
		y:    1,
		exits: []mud.Exit{
			mud.NewExit("_:forwell_inn_common_room", "n"),
			mud.NewExit("_:central_courtyard_ne", "e"),
			mud.NewExit("_:central_courtyard_c", "s"),
			mud.NewExit("_:central_courtyard_nw", "w"),
		},
	}))

	world.Rooms = append(world.Rooms, newStreet(quickRoom{
		name: "Central Courtyard NE",
		x:    1,
		y:    1,
		exits: []mud.Exit{
			mud.NewExit("_:map_shop", "n"),
			mud.NewExit("_:trading_post", "e"),
			mud.NewExit("_:central_courtyard_e", "s"),
			mud.NewExit("_:central_courtyard_n", "w"),
		},
	}))

	world.Rooms = append(world.Rooms, newStreet(quickRoom{
		name: "Central Courtyard W",
		x:    -1,
		y:    0,
		exits: []mud.Exit{
			mud.NewExit("_:central_courtyard_nw", "n"),
			mud.NewExit("_:central_courtyard_c", "e"),
			mud.NewExit("_:central_courtyard_sw", "s"),
		},
	}))

	world.Rooms = append(world.Rooms, newStreet(quickRoom{
		name: "Central Courtyard C",
		x:    0,
		y:    0,
		exits: []mud.Exit{
			mud.NewExit("_:central_courtyard_n", "n"),
			mud.NewExit("_:central_courtyard_e", "e"),
			mud.NewExit("_:central_courtyard_s", "s"),
			mud.NewExit("_:central_courtyard_w", "w"),
		},
	}))

	world.Rooms = append(world.Rooms, newStreet(quickRoom{
		name: "Central Courtyard E",
		x:    1,
		y:    0,
		exits: []mud.Exit{
			mud.NewExit("_:central_courtyard_ne", "n"),
			mud.NewExit("_:magic_shop", "e"),
			mud.NewExit("_:central_courtyard_se", "s"),
			mud.NewExit("_:central_courtyard_c", "w"),
		},
	}))

	world.Rooms = append(world.Rooms, newStreet(quickRoom{
		name: "Central Courtyard SW",
		x:    -1,
		y:    -1,
		exits: []mud.Exit{
			mud.NewExit("_:central_courtyard_w", "n"),
			mud.NewExit("_:central_courtyard_s", "e"),
		},
	}))

	world.Rooms = append(world.Rooms, newStreet(quickRoom{
		name: "Central Courtyard S",
		x:    0,
		y:    -1,
		exits: []mud.Exit{
			mud.NewExit("_:central_courtyard_c", "n"),
			mud.NewExit("_:central_courtyard_se", "e"),
			mud.NewExit("_:central_street_2", "s"),
			mud.NewExit("_:central_courtyard_sw", "w"),
		},
	}))

	world.Rooms = append(world.Rooms, newStreet(quickRoom{
		name: "Central Courtyard SE",
		x:    0,
		y:    -1,
		exits: []mud.Exit{
			mud.NewExit("_:central_courtyard_e", "n"),
			mud.NewExit("_:general_store", "e"),
			mud.NewExit("_:central_courtyard_s", "w"),
		},
	}))
}

func newStreet(quick quickRoom) mud.Room {
	s := newRoom(quick)
	s.Terrain = mud.Terrain{UID: "_:t_street"}

	if quick.desc == "" {
		quick.desc = quick.name
	}

	return s
}
