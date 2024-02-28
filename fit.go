package main

type FitName string

// Racial
const ()

const (
	Mountain      FitName = "Mountain"      // +1 str
	Stoic         FitName = "Stoic"         // +1 wlp
	SpiritSpeaker FitName = "SpiritSpeaker" // +1 wis
	Scholar       FitName = "Scholar"       // +1 int
	SilverTongue  FitName = "SilverTongue"  // +1 cha
	IronGuts      FitName = "IronGuts"      // +1 vit

	Bully      FitName = "Bully"      // + 2 str - 1 int
	Prude      FitName = "Prude"      // + 2 wlp - 1 cha
	Elder      FitName = "Elder"      // + 2 wis - 1 vit
	Bookworm   FitName = "Bookworm"   // + 2 int - 1 str
	Sensualist FitName = "Sensualist" // + 2 cha - 1 wlp
	Bear       FitName = "Bear"       // + 2 vit - 1 wis
)

var Fits = map[FitName]Fit{
	Mountain:      Fit{map[StatName]int{Strength: 1}},
	Stoic:         Fit{map[StatName]int{Willpower: 1}},
	SpiritSpeaker: Fit{map[StatName]int{Wisdom: 1}},
	Scholar:       Fit{map[StatName]int{Intelligence: 1}},
	SilverTongue:  Fit{map[StatName]int{Charisma: 1}},
	IronGuts:      Fit{map[StatName]int{Vitality: 1}},

	Bully:      Fit{map[StatName]int{Strength: 2, Intelligence: -1}},
	Prude:      Fit{map[StatName]int{Willpower: 2, Intelligence: -1}},
	Elder:      Fit{map[StatName]int{Wisdom: 2, Intelligence: -1}},
	Bookworm:   Fit{map[StatName]int{Intelligence: 2, Strength: -1}},
	Sensualist: Fit{map[StatName]int{Charisma: 2, Willpower: -1}},
	Bear:       Fit{map[StatName]int{Vitality: 2, Wisdom: -1}},
}

type Fit struct {
	StatChange map[StatName]int
}
