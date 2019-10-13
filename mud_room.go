package main

// type Room struct {
// 	id          int
// 	title       string
// 	description string
// 	exits       map[string]*Exit
// 	Inventory   []*Item
// }

// type Exit struct {
// 	source      int
// 	destination int
// }

/**
 * Load rooms from database
 * @TODO Actually load them from database...
 */
func loadRooms() map[int]*Room {
	rooms := make(map[int]*Room, 5)
	rooms[1] = &Room{
		UID:  "1",
		Name: "Ancient Inn common room",
		Desc: "You stand in the common room of the Ancient Inn of Tantallon. There are a number of chairs and tables scattered around the room, and there are two booths where people can go for private conversation. There is a large desk at the north end of the room, over which hangs an ornate clock. A doorway leads south into the world of Ancient Anguish and the adventure it has to offer.",
		Items: []Item{
			{
				Name: "Oak Chairs",
				Desc: "The chairs are sturdy, and made from oak. They complement the other furnishings nicely.",
			},
		},
		// Exits: map[string]*Exit{"NORTH": &Exit{1, 2}},
	}

	rooms[2] = &Room{
		UID:  "2",
		Name: "Ancient Inn north room",
		Desc: "You are at a small room north of the inn's common room. A large firepit is the dominating feature here, casting warmth and powerful shadows across the tables and chairs arranged around the room. A large window to the northwest displays the forest outside.",
		Items: []Item{
			{
				Name: "a warm firepit",
				Desc: "The firepit is set halfway into the northern wall, spreading warmth throughout the inn.",
			},
		},
		// Exits: map[string]*Exit{"SOUTH": &Exit{2, 1}},
	}

	return rooms
}

func (room *Room) showTo(player *Player) {
	//room.title + "\n\n" +
	str := room.Desc + room.listExits() + room.listContents()
	player.connection.Write(str)
}

func (room *Room) listContents() string {
	str := ""
	for _, item := range room.Items {
		str += item.Desc + "\n"
	}
	return str
}

func (room *Room) listExits() string {
	return "listExits"
	// keys := make([]string, len(room.Exits))

	// i := 0
	// for k := range room.Exits {
	// 	keys[i] = k
	// 	i++
	// }
	// return "\n Obvious directions are:\n  " + strings.Join(keys, ", ") + "\n"
}
