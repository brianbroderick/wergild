package main

import (
	"github.com/brianbroderick/wergild/internal/login"
	"github.com/brianbroderick/wergild/internal/mud"
)

func (world *worldSeed) setMobs() {
	erik := mud.NewMob()
	erik.UID = "_:erik"
	erik.Name = "Erik"
	erik.Slug = "Erik"
	erik.Title = "the apprentice smith"
	erik.Desc = "This tall lad already shows the stature of his uncle in sheer height. Bulging muscles have started to fill out his gangly frame, and his skin has taken on the dark hue of blacksmiths everywhere. His dark hair is clipped into a short burr, showing a faded scar near his temple. He has kind, merry brown eyes and you sense that he may have a soft spot for those new to this world."

	world.Mobs = append(world.Mobs, erik)

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
