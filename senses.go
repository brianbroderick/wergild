package main

func look(player *Player, arguments []string) {
	currentRoom := ServerInstance.getRoom(player.CurrentRoom)

	if len(arguments) == 0 {
		currentRoom.showTo(player)
		player.connection.Write("\n")
	} else if containsString([]string{"at", "in"}, arguments[0]) {
		for _, item := range currentRoom.Inventory {
			if containsString(item.Keys, arguments[1]) {
				player.connection.Write(item.Description + "\n")
				return
			}
		}
		player.connection.Write("You didn't find anything unusual about " + arguments[1] + ".\n")
	} else {
		player.connection.Write("Look AT or IN something, or what?\n")
	}

}
