package main

import "github.com/brianbroderick/wergild/internal/mud"

func (world *worldSeed) setRegions() {
	world.Regions = []mud.Region{
		mud.Region{
			UID:        "_:region_wergild",
			RegionName: "Wergild",
			CoorX:      0,
			CoorY:      0,
			Type:       "Region",
		},

		mud.Region{
			UID:        "_:region_forwell",
			RegionName: "Forwell",
			CoorX:      0,
			CoorY:      0,
			Type:       "Region",
			PartOf:     &mud.Region{UID: "_:region_wergild"},
		},

		mud.Region{
			UID:        "_:region_nw_tower",
			RegionName: "NW Tower",
			CoorX:      -8,
			CoorY:      6,
			Type:       "Region",
			PartOf:     &mud.Region{UID: "_:region_forwell"},
		},

		mud.Region{
			UID:        "_:region_iridium_gate",
			RegionName: "Iridium Gate",
			CoorX:      -4,
			CoorY:      6,
			Type:       "Region",
			PartOf:     &mud.Region{UID: "_:region_forwell"},
		},

		mud.Region{
			UID:        "_:region_rhodium_gate",
			RegionName: "Rhodium Gate",
			CoorX:      1,
			CoorY:      6,
			Type:       "Region",
			PartOf:     &mud.Region{UID: "_:region_forwell"},
		},

		mud.Region{
			UID:        "_:region_ne_tower",
			RegionName: "NE Tower",
			CoorX:      6,
			CoorY:      6,
			Type:       "Region",
			PartOf:     &mud.Region{UID: "_:region_forwell"},
		},

		mud.Region{
			UID:        "_:region_warehouse",
			RegionName: "Warehouse",
			CoorX:      -5,
			CoorY:      4,
			Type:       "Region",
			PartOf:     &mud.Region{UID: "_:region_forwell"},
		},

		mud.Region{
			UID:        "_:region_boneyard",
			RegionName: "Boneyard",
			CoorX:      5,
			CoorY:      4,
			Type:       "Region",
			PartOf:     &mud.Region{UID: "_:region_forwell"},
		},

		mud.Region{
			UID:        "_:region_blacksmith",
			RegionName: "Blacksmith",
			CoorX:      -1,
			CoorY:      2,
			Type:       "Region",
			PartOf:     &mud.Region{UID: "_:region_forwell"},
		},

		mud.Region{
			UID:        "_:region_inn",
			RegionName: "Inn",
			CoorX:      0,
			CoorY:      2,
			Type:       "Region",
			PartOf:     &mud.Region{UID: "_:region_forwell"},
		},

		mud.Region{
			UID:        "_:region_map_shop",
			RegionName: "Map shop",
			CoorX:      1,
			CoorY:      2,
			Type:       "Region",
			PartOf:     &mud.Region{UID: "_:region_forwell"},
		},

		mud.Region{
			UID:        "_:region_monestary",
			RegionName: "Monestary",
			CoorX:      4,
			CoorY:      2,
			Type:       "Region",
			PartOf:     &mud.Region{UID: "_:region_forwell"},
		},

		mud.Region{
			UID:        "_:region_scandium_gate",
			RegionName: "Scandium Gate",
			CoorX:      -7,
			CoorY:      1,
			Type:       "Region",
			PartOf:     &mud.Region{UID: "_:region_forwell"},
		},

		mud.Region{
			UID:        "_:region_trading_post",
			RegionName: "Trading Post",
			CoorX:      2,
			CoorY:      1,
			Type:       "Region",
			PartOf:     &mud.Region{UID: "_:region_forwell"},
		},

		mud.Region{
			UID:        "_:region_magicians_tower",
			RegionName: "Magicians Tower",
			CoorX:      -2,
			CoorY:      -1,
			Type:       "Region",
			PartOf:     &mud.Region{UID: "_:region_forwell"},
		},

		mud.Region{
			UID:        "_:region_magic_shop",
			RegionName: "Magic Shop",
			CoorX:      2,
			CoorY:      0,
			Type:       "Region",
			PartOf:     &mud.Region{UID: "_:region_forwell"},
		},

		mud.Region{
			UID:        "_:region_palladium_gate",
			RegionName: "Palladium Gate",
			CoorX:      6,
			CoorY:      0,
			Type:       "Region",
			PartOf:     &mud.Region{UID: "_:region_forwell"},
		},

		mud.Region{
			UID:        "_:region_general_store",
			RegionName: "General Store",
			CoorX:      2,
			CoorY:      -1,
			Type:       "Region",
			PartOf:     &mud.Region{UID: "_:region_forwell"},
		},

		mud.Region{
			UID:        "_:region_library",
			RegionName: "Library",
			CoorX:      -5,
			CoorY:      -2,
			Type:       "Region",
			PartOf:     &mud.Region{UID: "_:region_forwell"},
		},

		mud.Region{
			UID:        "_:region_weapon_store",
			RegionName: "Weapon Store",
			CoorX:      -1,
			CoorY:      -2,
			Type:       "Region",
			PartOf:     &mud.Region{UID: "_:region_forwell"},
		},

		mud.Region{
			UID:        "_:region_armor_store",
			RegionName: "Armor Store",
			CoorX:      -1,
			CoorY:      -3,
			Type:       "Region",
			PartOf:     &mud.Region{UID: "_:region_forwell"},
		},

		mud.Region{
			UID:        "_:region_bank",
			RegionName: "Bank",
			CoorX:      1,
			CoorY:      -3,
			Type:       "Region",
			PartOf:     &mud.Region{UID: "_:region_forwell"},
		},

		mud.Region{
			UID:        "_:region_iridium_gate",
			RegionName: "Iridium Gate",
			CoorX:      -7,
			CoorY:      -4,
			Type:       "Region",
			PartOf:     &mud.Region{UID: "_:region_forwell"},
		},

		mud.Region{
			UID:        "_:region_auction",
			RegionName: "Auction House",
			CoorX:      -1,
			CoorY:      -4,
			Type:       "Region",
			PartOf:     &mud.Region{UID: "_:region_forwell"},
		},

		mud.Region{
			UID:        "_:region_post_office",
			RegionName: "Post Office",
			CoorX:      2,
			CoorY:      -5,
			Type:       "Region",
			PartOf:     &mud.Region{UID: "_:region_forwell"},
		},

		mud.Region{
			UID:        "_:region_tungsten_gate",
			RegionName: "Tungsten Gate",
			CoorX:      6,
			CoorY:      -6,
			Type:       "Region",
			PartOf:     &mud.Region{UID: "_:region_forwell"},
		},

		mud.Region{
			UID:        "_:region_lord_viktor_mansion",
			RegionName: "Lord Viktor Mansion",
			CoorX:      1,
			CoorY:      -7,
			Type:       "Region",
			PartOf:     &mud.Region{UID: "_:region_forwell"},
		},

		mud.Region{
			UID:        "_:region_bakery",
			RegionName: "Bakery",
			CoorX:      -3,
			CoorY:      -9,
			Type:       "Region",
			PartOf:     &mud.Region{UID: "_:region_forwell"},
		},

		mud.Region{
			UID:        "_:region_cx",
			RegionName: "Commodities Exchange",
			CoorX:      -2,
			CoorY:      -9,
			Type:       "Region",
			PartOf:     &mud.Region{UID: "_:region_forwell"},
		},

		mud.Region{
			UID:        "_:region_ruthenium_gate",
			RegionName: "Ruthenium Gate",
			CoorX:      5,
			CoorY:      -9,
			Type:       "Region",
			PartOf:     &mud.Region{UID: "_:region_forwell"},
		},

		mud.Region{
			UID:        "_:region_warrior_hall",
			RegionName: "Warrior Hall",
			CoorX:      1,
			CoorY:      -10,
			Type:       "Region",
			PartOf:     &mud.Region{UID: "_:region_forwell"},
		},

		mud.Region{
			UID:        "_:region_indium_gate",
			RegionName: "Indium Gate",
			CoorX:      -6,
			CoorY:      -11,
			Type:       "Region",
			PartOf:     &mud.Region{UID: "_:region_forwell"},
		},

		mud.Region{
			UID:        "_:region_rhenium_gate",
			RegionName: "Rhenium Gate",
			CoorX:      0,
			CoorY:      -11,
			Type:       "Region",
			PartOf:     &mud.Region{UID: "_:region_forwell"},
		},
	}
}
