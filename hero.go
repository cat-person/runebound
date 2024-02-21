package main

const primary string = "primary"
const secondary string = "secondary"

type Hero struct {
	runes           []Rune
	armor           Armor
	weapon          Weapon
	secondaryWeapon Weapon
	helm            Helm
	book            Book
}

type Armor struct {
}

type Helm struct {
}

type Weapon struct {
	weaponType string
}

type Book struct {
	spells []string
}

func (hero Hero) castRunes() []string {
	result := make([]string, len(hero.runes))
	for idx, element := range hero.runes {
		result[idx] = element.cast()
	}
	return result
}
