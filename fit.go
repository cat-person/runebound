package main

type FeatName string

const (
	Mountain     FeatName = "Mountain"     // +1 str
	Stoic        FeatName = "Stoic"        // +1 wlp
	SpiritBound  FeatName = "SpiritBound"  // +1 wis
	Scholar      FeatName = "Scholar"      // +1 int
	SilverTongue FeatName = "SilverTongue" // +1 cha
	IronGuts     FeatName = "IronGuts"     // +1 vit

	Bully      FeatName = "Bully"      // + 2 str - 1 int
	Prude      FeatName = "Prude"      // + 2 wlp - 1 cha
	Elder      FeatName = "Elder"      // + 2 wis - 1 vit
	Bookworm   FeatName = "Bookworm"   // + 2 int - 1 str
	Sensualist FeatName = "Sensualist" // + 2 cha - 1 wlp
	Bear       FeatName = "Bear"       // + 2 vit - 1 wis
)

var Feats = map[FeatName]Feat{
	Mountain:     {map[StatName]int{Strength: 1}},
	Stoic:        {map[StatName]int{Willpower: 1}},
	SpiritBound:  {map[StatName]int{Wisdom: 1}},
	Scholar:      {map[StatName]int{Intelligence: 1}},
	SilverTongue: {map[StatName]int{Charisma: 1}},
	IronGuts:     {map[StatName]int{Vitality: 1}},

	Bully:      {map[StatName]int{Strength: 2, Intelligence: -1}},
	Prude:      {map[StatName]int{Willpower: 2, Intelligence: -1}},
	Elder:      {map[StatName]int{Wisdom: 2, Intelligence: -1}},
	Bookworm:   {map[StatName]int{Intelligence: 2, Strength: -1}},
	Sensualist: {map[StatName]int{Charisma: 2, Willpower: -1}},
	Bear:       {map[StatName]int{Vitality: 2, Wisdom: -1}},
}

type Feat struct {
	StatChange map[StatName]int
}
