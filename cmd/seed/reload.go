package main

import (
	"encoding/json"
	"log"

	"github.com/brianbroderick/agora"
	"github.com/brianbroderick/wergild/internal/login"
	"github.com/brianbroderick/wergild/internal/mud"
)

func reloadData() {
	agora.DropAll()
	agora.SetSchema(schemaString())
	// loadSeed()
}

type worldSeed struct {
	Regions []mud.Region  `json:"regions,omitempty"`
	Rooms   []mud.Room    `json:"rooms,omitempty"`
	Mobs    []mud.Mob     `json:"mobs,omitempty"`
	Users   []login.User  `json:"users,omitempty"`
	Terrain []mud.Terrain `json:"terrains,omitempty"`
}

func NewWorld() worldSeed {
	return worldSeed{
		Regions: make([]mud.Region, 10),
		Rooms:   make([]mud.Room, 100),
		Mobs:    make([]mud.Mob, 10),
		Users:   make([]login.User, 10),
		Terrain: make([]mud.Terrain, 15),
	}
}

func (world *worldSeed) setValues() {
	world.setRegions()
	world.setRooms()
	world.setMobs()
	world.setTerrain()
}

func loadSeed() {
	world := NewWorld()
	world.setValues()

	j, err := json.Marshal(world)
	if err != nil {
		log.Fatal(err)
	}

	agora.MutateDgraph(j)
}
