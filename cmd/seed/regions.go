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

		{
			UID:        "_:region_forwell",
			RegionName: "Forwell",
			CoorX:      0,
			CoorY:      0,
			Type:       "Region",
			PartOf:     &mud.Region{UID: "_:region_wergild"},
		},

		{
			UID:        "_:region_inn",
			RegionName: "Inn",
			CoorX:      0,
			CoorY:      0,
			Type:       "Region",
			PartOf:     &mud.Region{UID: "_:region_forwell"},
		},
	}
}
