package mud

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/brianbroderick/agora"
	"github.com/brianbroderick/wergild/internal/login"
)

func queryMob(slug string) (Mob, error) {
	query := `query Mob($slug: string){
		mobs(first:1, func: eq(mobSlug, $slug)) {
			uid mobName mobDesc mobSlug mobTitle mobRank 
			level exp coins 
			hp hpMax ap apMax 
			str agl intl tgh per 
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

	mob.myMessageToChannel(str)
}

func (mob *Mob) pulseUpdate() {
	currentRoom := WorldInstance.getRoom(mob.CurrentRoom)
	currentRoom.showEnv(mob)
}

func (mob *Mob) myChannelToConsole() {
	select {
	case msg := <-mob.me:
		if mob.conn != nil {
			mob.conn.Write(msg)
		}
	default:
		mob.messageErr()
	}
}

func (mob *Mob) myMessageToChannel(str string) {
	select {
	case mob.me <- str:
		mob.myChannelToConsole()
	default:
		mob.messageErr()
	}
}

func (mob *Mob) yourChannelToConsole() {
	select {
	case msg := <-mob.you:
		if mob.conn != nil {
			mob.conn.Write(msg)
		}
	default:
		mob.messageErr()
	}
}

func (mob *Mob) yourMessageToChannel(str string) {
	select {
	case mob.you <- str:
		mob.yourChannelToConsole()
	default:
		mob.messageErr()
	}
}

func (mob *Mob) messageErr() {
	mob.conn.Write("Slow down. Your last message was not processed.\n")
}

func CreateUserMob(user login.User) error {
	m := NewMob()
	m.Name = user.Name
	m.Slug = strings.ToLower(user.Name)

	j, err := json.Marshal(m)
	if err != nil {
		return err
	}

	agora.MutateDgraph(j)
	return nil
}

func NewMob() Mob {
	return Mob{
		UID:   "_:mob",
		Type:  "Mob",
		Title: "the utter novice",
		Level: 1,
		Hp:    50,
		HpMax: 50,
		Ap:    50,
		ApMax: 50,
		Str:   1,
		Agl:   1,
		Intl:  1,
		Tgh:   1,
		Per:   1,
	}
}
