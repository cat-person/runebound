package main

type AncestryName string

// Elf Progression
const (
	Elf    AncestryName = "Elf"
	Faerun AncestryName = "Faerun" // Wild elfs
	Sithi  AncestryName = "Sithi"  // Magic
	Drow   AncestryName = "Drow"   // Dark edgy stuff
)

// Dwarf Progression
const (
	Dwarf       AncestryName = "Dwarf"
	GoldenDwarf AncestryName = "GoldenDwarf" // More rewards / More life
	ShieldDwarf AncestryName = "ShieldDwarf" // Stronger and better protected dwarfs
	Duergar     AncestryName = "Duergar"     // Sneaky evil dwarfs
)

// Human progression
const (
	Human     AncestryName = "Human"
	Calishite AncestryName = "Calishite" // Diplomacy + More loot
	Mulan     AncestryName = "Mulan"     // Fighty humans
	Rashemi   AncestryName = "Rashemi"   // Magical humans
)

// Dragonborn
const (
	Dragonkin       AncestryName = "Dragonkin"
	BronzeDragonkin AncestryName = "BronzeDragonkin"
	BlackDragonkin  AncestryName = "BlackDragonkin"
	RedDragonkin    AncestryName = "RedDragonkin"
)

var ancestries map[AncestryName]Ancestry = map[AncestryName]Ancestry{
	Elf:    {"Pointy ears, blunt personality", map[StatName]int{Willpower: 2, Wisdom: 3, Intelligence: 1, Charisma: 2}, []AncestryName{Faerun, Sithi, Drow}},
	Faerun: {"Wild Elf", map[StatName]int{Strength: 2, Vitality: 2}, []AncestryName{}},
	Sithi:  {"Magic Elf", map[StatName]int{Willpower: 1, Wisdom: 3}, []AncestryName{}},
	Drow:   {"Dark, edgy Elf", map[StatName]int{Strength: 2, Willpower: 1, Vitality: 1}, []AncestryName{}},

	Dwarf:       {"5'11 Alchoholic", map[StatName]int{Strength: 3, Willpower: 2, Vitality: 3}, []AncestryName{GoldenDwarf, ShieldDwarf, Duergar}},
	GoldenDwarf: {"Avatar of greed: Will get More rewards / More life", map[StatName]int{Intelligence: 2, Charisma: 2}, []AncestryName{}},
	ShieldDwarf: {"Stronger and better protected dwarfs", map[StatName]int{Willpower: 2, Vitality: 2}, []AncestryName{}},
	Duergar:     {"Sneaky evil dwarfs", map[StatName]int{Strength: 2, Willpower: 1, Wisdom: 1}, []AncestryName{}},

	Human:     {"6'0 Alchoholic", map[StatName]int{Strength: 2, Willpower: 1, Wisdom: 1, Intelligence: 1, Charisma: 1, Vitality: 2}, []AncestryName{Calishite, Mulan, Rashemi}},
	Calishite: {"Theatrical and narcissistic this humans prefer talking to pretty much anything: Diplomacy + More loot", map[StatName]int{Intelligence: 1, Charisma: 3}, []AncestryName{}},
	Mulan:     {"Brutal warriors", map[StatName]int{Strength: 2, Willpower: 1, Vitality: 1}, []AncestryName{}},
	Rashemi:   {"Magic aligned humans", map[StatName]int{Willpower: 1, Wisdom: 3}, []AncestryName{}},

	Dragonkin:       {"Dragons, you like dragons right?", map[StatName]int{Strength: 3, Willpower: 3, Vitality: 2}, []AncestryName{BronzeDragonkin, BlackDragonkin, RedDragonkin}},
	BronzeDragonkin: {"Dragonkin with more defence / atack", map[StatName]int{Willpower: 2, Vitality: 2}, []AncestryName{}},
	BlackDragonkin:  {"Evil sneaky magical Dragonkin", map[StatName]int{Strength: 2, Intelligence: 2}, []AncestryName{}},
	RedDragonkin:    {"Evil fighty magical Dragonkin", map[StatName]int{Strength: 1, Wisdom: 2, Charisma: 1}, []AncestryName{}},
}

type Ancestry struct {
	Description string
	StatChange  map[StatName]int
	Progression []AncestryName
}

func getAncestryByCode(ancestryName AncestryName) Ancestry {
	value, keyExist := ancestries[ancestryName]

	if keyExist {
		return value
	}
	return Ancestry{
		"Not really an ancestry",
		map[StatName]int{},
		[]AncestryName{},
	}
}
