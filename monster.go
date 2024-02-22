package main

type Monster struct {
	runes  []Rune
	level  uint
	health int
}

func (monster Monster) castRunes() []string {
	result := make([]string, len(monster.runes))
	for idx, element := range monster.runes {
		result[idx] = element.cast()
	}
	return result
}

var werewolf = Monster{
	[]Rune{Rune{
		head:  "",
		tails: "",
	}, Rune{
		head:  "",
		tails: "",
	}, Rune{
		head:  "",
		tails: "",
	}},
	1,
	9,
}
