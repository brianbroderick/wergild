package main

import "github.com/brianbroderick/wergild/internal/mud"

func (world *worldSeed) setRooms() {
	world.inn()
}

func (world *worldSeed) inn() {
	world.Rooms = append(world.Rooms, newRoom(quickRoom{
		name:   "Beginning",
		desc:   "You are in a formless dark void where nothing surrounds you, not even air. You are the only thing in the universe. Perhaps you should imagine something.",
		region: "_:region_wergild",
		x:      0,
		y:      0,
		z:      0,
	}))
}

type quickRoom struct {
	name    string
	desc    string
	region  string
	terrain string
	x       int
	y       int
	z       int
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
		quick.region = "_:region_wergild"
	}

	if quick.terrain == "" {
		quick.terrain = "_:t_void"
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
		CoorZ:            quick.z,
		Exits:            quick.exits,
		PointsOfInterest: quick.poi,
		Smell:            quick.smell,
		Listen:           quick.listen,
		Mobs:             quick.mobs,
	}
}
