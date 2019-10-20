package mud

import "bytes"

// ScoreStatement represents a command for looking at a room or object.
type ScoreStatement struct {
	player *Player
}

func (s *ScoreStatement) String() string {
	var buf bytes.Buffer
	_, _ = buf.WriteString("SCORE")

	return buf.String()
}

// parseScoreStatement parses a look command and returns a Statement AST object.
// This function assumes the Score token has already been consumed.
func (p *Parser) parseScoreStatement() (*ScoreStatement, error) {
	return &ScoreStatement{}, nil
}

func (s *ScoreStatement) execute() {
	s.player.connection.Write(`Hamanu the Brave Warrior of the Black Bear (neutral)

Str: 16 (16)    Race : Dwarf (male)          Exp    : 4,314,431
Agl: 16 (16)    Class: Artificer (22)        Money  : 25,361 coins
Int: 14 (14)    Guild: The Black Bear        Banks  : 380,347 coins
Tgh: 14 (14)    Age  : 8d  14h  35m  52s     Traits : 7 points
Wis: 15 (15)    Language: Common language    You are: Unencumbered
	
Hits : 166 (166)   Defend   : None           You are: Sober
Aps  : 174 (174)   Aiming at: Body           You are: Thirsty
Wimpy: 33 hits     Attack   : Crush          You are: Hungry
Dir  : Random      Hunted by: No one         You are: Unpoisoned
`)
}

type aPlayer struct {
	Name        string
	Age         int
	CurrentRoom int
	connection  *Connection
	inventory   []*Item

	race *Race
	lang string

	level     int
	exp       int
	coins     int
	bankCoins int

	class *Class
	guild *Guild

	hp       int    // hit points - hitting zero means death
	hpMax    int    // max hit points
	ap       int    // action points. - hitting zero means you are too fatigued to do the action
	apMax    int    // max action points
	wimpy    int    // hp when it's time to run
	wimpyDir string // direction to run, if possible

	encumb int
	sober  int
	thirst int
	hunger int
	poison int

	defend string
	aim    string
	attack string

	str  int
	agl  int
	intl int
	tgh  int
	wis  int

	strMod  int
	aglMod  int
	intlMod int
	tghMod  int
	wisMod  int
}

func (s *ScoreStatement) setPlayer(player *Player) {
	s.player = player
}
