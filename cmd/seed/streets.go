package main

import "github.com/brianbroderick/wergild/internal/mud"

func (world *worldSeed) streets() {
	// Central Street
	world.Rooms = append(world.Rooms, newStreet(quickRoom{
		name: "200 S Central Street",
		x:    0,
		y:    -2,
		exits: []mud.Exit{
			mud.NewExit("_:central_courtyard_s", "n"),
			mud.NewExit("_:300_s_central_street", "s"),
		},
	}))

	world.Rooms = append(world.Rooms, newStreet(quickRoom{
		name: "300 S Central Street",
		x:    0,
		y:    -3,
		exits: []mud.Exit{
			mud.NewExit("_:200_s_central_street", "n"),
			mud.NewExit("_:400_s_central_street", "s"),
		},
	}))

	world.Rooms = append(world.Rooms, newStreet(quickRoom{
		name: "400 S Central Street",
		x:    0,
		y:    -4,
		exits: []mud.Exit{
			mud.NewExit("_:300_s_central_street", "n"),
			mud.NewExit("_:500_s_central_street", "s"),
		},
	}))

	world.Rooms = append(world.Rooms, newStreet(quickRoom{
		name: "500 S Central Street",
		x:    0,
		y:    -5,
		exits: []mud.Exit{
			mud.NewExit("_:400_s_central_street", "n"),
			mud.NewExit("_:south_well", "s"),
		},
	}))

	world.Rooms = append(world.Rooms, newStreet(quickRoom{
		name: "700 S Central Street",
		x:    0,
		y:    -7,
		exits: []mud.Exit{
			mud.NewExit("_:south_well", "n"),
			mud.NewExit("_:800_s_central_street", "s"),
		},
	}))

	world.Rooms = append(world.Rooms, newStreet(quickRoom{
		name: "800 S Central Street",
		x:    0,
		y:    -8,
		exits: []mud.Exit{
			mud.NewExit("_:700_s_central_street", "n"),
			mud.NewExit("_:900_s_central_street", "s"),
			mud.NewExit("_:hafnium_avenue_4", "w"),
		},
	}))

	world.Rooms = append(world.Rooms, newStreet(quickRoom{
		name: "900 S Central Street",
		x:    0,
		y:    -9,
		exits: []mud.Exit{
			mud.NewExit("_:800_s_central_street", "n"),
			mud.NewExit("_:iron_way_1", "e"),
			mud.NewExit("_:1000_s_central_street", "s"),
		},
	}))

	world.Rooms = append(world.Rooms, newStreet(quickRoom{
		name: "1000 S Central Street",
		x:    0,
		y:    -10,
		exits: []mud.Exit{
			mud.NewExit("_:900_s_central_street", "n"),
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
			mud.NewExit("_:500_w_scandium_street", "e"),
			mud.NewExit("_:indium_street_7", "s"),
			mud.NewExit("_:scandium_gate", "w"),
		},
	}))

	world.Rooms = append(world.Rooms, newStreet(quickRoom{
		name: "500 W Scandium Street",
		x:    -5,
		y:    1,
		exits: []mud.Exit{
			mud.NewExit("_:iridium_and_scandium", "e"),
			mud.NewExit("_:indium_and_scandium", "w"),
		},
	}))

	world.Rooms = append(world.Rooms, newStreet(quickRoom{
		name: "Iridium and Scandium",
		x:    -4,
		y:    1,
		exits: []mud.Exit{
			mud.NewExit("_:200_n_iridium_street", "n"),
			mud.NewExit("_:300_w_scandium_street", "e"),
			mud.NewExit("_:400_n_iridium_street", "s"),
			mud.NewExit("_:500_w_scandium_street", "w"),
		},
	}))

	world.Rooms = append(world.Rooms, newStreet(quickRoom{
		name: "300 W Scandium Street",
		x:    -3,
		y:    1,
		exits: []mud.Exit{
			mud.NewExit("_:scandium_well", "e"),
			mud.NewExit("_:iridium_and_scandium", "w"),
		},
	}))

	world.Rooms = append(world.Rooms, newStreet(quickRoom{
		name: "200 N Scandium Street",
		x:    -2,
		y:    2,
		exits: []mud.Exit{
			mud.NewExit("_:300_n_scandium_street", "n"),
			mud.NewExit("_:scandium_well", "s"),
		},
	}))

	world.Rooms = append(world.Rooms, newStreet(quickRoom{
		name: "300 N Scandium Street",
		x:    -2,
		y:    3,
		exits: []mud.Exit{
			mud.NewExit("_:100_w_scandium_street", "e"),
			mud.NewExit("_:200_n_scandium_street", "s"),
		},
	}))

	world.Rooms = append(world.Rooms, newStreet(quickRoom{
		name: "100 W Scandium Street",
		x:    -1,
		y:    3,
		exits: []mud.Exit{
			mud.NewExit("_:center_scandium_street", "e"),
			mud.NewExit("_:300_n_scandium_street", "w"),
		},
	}))

	world.Rooms = append(world.Rooms, newStreet(quickRoom{
		name: "Center Scandium Street",
		x:    0,
		y:    3,
		exits: []mud.Exit{
			mud.NewExit("_:300_n_ruthenium_road", "e"),
			mud.NewExit("_:100_w_scandium_street", "w"),
		},
	}))

	// Ruthenium Road

	world.Rooms = append(world.Rooms, newStreet(quickRoom{
		name: "500 N Ruthenium Road",
		x:    1,
		y:    5,
		exits: []mud.Exit{
			mud.NewExit("_:ruthenium_gate", "n"),
			mud.NewExit("_:center_rhodium_street", "w"),
			mud.NewExit("_:400_n_ruthenium_road", "s"),
		},
	}))

	world.Rooms = append(world.Rooms, newStreet(quickRoom{
		name: "400 N Ruthenium Road",
		x:    1,
		y:    4,
		exits: []mud.Exit{
			mud.NewExit("_:500_n_ruthenium_road", "n"),
			mud.NewExit("_:300_n_ruthenium_road", "s"),
		},
	}))

	world.Rooms = append(world.Rooms, newStreet(quickRoom{
		name: "300 N Ruthenium Road",
		x:    1,
		y:    3,
		exits: []mud.Exit{
			mud.NewExit("_:400_n_ruthenium_road", "n"),
			mud.NewExit("_:center_scandium_street", "w"),
			// mud.NewExit("_:continue", "e"),
		},
	}))

	// Rhodium Street

	world.Rooms = append(world.Rooms, newStreet(quickRoom{
		name: "600 W Rhodium Street",
		x:    -6,
		y:    5,
		exits: []mud.Exit{
			mud.NewExit("_:500_w_rhodium_street", "e"),
			mud.NewExit("_:500_n_indium_street", "w"),
		},
	}))

	world.Rooms = append(world.Rooms, newStreet(quickRoom{
		name: "500 W Rhodium Street",
		x:    -5,
		y:    5,
		exits: []mud.Exit{
			mud.NewExit("_:iridium_and_rhodium", "e"),
			mud.NewExit("_:600_w_rhodium_street", "w"),
		},
	}))

	world.Rooms = append(world.Rooms, newStreet(quickRoom{
		name: "300 W Rhodium Street",
		x:    -3,
		y:    5,
		exits: []mud.Exit{
			mud.NewExit("_:200_w_rhodium_street", "e"),
			mud.NewExit("_:iridium_and_rhodium", "w"),
		},
	}))

	world.Rooms = append(world.Rooms, newStreet(quickRoom{
		name: "200 W Rhodium Street",
		x:    -2,
		y:    5,
		exits: []mud.Exit{
			mud.NewExit("_:100_w_rhodium_street", "e"),
			mud.NewExit("_:300_w_rhodium_street", "w"),
		},
	}))

	world.Rooms = append(world.Rooms, newStreet(quickRoom{
		name: "100 W Rhodium Street",
		x:    -1,
		y:    5,
		exits: []mud.Exit{
			mud.NewExit("_:center_rhodium_street", "e"),
			mud.NewExit("_:200_w_rhodium_street", "w"),
		},
	}))

	world.Rooms = append(world.Rooms, newStreet(quickRoom{
		name: "Center Rhodium Street",
		x:    0,
		y:    5,
		exits: []mud.Exit{
			mud.NewExit("_:500_n_ruthenium_road", "e"),
			mud.NewExit("_:100_w_rhodium_street", "w"),
		},
	}))

	// Iridium Street

	world.Rooms = append(world.Rooms, newStreet(quickRoom{
		name: "Iridium and Rhodium",
		x:    -4,
		y:    5,
		exits: []mud.Exit{
			mud.NewExit("_:iridium_gate", "n"),
			mud.NewExit("_:300_w_rhodium_street", "e"),
			mud.NewExit("_:400_n_iridium_street", "s"),
			mud.NewExit("_:500_w_rhodium_street", "w"),
		},
	}))

	world.Rooms = append(world.Rooms, newStreet(quickRoom{
		name: "400 N Iridium Street",
		x:    -4,
		y:    4,
		exits: []mud.Exit{
			mud.NewExit("_:iridium_and_rhodium", "n"),
			mud.NewExit("_:300_n_iridium_street", "s"),
		},
	}))

	world.Rooms = append(world.Rooms, newStreet(quickRoom{
		name: "300 N Iridium Street",
		x:    -4,
		y:    3,
		exits: []mud.Exit{
			mud.NewExit("_:400_n_iridium_street", "n"),
			mud.NewExit("_:200_n_iridium_street", "s"),
		},
	}))

	world.Rooms = append(world.Rooms, newStreet(quickRoom{
		name: "200 N Iridium Street",
		x:    -4,
		y:    2,
		exits: []mud.Exit{
			mud.NewExit("_:300_n_iridium_street", "n"),
			mud.NewExit("_:iridium_and_scandium", "s"),
		},
	}))

	world.Rooms = append(world.Rooms, newStreet(quickRoom{
		name: "400 W Iridium Street",
		x:    -4,
		y:    0,
		exits: []mud.Exit{
			mud.NewExit("_:iridium_and_scandium", "n"),
			mud.NewExit("_:100_s_iridium_street", "s"),
		},
	}))

	world.Rooms = append(world.Rooms, newStreet(quickRoom{
		name: "100 S Iridium Street",
		x:    -4,
		y:    -1,
		exits: []mud.Exit{
			mud.NewExit("_:400_w_iridium_street", "n"),
			mud.NewExit("_:200_s_iridium_street", "s"),
		},
	}))

	world.Rooms = append(world.Rooms, newStreet(quickRoom{
		name: "200 S Iridium Street",
		x:    -4,
		y:    -2,
		exits: []mud.Exit{
			mud.NewExit("_:100_s_iridium_street", "n"),
			mud.NewExit("_:300_w_iridium_street", "e"),
		},
	}))

	world.Rooms = append(world.Rooms, newStreet(quickRoom{
		name: "300 W Iridium Street",
		x:    -3,
		y:    -2,
		exits: []mud.Exit{
			mud.NewExit("_:200_s_iridium_street", "w"),
			mud.NewExit("_:200_w_iridium_street", "e"),
		},
	}))

	world.Rooms = append(world.Rooms, newStreet(quickRoom{
		name: "200 W Iridium Street",
		x:    -2,
		y:    -2,
		exits: []mud.Exit{
			mud.NewExit("_:300_s_iridium_street", "s"),
			mud.NewExit("_:300_w_iridium_street", "w"),
		},
	}))

	world.Rooms = append(world.Rooms, newStreet(quickRoom{
		name: "300 S Iridium Street",
		x:    -2,
		y:    -3,
		exits: []mud.Exit{
			mud.NewExit("_:200_w_iridium_street", "n"),
			mud.NewExit("_:400_s_iridium_street", "s"),
		},
	}))

	world.Rooms = append(world.Rooms, newStreet(quickRoom{
		name: "400 S Iridium Street",
		x:    -2,
		y:    -4,
		exits: []mud.Exit{
			mud.NewExit("_:300_s_iridium_street", "n"),
			mud.NewExit("_:iridium_and_platinum", "s"),
		},
	}))

	world.Rooms = append(world.Rooms, newStreet(quickRoom{
		name: "Iridium and Platinum",
		x:    -2,
		y:    -5,
		exits: []mud.Exit{
			mud.NewExit("_:400_s_iridium_street", "n"),
			mud.NewExit("_:600_s_iridium_street", "s"),
		},
	}))

	world.Rooms = append(world.Rooms, newStreet(quickRoom{
		name: "600 S Iridium Street",
		x:    -2,
		y:    -6,
		exits: []mud.Exit{
			mud.NewExit("_:iridium_and_platinum", "n"),
			mud.NewExit("_:700_s_iridium_street", "s"),
		},
	}))

	world.Rooms = append(world.Rooms, newStreet(quickRoom{
		name: "700 S Iridium Street",
		x:    -2,
		y:    -7,
		exits: []mud.Exit{
			mud.NewExit("_:600_s_iridium_street", "n"),
			mud.NewExit("_:iridium_and_hafnium", "n"),
		},
	}))

	world.Rooms = append(world.Rooms, newStreet(quickRoom{
		name: "Iridium and Hafnium",
		x:    -2,
		y:    -8,
		exits: []mud.Exit{
			mud.NewExit("_:700_s_iridium_street", "n"),
			mud.NewExit("_:hafnium_avenue_4", "e"),
			mud.NewExit("_:hafnium_avenue_2", "w"),
		},
	}))

	// Indium Street

	world.Rooms = append(world.Rooms, newStreet(quickRoom{
		name: "500 N Indium Street",
		x:    -7,
		y:    5,
		exits: []mud.Exit{
			mud.NewExit("_:600_w_rhodium_street", "e"),
			mud.NewExit("_:indium_street_2", "s"),
		},
	}))

	world.Rooms = append(world.Rooms, newStreet(quickRoom{
		name: "Indium Street 2",
		x:    -7,
		y:    4,
		exits: []mud.Exit{
			mud.NewExit("_:500_n_indium_street", "n"),
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
			mud.NewExit("_:1000_s_central_street", "e"),
			mud.NewExit("_:rhenium_road_4", "w"),
		},
	}))

	// Hafnium Avenue

	world.Rooms = append(world.Rooms, newStreet(quickRoom{
		name: "Hafnium Avenue 1",
		x:    -4,
		y:    -8,
		exits: []mud.Exit{
			mud.NewExit("_:hafnium_avenue_2", "e"),
			mud.NewExit("_:rhenium_road_2", "s"),
		},
	}))

	world.Rooms = append(world.Rooms, newStreet(quickRoom{
		name: "Hafnium Avenue 2",
		x:    -3,
		y:    -8,
		exits: []mud.Exit{
			mud.NewExit("_:iridium_and_hafnium", "e"),
			mud.NewExit("_:hafnium_avenue_1", "w"),
		},
	}))

	world.Rooms = append(world.Rooms, newStreet(quickRoom{
		name: "Hafnium Avenue 4",
		x:    -1,
		y:    -8,
		exits: []mud.Exit{
			mud.NewExit("_:800_s_central_street", "e"),
			mud.NewExit("_:iridium_and_hafnium", "w"),
		},
	}))

	// Iron Way

	world.Rooms = append(world.Rooms, newStreet(quickRoom{
		name: "Iron Way 1",
		x:    1,
		y:    -9,
		exits: []mud.Exit{
			mud.NewExit("_:iron_way_2", "e"),
			mud.NewExit("_:900_s_central_street", "w"),
		},
	}))

	world.Rooms = append(world.Rooms, newStreet(quickRoom{
		name: "Iron Way 2",
		x:    2,
		y:    -9,
		exits: []mud.Exit{
			mud.NewExit("_:iron_way_3", "e"),
			mud.NewExit("_:iron_way_1", "w"),
		},
	}))

	world.Rooms = append(world.Rooms, newStreet(quickRoom{
		name: "Iron Way 3",
		x:    3,
		y:    -9,
		exits: []mud.Exit{
			mud.NewExit("_:ruthenium_road_16", "n"),
			mud.NewExit("_:iron_way_4", "e"),
			mud.NewExit("_:iron_way_2", "w"),
		},
	}))

	world.Rooms = append(world.Rooms, newStreet(quickRoom{
		name: "Iron Way 4",
		x:    4,
		y:    -9,
		exits: []mud.Exit{
			mud.NewExit("_:ruthenium_gate", "e"),
			mud.NewExit("_:iron_way_3", "w"),
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
			mud.NewExit("_:scandium_well", "w"),
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
			mud.NewExit("_:200_s_central_street", "s"),
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
