package main

import "github.com/brianbroderick/wergild/internal/mud"

func (world *worldSeed) setRegions() {
	world.Regions = []mud.Region{
		{
			UID:        "_:region_wergild",
			RegionName: "Wergild",
			CoorX:      0,
			CoorY:      0,
			Type:       "Region",
		},
	}
}
