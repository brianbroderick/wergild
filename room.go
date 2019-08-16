package main

type room struct {
	id          int
	title       string
	description string
	exits       map[string]*room
	// Inventory   []*Item
}
