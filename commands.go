package main

import (
	"errors"
	"sort"
	"strings"
)

type command struct {
	closure                func(player *player, arguments []string)
	executionTimeInSeconds int
}

var commandKeys []string

func prepareCommands() {
	for k := range commandList {
		commandKeys = append(commandKeys, k)
	}
	sort.Strings(commandKeys)
}

func getCommand(userInput string) (*command, error) {

	for _, k := range commandKeys {
		if strings.HasPrefix(k, userInput) {
			return commandList[k], nil
		}
	}

	return nil, errors.New("I'm not sure what you mean.  Try again?")
}

var commandList = map[string]*command{
	"look": {
		closure: func(player *player, arguments []string) {
			// currentRoom := serverInstance.getRoom(player.CurrentRoom)

			// if len(arguments) == 0 {
			// currentRoom.showTo(player)
			player.connection.write("Looking around \n")
			// } else {
			// find something or someone in the room that matches the keyword
			// for _, item := range currentRoom.Inventory {
			// 	if containsString(item.Keys, arguments[0]) {
			// 		player.connection.Write(item.Description)
			// 		return
			// 	}
			// }
			// }

		},
		executionTimeInSeconds: 0,
	},
}
