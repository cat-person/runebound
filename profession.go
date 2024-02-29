package main

type ProfessionName string

// Common
const (
	Blacksmith ProfessionName = "Blacksmith"
	Herbalist  ProfessionName = "Herbalist"
	Fighter    ProfessionName = "Fighter"
	Criminal   ProfessionName = "Criminal"
)

type Profession struct {
	Description string
	StatChange  map[StatName]int
}

var Professions map[ProfessionName]Profession = map[ProfessionName]Profession{
	Blacksmith: {"Strong and independent sword and shield maker", map[StatName]int{Strength: 1, Willpower: 1, Vitality: 1}},
	Herbalist:  {"Herbalist knows a thing or two about grass and berries", map[StatName]int{Wisdom: 2, Vitality: 1}},
	Fighter:    {"There is nobody better in skullcrushing routine than good old fighter", map[StatName]int{Strength: 2, Vitality: 1}},
	Criminal:   {"You are scum you are horrible being but if that's come with a big paycheck you are totally ok with it", map[StatName]int{Intelligence: 2, Charisma: 1}},
}
