package main

import (
	"errors"
	"sort"
	"strings"
)

type Command struct {
	closure                func(player *Player, verb string, arguments []string)
	verb                   string
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

	return nil, errors.New("what?")
}

/**
 * Without further ado, the command list
 * TODO: Figure out good way to "nick" or alias commands
 */
var commandList = map[string]*Command{
	"exit": {
		closure: func(player *Player, verb string, arguments []string) {
			player.connection.Write("You slowly fade away.\n")
			player.connection.conn.Close()
			ServerInstance.playerExited(player.connection, nil)
			ServerInstance.onClientConnectionClosed(player.connection, nil)
		},
		executionTimeInSeconds: 0,
	},
	// Senses:
	// Look, listing room and inhabitants
	"look": {
		closure: func(player *Player, verb string, arguments []string) {
			look(player, arguments)
		},
		executionTimeInSeconds: 0,
	},
	// Travel Commands
	"travel": {
		closure: func(player *Player, verb string, arguments []string) {
			moveTo(player, arguments)
		},
		executionTimeInSeconds: 0,
	},
	// "north": {
	// 	closure: func(player *Player, verb string, arguments []string) {
	// 		moveTo(player, verb, arguments)
	// 	},
	// 	executionTimeInSeconds: 0,
	// },
	// "south": {
	// 	closure: func(player *Player, verb string, arguments []string) {
	// 		moveTo(player, verb, arguments)
	// 	},
	// 	executionTimeInSeconds: 0,
	// },
	// "east": {
	// 	closure: func(player *Player, verb string, arguments []string) {
	// 		moveTo(player, verb, arguments)
	// 	},
	// 	executionTimeInSeconds: 0,
	// },
	// "west": {
	// 	closure: func(player *Player, verb string, arguments []string) {
	// 		moveTo(player, verb, arguments)
	// 	},
	// 	executionTimeInSeconds: 0,
	// },
}
