package main

import "github.com/brianbroderick/wergild/internal/mud"

func (world *worldSeed) setTerrain() {
	world.Terrain = []mud.Terrain{
		mud.Terrain{
			UID:  "_:t_beach",
			Name: "Beach",
			Type: "Terrain",
		},
		mud.Terrain{
			UID:  "_:t_building",
			Name: "Building",
			Type: "Terrain",
		},
		mud.Terrain{
			UID:  "_:t_desert",
			Name: "Desert",
			Type: "Terrain",
		},
		mud.Terrain{
			UID:  "_:t_forest",
			Name: "Forest",
			Type: "Terrain",
		},
		mud.Terrain{
			UID:  "_:t_grassland",
			Name: "Grassland",
			Type: "Terrain",
		},
		mud.Terrain{
			UID:  "_:t_hills",
			Name: "Hills",
			Type: "Terrain",
		},
		mud.Terrain{
			UID:  "_:t_poi",
			Name: "Point of Interest",
			Type: "Terrain",
		},
		mud.Terrain{
			UID:  "_:t_settlement",
			Name: "Settlement",
			Type: "Terrain",
		},
		mud.Terrain{
			UID:  "_:t_street",
			Name: "Street",
			Type: "Terrain",
		},
		mud.Terrain{
			UID:  "_:t_swamp",
			Name: "Swamp",
			Type: "Terrain",
		},
		mud.Terrain{
			UID:  "_:t_tundra",
			Name: "Tundra",
			Type: "Terrain",
		},
		mud.Terrain{
			UID:  "_:t_underground",
			Name: "Underground",
			Type: "Terrain",
		},
		mud.Terrain{
			UID:  "_:t_water",
			Name: "Water",
			Type: "Terrain",
		},
	}
}
