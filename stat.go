package main

type StatName string

const (
	Strength     StatName = "Strength"     // Ability to deal more damage // ability to wield heavier weapon // ability to perform more epic physical attacks
	Vitality     StatName = "Vitality"     // Ability to take damage and recover from it // Ability to drink more potions
	Charisma     StatName = "Charisma"     // Ability to perform and resolve stuff in non lethal way
	Intelligence StatName = "Intelligence" // Ability to think quickly // Initiative // Trick usage
	Wisdom       StatName = "Wisdom"       // Ability to cast stronger spells
	Willpower    StatName = "Willpower"    // Negate non physical effects // Endure more bad stuff
)

func ZeroStats() map[StatName]int {
	return map[StatName]int{
		Strength:     0,
		Vitality:     0,
		Charisma:     0,
		Intelligence: 0,
		Wisdom:       0,
		Willpower:    0,
	}
}
