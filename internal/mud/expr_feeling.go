package mud

import (
	"fmt"

	"github.com/brianbroderick/wergild/internal/mql"
)

// FeelingExpression represents a command for looking at a room or object.
type FeelingExpression struct {
	mob  *Mob
	stmt *mql.FeelingStatement
}

func (s *FeelingExpression) Execute() {
	me := ""
	you := ""
	switch s.stmt.Ident {
	case "ack":
		me, you = s.standardTo("ack", "acks")
	case "ah":
		me, you = s.standardTo("say 'ah'", "says 'ah'")
	case "agree":
		me, you = s.agree()
	case "apologise":
		me, you = s.apologise()
	case "applaud":
		me, you = s.applaud()
	case "beep":
		me, you = s.beep()
	case "bite":
		me, you = s.bite()
	case "bow":
		me, you = s.standardTo("bow", "bows")
	case "laugh":
		me, you = s.laugh()
	default:
		s.mob.myMessageToChannel("what?\n")
	}

	if me != "" {
		s.mob.myMessageToChannel(me)
	}

	room := WorldInstance.getRoom(s.mob.CurrentRoom)
	for _, mob := range room.MobMap {
		if s.mob.Slug == mob.Slug {
			continue
		}
		mob.yourMessageToChannel(you)
	}
}

func (s *FeelingExpression) agree() (me string, you string) {
	if s.stmt.Object == "" {
		me = "You agree wholeheartedly.\n"
		you = s.mob.Name + " agrees wholeheartedly.\n"
	} else {
		me = fmt.Sprintf("You agree wholeheartedly with %s.\n", s.stmt.Object)
		you = fmt.Sprintf("%s agrees wholeheartedly with %s.\n", s.mob.Name, s.stmt.Object)
	}
	return me, you
}

func (s *FeelingExpression) apologise() (me string, you string) {
	if s.stmt.Object == "" {
		me = "You apologise for breathing.\n"
		you = s.mob.Name + " apologises for breathing.\n"
	} else {
		me = fmt.Sprintf("You beg %s's pardon.\n", s.stmt.Object)
		you = fmt.Sprintf("%s begs %s's pardon.\n", s.mob.Name, s.stmt.Object)
	}
	return me, you
}

func (s *FeelingExpression) applaud() (me string, you string) {
	if s.stmt.Object == "" {
		me = "You applaud wholeheartedly.\n"
		you = s.mob.Name + " applauds wholeheartedly.\n"
	} else {
		me = fmt.Sprintf("You applaud %s wholeheartedly.\n", s.stmt.Object)
		you = fmt.Sprintf("%s applauds %s wholeheartedly.\n", s.mob.Name, s.stmt.Object)
	}
	return me, you
}

func (s *FeelingExpression) beep() (me string, you string) {
	if s.stmt.Object == "" {
		me = "Beep! Beep!\n"
		you = s.mob.Name + " beep beeps!\n"
	} else {
		me = fmt.Sprintf("You beep %s on the nose.\n", s.stmt.Object)
		you = fmt.Sprintf("%s beeps %s on the nose.\n", s.mob.Name, s.stmt.Object)
	}
	return me, you
}

func (s *FeelingExpression) bite() (me string, you string) {
	if s.stmt.Object == "" {
		me = "You bite your tongue! Aarghh!\n"
		you = s.mob.Name + " bites their tongue!\n"
	} else {
		me = fmt.Sprintf("You bite %s.\n", s.stmt.Object)
		you = fmt.Sprintf("%s bites %s.\n", s.mob.Name, s.stmt.Object)
	}
	return me, you
}

func (s *FeelingExpression) laugh() (me string, you string) {
	if s.stmt.Object == "" {
		me = "You fall down laughing.\n"
		you = s.mob.Name + " falls down laughing.\n"
	} else {
		me = fmt.Sprintf("You laugh at %s.\n", s.stmt.Object)
		you = fmt.Sprintf("%s laughs at %s.\n", s.mob.Name, s.stmt.Object)
	}
	return me, you
}

func (s *FeelingExpression) standardTo(mePhrase string, youPhrase string) (me string, you string) {
	if s.stmt.Object == "" {
		me = fmt.Sprintf("You %s.\n", mePhrase)
		you = fmt.Sprintf("%s %s.\n", s.mob.Name, youPhrase)
	} else {
		me = fmt.Sprintf("You %s to %s.\n", mePhrase, s.stmt.Object)
		you = fmt.Sprintf("%s %s to %s.\n", s.mob.Name, youPhrase, s.stmt.Object)
	}
	return me, you
}

// ack, agree, ah, apologise, applaud, beep, bite, bkiss, blink, blush,
// boggle, bored, bounce, bow, breathe, burp, cackle, caress, cheer, choke,
// chortle, chuckle, clap, comfort, comp, complain, confuse, confused,
// cough, cower, cringe, cry, cuddle, curious, curtsey, dance, daydream,
// despair, die, disagree, drool, duck, duh, ear, faint, fart, flash, flex,
// flip, flutter, fondle, forgive, french, frown, fume, gasp, gaze, ggrovel,
// gibber, giggle, glare, grab, grimace, grin, groan, grope, grovel, growl,
// grumble, grunt, guffaw, hand, happy, hiccup, high5, hkiss, hold, hop,
// howl, hsigh, hug, hum, ignore, insult, ising, interrupt, kick, kiss,
// knee, kneel, lag, laugh, leak, lick, love, meditate, moan, mock, mgrin,
// mumble, nibble, nod, nudge, oh, ouch, panic, pant, pat, peer, pick,
// pinch, point, poke, ponder, pout, puke, punch, purr, puzzle, raise,
// recoil, roll, ruffle, sad, scratch, scream, shake, shiver, shrug,
// shudder, sigh, sing, slap, smirk, smile, smother, snap, snarl, sneer,
// sneeze, snicker, sniff, snore, snort, snuggle, snivel, sob, spank, spit,
// squeeze, ssnarl, ssteam, stare, steam, strangle, strut, sulk, tackle,
// tap, taunt, thank, think, thug, tickle, tongue, tremble, ttackle, tthink,
// twiddle, wave, whimper, whine, whistle, wiggle, wince, wink, worship,
// wrinkle, yawn, yodel, yuck, xhappy, xlaugh, xsad, and xsob.
