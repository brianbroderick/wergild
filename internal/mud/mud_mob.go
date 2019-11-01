package mud

import (
	"fmt"
	"strings"

	"github.com/brianbroderick/agora"
)

func queryMob(slug string) (Mob, error) {
	query := `query Mob($slug: string){
		mobs(first:1, func: eq(mobSlug, $slug)) {
			uid mobName mobDesc mobSlug mobTitle mobRank age lang gender 
			level exp coins bankCoins 
			hp hpMax ap apMax wimpy wimpyDir 
			encumb sober thirst hunger poison 
			defend aim attack 
			str agl intl tgh per strMod aglMod intlMod tghMod perMod
			user { userName }
		}
	}`

	variables := make(map[string]string)
	variables["$slug"] = slug

	var r DgraphResponse
	err := agora.ResolveQueryWithVars(&r, query, variables)
	if err != nil {
		return Mob{}, err
	}
	if len(r.Mobs) == 1 {
		return r.Mobs[0], nil
	}

	return Mob{}, fmt.Errorf("Mob not found.")
}

// probably not needed?
func (mob *Mob) getCurrentRoom() string {
	return mob.CurrentRoom
}

func (mob *Mob) sendPrompt() {
	if mob == nil || mob.conn == nil {
		return
	}

	str := fmt.Sprintf("%d:%d> ",
		mob.Hp,
		mob.Ap)

	mob.conn.Write(str)
}

func (mob *Mob) do(message string) {
	stmt, err := NewParser(strings.NewReader(message)).ParseStatement()
	if err != nil {
		mob.conn.Write(fmt.Sprint(err) + "\n")
		return
	}

	stmt.setMob(mob)
	stmt.execute()
}
