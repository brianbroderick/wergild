package main

import "strings"

type Nick struct {
	source    string
	convertTo string
}

var globalNicks = map[string]*Nick{
	"quit": {source: "quit", convertTo: "exit"},
	"l":    {source: "l", convertTo: "look"},
	"n":    {source: "n", convertTo: "travel north"},
	"e":    {source: "e", convertTo: "travel east"},
	"w":    {source: "w", convertTo: "travel west"},
	"s":    {source: "s", convertTo: "travel south"},
	"u":    {source: "u", convertTo: "travel up"},
	"d":    {source: "d", convertTo: "travel down"},
}

func applyNick(message string) string {
	words := strings.Fields(message)

	for i, word := range words {
		if val, ok := globalNicks[word]; ok {
			words[i] = val.convertTo
		}
	}

	return strings.Join(words, " ")
}
