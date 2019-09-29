package main

func moveTo(player *Player, arguments []string) {
	room := ServerInstance.getRoom(player.CurrentRoom)

	if val, ok := room.exits[arguments[0]]; ok {
		// fmt.Printf("%v", val.destination)
		player.CurrentRoom = val.destination
		newRoom := ServerInstance.getRoom(player.CurrentRoom)
		newRoom.showTo(player)
	}
}
