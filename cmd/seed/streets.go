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
			mud.NewExit("_:rhenium_road_6", "w"),
		},
	}))

	// Scandium Street

	world.Rooms = append(world.Rooms, newStreet(quickRoom{
		name: "Indium and Scandium",
		x:    -6,
		y:    1,
		exits: []mud.Exit{
			mud.NewExit("_:indium_street_5", "n"),
			mud.NewExit("_:scandium_street_2", "e"),
			mud.NewExit("_:indium_street_7", "s"),
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
		x:    -3,
		y:    1,
		exits: []mud.Exit{
			mud.NewExit("_:well_3", "e"),
			mud.NewExit("_:iridium_and_scandium", "w"),
		},
	}))

	world.Rooms = append(world.Rooms, newStreet(quickRoom{
		name: "Scandium Street 6",
		x:    -2,
		y:    2,
		exits: []mud.Exit{
			mud.NewExit("_:scandium_street_7", "n"),
			mud.NewExit("_:well_3", "s"),
		},
	}))

	world.Rooms = append(world.Rooms, newStreet(quickRoom{
		name: "Scandium Street 7",
		x:    -2,
		y:    3,
		exits: []mud.Exit{
			mud.NewExit("_:scandium_street_8", "e"),
			mud.NewExit("_:scandium_street_6", "s"),
		},
	}))

	world.Rooms = append(world.Rooms, newStreet(quickRoom{
		name: "Scandium Street 8",
		x:    -1,
		y:    3,
		exits: []mud.Exit{
			mud.NewExit("_:scandium_street_9", "e"),
			mud.NewExit("_:scandium_street_7", "w"),
		},
	}))

	world.Rooms = append(world.Rooms, newStreet(quickRoom{
		name: "Scandium Street 9",
		x:    0,
		y:    3,
		exits: []mud.Exit{
			mud.NewExit("_:ruthenium_street_3", "e"),
			mud.NewExit("_:scandium_street_8", "w"),
		},
	}))

	// Ruthenium Street

	world.Rooms = append(world.Rooms, newStreet(quickRoom{
		name: "Ruthenium Street 1",
		x:    1,
		y:    5,
		exits: []mud.Exit{
			mud.NewExit("_:ruthenium_gate", "n"),
			mud.NewExit("_:rhodium_street_7", "w"),
			// mud.NewExit("_:ruthenium_street_2", "s"),
		},
	}))

	// Rhodium Street

	world.Rooms = append(world.Rooms, newStreet(quickRoom{
		name: "Rhodium Street 1",
		x:    -6,
		y:    5,
		exits: []mud.Exit{
			mud.NewExit("_:rhodium_street_2", "e"),
			mud.NewExit("_:indium_street_1", "w"),
		},
	}))

	world.Rooms = append(world.Rooms, newStreet(quickRoom{
		name: "Rhodium Street 2",
		x:    -5,
		y:    5,
		exits: []mud.Exit{
			mud.NewExit("_:iridium_and_rhodium", "e"),
			mud.NewExit("_:rhodium_street_1", "w"),
		},
	}))

	world.Rooms = append(world.Rooms, newStreet(quickRoom{
		name: "Rhodium Street 4",
		x:    -3,
		y:    5,
		exits: []mud.Exit{
			mud.NewExit("_:rhodium_street_5", "e"),
			mud.NewExit("_:iridium_and_rhodium", "w"),
		},
	}))

	world.Rooms = append(world.Rooms, newStreet(quickRoom{
		name: "Rhodium Street 5",
		x:    -2,
		y:    5,
		exits: []mud.Exit{
			mud.NewExit("_:rhodium_street_6", "e"),
			mud.NewExit("_:rhodium_street_4", "w"),
		},
	}))

	world.Rooms = append(world.Rooms, newStreet(quickRoom{
		name: "Rhodium Street 6",
		x:    -1,
		y:    5,
		exits: []mud.Exit{
			mud.NewExit("_:rhodium_street_7", "e"),
			mud.NewExit("_:rhodium_street_5", "w"),
		},
	}))

	world.Rooms = append(world.Rooms, newStreet(quickRoom{
		name: "Rhodium Street 7",
		x:    0,
		y:    5,
		exits: []mud.Exit{
			mud.NewExit("_:ruthenium_street_1", "e"),
			mud.NewExit("_:rhodium_street_6", "w"),
		},
	}))

	// Iridium Street

	world.Rooms = append(world.Rooms, newStreet(quickRoom{
		name: "Iridium and Rhodium",
		x:    -4,
		y:    5,
		exits: []mud.Exit{
			mud.NewExit("_:iridium_gate", "n"),
			mud.NewExit("_:rhodium_street_4", "e"),
			mud.NewExit("_:iridium_street_2", "s"),
			mud.NewExit("_:rhodium_street_2", "w"),
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
		y:    -2,
		exits: []mud.Exit{
			mud.NewExit("_:iridium_street_7", "n"),
			mud.NewExit("_:iridium_street_9", "e"),
		},
	}))

	world.Rooms = append(world.Rooms, newStreet(quickRoom{
		name: "Iridium Street 9",
		x:    -3,
		y:    -2,
		exits: []mud.Exit{
			mud.NewExit("_:iridium_street_8", "w"),
			mud.NewExit("_:iridium_street_10", "e"),
		},
	}))

	world.Rooms = append(world.Rooms, newStreet(quickRoom{
		name: "Iridium Street 10",
		x:    -2,
		y:    -2,
		exits: []mud.Exit{
			mud.NewExit("_:iridium_street_11", "s"),
			mud.NewExit("_:iridium_street_9", "w"),
		},
	}))

	world.Rooms = append(world.Rooms, newStreet(quickRoom{
		name: "Iridium Street 11",
		x:    -2,
		y:    -3,
		exits: []mud.Exit{
			mud.NewExit("_:iridium_street_10", "n"),
			mud.NewExit("_:iridium_street_12", "s"),
		},
	}))

	world.Rooms = append(world.Rooms, newStreet(quickRoom{
		name: "Iridium Street 12",
		x:    -2,
		y:    -4,
		exits: []mud.Exit{
			mud.NewExit("_:iridium_street_11", "n"),
			mud.NewExit("_:iridium_and_platinum", "s"),
		},
	}))

	world.Rooms = append(world.Rooms, newStreet(quickRoom{
		name: "Iridium and Platinum",
		x:    -2,
		y:    -5,
		exits: []mud.Exit{
			mud.NewExit("_:iridium_street_12", "n"),
			mud.NewExit("_:iridium_street_14", "s"),
		},
	}))

	world.Rooms = append(world.Rooms, newStreet(quickRoom{
		name: "Iridium Street 14",
		x:    -2,
		y:    -6,
		exits: []mud.Exit{
			mud.NewExit("_:iridium_and_platinum", "n"),
			mud.NewExit("_:iridium_street_15", "s"),
		},
	}))

	world.Rooms = append(world.Rooms, newStreet(quickRoom{
		name: "Iridium Street 15",
		x:    -2,
		y:    -7,
		exits: []mud.Exit{
			mud.NewExit("_:iridium_street_14", "n"),
			mud.NewExit("_:iridium_street_16", "n"),
		},
	}))

	world.Rooms = append(world.Rooms, newStreet(quickRoom{
		name: "Iridium Street 16",
		x:    -2,
		y:    -8,
		exits: []mud.Exit{
			mud.NewExit("_:iridium_street_15", "n"),
		},
	}))

	// Indium Street

	world.Rooms = append(world.Rooms, newStreet(quickRoom{
		name: "Indium Street 1",
		x:    -7,
		y:    5,
		exits: []mud.Exit{
			mud.NewExit("_:rhodium_street_1", "e"),
			mud.NewExit("_:indium_street_2", "s"),
		},
	}))

	world.Rooms = append(world.Rooms, newStreet(quickRoom{
		name: "Indium Street 2",
		x:    -7,
		y:    4,
		exits: []mud.Exit{
			mud.NewExit("_:indium_street_1", "n"),
			mud.NewExit("_:indium_street_3", "s"),
		},
	}))

	world.Rooms = append(world.Rooms, newStreet(quickRoom{
		name: "Indium Street 3",
		x:    -7,
		y:    3,
		exits: []mud.Exit{
			mud.NewExit("_:indium_street_2", "n"),
			mud.NewExit("_:indium_street_4", "e"),
		},
	}))

	world.Rooms = append(world.Rooms, newStreet(quickRoom{
		name: "Indium Street 4",
		x:    -6,
		y:    3,
		exits: []mud.Exit{
			mud.NewExit("_:indium_street_5", "s"),
			mud.NewExit("_:indium_street_3", "w"),
		},
	}))

	world.Rooms = append(world.Rooms, newStreet(quickRoom{
		name: "Indium Street 5",
		x:    -6,
		y:    2,
		exits: []mud.Exit{
			mud.NewExit("_:indium_street_4", "n"),
			mud.NewExit("_:indium_and_scandium", "s"),
		},
	}))

	world.Rooms = append(world.Rooms, newStreet(quickRoom{
		name: "Indium Street 7",
		x:    -6,
		y:    0,
		exits: []mud.Exit{
			mud.NewExit("_:indium_and_scandium", "n"),
			mud.NewExit("_:indium_street_8", "s"),
		},
	}))

	world.Rooms = append(world.Rooms, newStreet(quickRoom{
		name: "Indium Street 8",
		x:    -6,
		y:    -1,
		exits: []mud.Exit{
			mud.NewExit("_:indium_street_7", "n"),
			mud.NewExit("_:indium_street_9", "s"),
		},
	}))

	world.Rooms = append(world.Rooms, newStreet(quickRoom{
		name: "Indium Street 9",
		x:    -6,
		y:    -2,
		exits: []mud.Exit{
			mud.NewExit("_:indium_street_8", "n"),
			mud.NewExit("_:indium_street_10", "s"),
		},
	}))

	world.Rooms = append(world.Rooms, newStreet(quickRoom{
		name: "Indium Street 10",
		x:    -6,
		y:    -3,
		exits: []mud.Exit{
			mud.NewExit("_:indium_street_9", "n"),
			mud.NewExit("_:indium_and_platinum", "s"),
		},
	}))

	world.Rooms = append(world.Rooms, newStreet(quickRoom{
		name: "Indium and Platinum",
		x:    -6,
		y:    -4,
		exits: []mud.Exit{
			mud.NewExit("_:indium_street_10", "n"),
			mud.NewExit("_:platinum_street_2", "e"),
			mud.NewExit("_:indium_street_12", "s"),
			mud.NewExit("_:platinum_gate", "w"),
		},
	}))

	world.Rooms = append(world.Rooms, newStreet(quickRoom{
		name: "Indium Street 12",
		x:    -6,
		y:    -5,
		exits: []mud.Exit{
			mud.NewExit("_:indium_and_platinum", "n"),
			mud.NewExit("_:indium_street_13", "s"),
		},
	}))

	world.Rooms = append(world.Rooms, newStreet(quickRoom{
		name: "Indium Street 13",
		x:    -6,
		y:    -6,
		exits: []mud.Exit{
			mud.NewExit("_:indium_street_12", "n"),
			mud.NewExit("_:indium_street_14", "s"),
		},
	}))

	world.Rooms = append(world.Rooms, newStreet(quickRoom{
		name: "Indium Street 14",
		x:    -6,
		y:    -7,
		exits: []mud.Exit{
			mud.NewExit("_:indium_street_13", "n"),
			mud.NewExit("_:indium_street_15", "s"),
		},
	}))

	world.Rooms = append(world.Rooms, newStreet(quickRoom{
		name: "Indium Street 15",
		x:    -6,
		y:    -8,
		exits: []mud.Exit{
			mud.NewExit("_:indium_street_14", "n"),
			mud.NewExit("_:indium_street_16", "s"),
		},
	}))

	world.Rooms = append(world.Rooms, newStreet(quickRoom{
		name: "Indium Street 16",
		x:    -6,
		y:    -9,
		exits: []mud.Exit{
			mud.NewExit("_:indium_street_15", "n"),
			mud.NewExit("_:rhenium_road_1", "e"),
			mud.NewExit("_:indium_street_17", "s"),
		},
	}))

	world.Rooms = append(world.Rooms, newStreet(quickRoom{
		name: "Indium Street 17",
		x:    -6,
		y:    -10,
		exits: []mud.Exit{
			mud.NewExit("_:indium_street_16", "n"),
			mud.NewExit("_:indium_gate", "s"),
		},
	}))

	// Rhenium Street

	world.Rooms = append(world.Rooms, newStreet(quickRoom{
		name: "Rhenium Road 1",
		x:    -5,
		y:    -9,
		exits: []mud.Exit{
			mud.NewExit("_:rhenium_road_2", "e"),
			mud.NewExit("_:indium_street_16", "w"),
		},
	}))

	world.Rooms = append(world.Rooms, newStreet(quickRoom{
		name: "Rhenium Road 2",
		x:    -4,
		y:    -9,
		exits: []mud.Exit{
			mud.NewExit("_:hafnium_street_1", "n"),
			mud.NewExit("_:rhenium_road_3", "s"),
			mud.NewExit("_:rhenium_road_1", "w"),
		},
	}))

	world.Rooms = append(world.Rooms, newStreet(quickRoom{
		name: "Rhenium Road 3",
		x:    -4,
		y:    -10,
		exits: []mud.Exit{
			mud.NewExit("_:rhenium_road_2", "n"),
			mud.NewExit("_:rhenium_road_4", "e"),
		},
	}))

	world.Rooms = append(world.Rooms, newStreet(quickRoom{
		name: "Rhenium Road 4",
		x:    -3,
		y:    -10,
		exits: []mud.Exit{
			mud.NewExit("_:rhenium_road_5", "e"),
			mud.NewExit("_:rhenium_road_3", "w"),
		},
	}))

	world.Rooms = append(world.Rooms, newStreet(quickRoom{
		name: "Rhenium Road 5",
		x:    -2,
		y:    -10,
		exits: []mud.Exit{
			mud.NewExit("_:rhenium_road_6", "e"),
			mud.NewExit("_:rhenium_road_4", "w"),
		},
	}))

	world.Rooms = append(world.Rooms, newStreet(quickRoom{
		name: "Rhenium Road 6",
		x:    -1,
		y:    -10,
		exits: []mud.Exit{
			mud.NewExit("_:central_street_10", "e"),
			mud.NewExit("_:rhenium_road_4", "w"),
		},
	}))
}

// Central Courtyard

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
