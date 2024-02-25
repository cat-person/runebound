package main

type MonsterName string

const (
	Werewolf MonsterName = "Werewolf"
)

type Monster struct {
	Stats   map[string]int
	Actions []ActionName
	Runes   []RuneName
}

var Monsters = map[string]Monster{
	"Werewolf": {
		Stats: map[string]int{
			"health": 8,
		},
		Actions: []ActionName{},
		Runes:   []RuneName{Copper, Copper, Red},
	},
}

// func (monster Monster) castRunes() []common.RuneSide {
// 	result := make([]common.RuneSide, len(monster.runes))
// 	for idx, element := range monster.runes {
// 		result[idx] = element.cast()
// 	}
// 	return result
// }

// func getWerewolf() Monster {
// 	return Werewolf
// }

// var Werewolf = Monster{
// 	[]Rune{{
// 		Head:  common.RuneSide{Attack, 2},
// 		Tails: common.RuneSide{Defense, 1},
// 	}, {
// 		Head:  common.RuneSide{Attack, 1},
// 		Tails: common.RuneSide{Trick, 1},
// 	}},
// 	1,
// 	9,
// }
