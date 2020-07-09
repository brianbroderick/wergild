package mud

type Level struct {
	lvl  int
	name string
	xp   int
}

// These are how a string is mapped to the token
var levels = map[int]Level{
	1:  {lvl: 1, name: "the utter novice", xp: 0},
	2:  {lvl: 2, name: "the curious wanderer", xp: 676},
	3:  {lvl: 3, name: "the knowledgeable wanderer", xp: 1014},
	4:  {lvl: 4, name: "the experienced wanderer", xp: 1522},
	5:  {lvl: 5, name: "the great wanderer", xp: 2283},
	6:  {lvl: 6, name: "the simple seeker", xp: 3425},
	7:  {lvl: 7, name: "the roaming seeker ", xp: 5138},
	8:  {lvl: 8, name: "the proficient seeker", xp: 7707},
	9:  {lvl: 9, name: "the efficient seeker", xp: 11561},
	10: {lvl: 10, name: "the ambitious searcher", xp: 17341},
	11: {lvl: 11, name: "the confident searcher", xp: 26012},
	12: {lvl: 12, name: "the expert searcher", xp: 39018},
	13: {lvl: 13, name: "the veteran searcher", xp: 58527},
	14: {lvl: 14, name: "the rookie adventurer", xp: 87791},
	15: {lvl: 15, name: "the learning adventurer", xp: 131687},
	16: {lvl: 16, name: "the proficient adventurer", xp: 197530},
	17: {lvl: 17, name: "the seasoned adventurer", xp: 296296},
	18: {lvl: 18, name: "the wise adventurer ", xp: 444444},
	19: {lvl: 19, name: "the marvelous adventurer", xp: 666666},
	20: {lvl: 20, name: "the eager explorer", xp: 1000000},
}

// 	21     the insightful explorer                            2,000,000
// 	22     the well-known explorer                            3,500,000
// 	23     the professional explorer                          5,500,000
// 	24     the worldly explorer                               8,000,000
// 	25     the master explorer of Anguish                    13,000,000
// 	26     the brave quester                                 20,000,000
// 	27     the praiseworthy quester                          35,000,000
// 	28     the knowing quester                               50,000,000
// 	29     the amazing quester                               65,000,000
// 	30     the famous quester of Anguish                     80,000,000
// 	31     the natural discoverer                           100,000,000
// 	32     the talented discoverer                          140,000,000
// 	33     the masterful discoverer                         190,000,000
// 	34     the unrivaled discoverer                         250,000,000
// 	35     the legendary discoverer of Anguish              320,000,000
// 	36     the talked-about vagabond                        400,000,000
// 	37     the well-known vagabond                          500,000,000
// 	38     the popular vagabond                             600,000,000
// 	39     the famed vagabond                               700,000,000
// 	40     the historic vagabond of Anguish                 800,000,000
// 	41     the nomad of the lands of Anguish                900,000,000
// 	42     the great nomad of the lands of Anguish        1,000,000,000
// 	43     the rugged nomad of the lands of Anguish       1,100,000,000
// 	44     the weathered nomad of the lands of Anguish    1,200,000,000
// 	45     the celebrated nomad of the lands of Anguish   1,300,000,000
// 	46     the old and wise nomad of the lands of Anguish 1,450,000,000
// 	47     the sagacious nomad of the lands of Anguish    1,600,000,000
// 	48     the historic nomad of the lands of Anguish     1,750,000,000
// 	49     the ancient nomad of the lands of Anguish      1,900,000,000
// 	50     the all-knowing nomad of the lands of Anguish  2,050,000,000`
