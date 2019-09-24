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

	return nil, errors.New("what?")
}

/**
 * Without further ado, the command list
 */
var commandList = map[string]*Command{

	// Look, listing room and inhabitants
	"l": {
		closure: func(player *Player, arguments []string) {
			look(player, arguments)
		},
		executionTimeInSeconds: 0,
	},
	"look": {
		closure: func(player *Player, arguments []string) {
			look(player, arguments)
		},
		executionTimeInSeconds: 0,
	},
	"exit": {
		closure: func(player *Player, arguments []string) {
			player.connection.Write("You slowly fade away.\n")
			player.connection.conn.Close()
			ServerInstance.playerExited(player.connection, nil)
			ServerInstance.onClientConnectionClosed(player.connection, nil)
		},
		executionTimeInSeconds: 0,
	},
}
