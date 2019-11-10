package mud

import (
	"io/ioutil"
)

var (
	ServerInstance *Server
	WorldInstance  *World
	// BcryptCost is set lower when testing to speed up tests. It determines how much CPU
	// is used when hashing passwords
	BcryptCost = 14
)

func init() {
	ServerInstance = &Server{
		mobs:  make(map[string]*Mob),
		users: make(map[string]*Mob),
	}
	BuildWorld()
	ServerInstance.Start()
}

func motd() string {
	motdBytes, _ := ioutil.ReadFile("welcome/1.txt")
	return string(motdBytes) + "\n"
}
