package main

import (
	"github.com/brianbroderick/wergild/internal/login"
	"github.com/brianbroderick/wergild/internal/mud"
)

func (world *worldSeed) setMobs() {
	azkul := mud.NewMob()
	azkul.UID = "_:azkul"
	azkul.Name = "Azkul"
	azkul.Slug = "azkul"
	azkul.User = login.User{
		UID:  "_:user_azkul",
		Name: "azkul",
		Pass: "password",
		Type: "User",
	}

	world.Mobs = append(world.Mobs, azkul)
}
