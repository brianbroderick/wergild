package main

import "github.com/brianbroderick/wergild/internal/mud"

func (world *worldSeed) setTerrain() {
	world.Terrain = []mud.Terrain{
		{
			UID:  "_:t_beach",
			Name: "Beach",
			Type: "Terrain",
		},
		{
			UID:  "_:t_building",
			Name: "Building",
			Type: "Terrain",
		},
		{
			UID:  "_:t_desert",
			Name: "Desert",
			Type: "Terrain",
		},
		{
			UID:  "_:t_forest",
			Name: "Forest",
			Type: "Terrain",
		},
		{
			UID:  "_:t_grassland",
			Name: "Grassland",
			Type: "Terrain",
		},
		{
			UID:  "_:t_hills",
			Name: "Hills",
			Type: "Terrain",
		},
		{
			UID:  "_:t_jungle",
			Name: "Jungle",
			Type: "Terrain",
		},
		{
			UID:  "_:t_poi",
			Name: "Point of Interest",
			Type: "Terrain",
		},
		// village, town, city
		{
			UID:  "_:t_settlement",
			Name: "Settlement",
			Type: "Terrain",
		},
		// A shop is a specialized building. Shop should override a building
		{
			UID:  "_:t_shop",
			Name: "Shop",
			Type: "Terrain",
		},
		{
			UID:  "_:t_street",
			Name: "Street",
			Type: "Terrain",
		},
		{
			UID:  "_:t_swamp",
			Name: "Swamp",
			Type: "Terrain",
		},
		{
			UID:  "_:t_tundra",
			Name: "Tundra",
			Type: "Terrain",
		},
		{
			UID:  "_:t_underground",
			Name: "Underground",
			Type: "Terrain",
		},
		{
			UID:  "_:t_water",
			Name: "Water",
			Type: "Terrain",
		},
	}
}
