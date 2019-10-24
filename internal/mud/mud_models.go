package mud

import "time"

type Region struct {
	UID        string `json:"uid,omitempty"`
	Type       string `json:"dgraph.type,omitempty"`
	RegionName string `json:"regionName,omitempty"`
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
	// Non-persistent fields
	ExitMap map[string]string
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
	UID         string `json:"uid,omitempty"`
	Type        string `json:"dgraph.type,omitempty"`
	TerrainType string `json:"terrainType,omitempty"`
}

type Mob struct {
	UID           string    `json:"uid,omitempty"`
	Type          string    `json:"dgraph.type,omitempty"`
	Name          string    `json:"mobName,omitempty"`
	Desc          string    `json:"mobDesc,omitempty"`
	Slug          string    `json:"mobSlug,omitempty"`
	Age           int       `json:"age,omitempty"`
	Lang          string    `json:"lang,omitempty"`
	Level         int       `json:"level,omitempty"`
	Exp           int       `json:"exp,omitempty"`
	Coins         int       `json:"coins,omitempty"`
	BankCoins     int       `json:"bankCoins,omitempty"`
	Hp            int       `json:"hp,omitempty"`
	HpMax         int       `json:"hpMax,omitempty"`
	Ap            int       `json:"ap,omitempty"`
	ApMax         int       `json:"apMax,omitempty"`
	Wimpy         int       `json:"wimpy,omitempty"`
	WimpyDir      string    `json:"wimpyDir,omitempty"`
	Encumb        int       `json:"encumb,omitempty"`
	Sober         int       `json:"sober,omitempty"`
	Thirst        int       `json:"thirst,omitempty"`
	Hunger        int       `json:"hunger,omitempty"`
	Poison        int       `json:"poison,omitempty"`
	Defend        string    `json:"defend,omitempty"`
	Aim           string    `json:"aim,omitempty"`
	Attack        string    `json:"attack,omitempty"`
	Str           int       `json:"str,omitempty"`
	Agl           int       `json:"agl,omitempty"`
	Intl          int       `json:"intl,omitempty"`
	Tgh           int       `json:"tgh,omitempty"`
	Per           int       `json:"per,omitempty"`
	StrMod        int       `json:"strMod,omitempty"`
	AglMod        int       `json:"aglMod,omitempty"`
	IntlMod       int       `json:"intlMod,omitempty"`
	TghMod        int       `json:"tghMod,omitempty"`
	PerMod        int       `json:"perMod,omitempty"`
	InsertedMobAt time.Time `json:"insertedMobAt,omitempty"`
	Items         []Item    `json:"items,omitempty"`
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
}
