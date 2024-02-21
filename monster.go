package main

type Monster struct {
	runes           []Rune
	health          int
	attack          int
	secondaryAttack int
	abilities       []Ability
}

func (monster Monster) castRunes() []string {
	result := make([]string, len(monster.runes))
	for idx, element := range monster.runes {
		result[idx] = element.cast()
	}
	return result
}
