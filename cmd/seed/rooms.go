package main

import "github.com/brianbroderick/wergild/internal/mud"

func (world *worldSeed) setRooms() {
	world.inn()
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
		},
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
