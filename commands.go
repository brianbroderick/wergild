package main

import (
	"errors"
	"sort"
	"strings"
)

type Command struct {
	closure                func(player *Player, arguments []string)
	executionTimeInSeconds int
}

var commandKeys []string

/**
 * Sort all command keys into a reusable lookup
 */
func prepareCommands() {
	for k, _ := range commandList {
		commandKeys = append(commandKeys, k)
	}
	sort.Strings(commandKeys)
}

/**
 * Given some user input, return the related command
 */
func getCommand(userInput string) (*Command, error) {

	for _, k := range commandKeys {
		if strings.HasPrefix(k, userInput) {
			return commandList[k], nil
		}

	}

	return nil, errors.New("I'm not sure what you mean.  Try again?")
}

/**
 * Without further ado, the command list
 */
var commandList = map[string]*Command{

	/**
	 * Look, listing room and inhabitants
	 */
	"look": {
		closure: func(player *Player, arguments []string) {
			currentRoom := ServerInstance.getRoom(player.CurrentRoom)

			if len(arguments) == 0 {
				currentRoom.showTo(player)
				player.connection.Write("\n")
			} else {
				// find something or someone in the room that matches the keyword
				// for _, item := range currentRoom.Inventory {
				// 	if containsString(item.Keys, arguments[0]) {
				// 		player.connection.Write(item.Description)
				// 		return
				// 	}
				// }
			}

		},
		executionTimeInSeconds: 0,
	},
}
