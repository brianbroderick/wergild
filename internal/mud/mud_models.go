package mud

import (
	"time"

	"github.com/brianbroderick/wergild/internal/login"
)

type Region struct {
	UID        string  `json:"uid,omitempty"`
	Type       string  `json:"dgraph.type,omitempty"`
	RegionName string  `json:"regionName,omitempty"`
	PartOf     *Region `json:"region,omitempty"`
	CoorX      int     `json:"coorX,omitempty"`
	CoorY      int     `json:"coorY,omitempty"`
	CoorZ      int     `json:"coorZ,omitempty"`
}

type Room struct {
	UID              string            `json:"uid,omitempty"`
	Type             string            `json:"dgraph.type,omitempty"`
	Region           Region            `json:"region,omitempty"`
	CoorX            int               `json:"coorX,omitempty"`
	CoorY            int               `json:"coorY,omitempty"`
	CoorZ            int               `json:"coorZ,omitempty"`
	Hash             string            `json:"locationHash,omitempty"`
	Name             string            `json:"roomName,omitempty"`
	Slug             string            `json:"roomSlug,omitempty"`
	Desc             string            `json:"roomDesc,omitempty"`
	Listen           string            `json:"roomListen,omitempty"`
	Smell            string            `json:"roomSmell,omitempty"`
	Env              []string          `json:"roomEnv,omitempty"`
	LightLevel       int               `json:"lightLevel,omitempty"`
	PointsOfInterest []PointOfInterest `json:"pointsOfInterest,omitempty"`
	Exits            []Exit            `json:"exits,omitempty"`
	Terrain          Terrain           `json:"terrain,omitempty"`
	Items            []Item            `json:"items,omitempty"`
	Mobs             []Mob             `json:"mobs,omitempty"`
	// Non-persistent fields
	ExitMap map[string]string
	MobMap  map[string]*Mob // Which mobs are currently in the room
}

// Room connects to an Exit. It's a one way exit, so there's no need for a source
type Exit struct {
	UID       string `json:"uid,omitempty"`
	Type      string `json:"dgraph.type,omitempty"`
	Dest      []Room `json:"dest,omitempty"`
	Direction string `json:"direction,omitempty"`
	Portal    string `json:"portal,omitempty"` // exit type like open, door, etc
}

type PointOfInterest struct {
	UID    string `json:"uid,omitempty"`
	Type   string `json:"dgraph.type,omitempty"`
	Name   string `json:"poiName,omitempty"`
	Desc   string `json:"poiDesc,omitempty"`
	Listen string `json:"poiListen,omitempty"`
	Smell  string `json:"poiSmell,omitempty"`
	Touch  string `json:"poiTouch,omitempty"`
	Search []Item `json:"search,omitempty"`
}

type Terrain struct {
	UID  string `json:"uid,omitempty"`
	Type string `json:"dgraph.type,omitempty"`
	Name string `json:"terrainName,omitempty"`
}

type Mob struct {
	conn          *Connection
	cmd           chan string // channel for any commands received
	me            chan string // channel to send messages from my point of view
	you           chan string // channel to send messages from your point of view
	User          login.User  `json:"user,omitempty"`
	CurrentRoom   string
	UID           string         `json:"uid,omitempty"`
	Type          string         `json:"dgraph.type,omitempty"`
	Name          string         `json:"mobName,omitempty"`
	Title         string         `json:"mobTitle,omitempty"`
	Rank          string         `json:"mobRank,omitempty"`
	Desc          string         `json:"mobDesc,omitempty"`
	Slug          string         `json:"mobSlug,omitempty"`
	Level         int            `json:"level,omitempty"`
	Exp           int            `json:"exp,omitempty"`
	Coins         int            `json:"coins,omitempty"`
	Hp            int            `json:"hp,omitempty"`
	HpMax         int            `json:"hpMax,omitempty"`
	Ap            int            `json:"ap,omitempty"`
	ApMax         int            `json:"apMax,omitempty"`
	Str           int            `json:"str,omitempty"`
	Agl           int            `json:"agl,omitempty"`
	Intl          int            `json:"intl,omitempty"`
	Tgh           int            `json:"tgh,omitempty"`
	Per           int            `json:"per,omitempty"`
	InsertedMobAt time.Time      `json:"insertedMobAt,omitempty"`
	Items         []Item         `json:"items,omitempty"`
	Conversations []Conversation `json:"conversations,omitempty"`
}

type Conversation struct {
	UID     string `json:"uid,omitempty"`
	Type    string `json:"dgraph.type,omitempty"`
	Trigger string `json:"trigger,omitempty"`
	Say     string `json:"say,omitempty"`
}

type Item struct {
	UID       string `json:"uid,omitempty"`
	Type      string `json:"dgraph.type,omitempty"`
	Name      string `json:"itemName,omitempty"`
	Desc      string `json:"itemDesc,omitempty"`
	Listen    string `json:"itemListen,omitempty"`
	Smell     string `json:"itemSmell,omitempty"`
	Touch     string `json:"itemTouch,omitempty"`
	CoinValue int    `json:"coinValue,omitempty"`
	Weight    int    `json:"weight,omitempty"`
	Items     []Item `json:"items,omitempty"`
}

type MobClass struct {
	UID  string `json:"uid,omitempty"`
	Type string `json:"dgraph.type,omitempty"`
	Name string `json:"className,omitempty"`
}

type MobRace struct {
	UID  string `json:"uid,omitempty"`
	Type string `json:"dgraph.type,omitempty"`
	Race string `json:"race,omitempty"`
}

type Guild struct {
	UID  string `json:"uid,omitempty"`
	Type string `json:"dgraph.type,omitempty"`
	Name string `json:"guildName,omitempty"`
}

type DgraphResponse struct {
	Rooms            []Room            `json:"room,omitempty"`
	PointsOfInterest []PointOfInterest `json:"pointsOfInterest,omitempty"`
	Errors           []error           `json:"errors,omitempty"`
	Mobs             []Mob             `json:"mobs,omitempty"`
}
