package main

import "fmt"

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

			if len(arguments) > 2 && containsString(item.Keys, arguments[2]) {
				player.connection.Write(item.Description + "\n")
				return
			}
		}
		player.connection.Write("You didn't see anything unusual about " + arguments[1] + ".\n")
	} else if arguments[0] == "travel" && arguments[1] == "up" {
		player.connection.Write("You look up towards the sky.\n")
	} else if arguments[0] == "travel" && arguments[1] == "down" {
		player.connection.Write("You look down towards the ground.\n")
	} else if arguments[0] == "travel" {
		player.connection.Write(fmt.Sprintf("You look towards the %s.\n", arguments[1]))
	} else {
		player.connection.Write("Look AT or IN something, or what?\n")
	}

}
