package main

import (
	"errors"
	"math/rand"
)

type Runes []RuneName

type RuneName string

const (
	Copper  RuneName = "Copper"
	Bronze  RuneName = "Bronze"
	Golden  RuneName = "Golden"
	Red     RuneName = "Red"
	Scarlet RuneName = "Scarlet"
	Green   RuneName = "Green"
	Olive   RuneName = "Olive"
	Blue    RuneName = "Blue"
	Indigo  RuneName = "Indigo"
)

type RuneType uint16

const (
	Attack RuneType = iota
	Defense
	Magic
	Trick
)

type Rune struct {
	Head  RuneSide
	Tails RuneSide
}

type RuneSide struct {
	RuneType RuneType
	Power    int
}

func (rune Rune) cast() RuneSide {
	if rand.Intn(2) == 0 {
		return rune.Head
	} else {
		return rune.Tails
	}
}

func (runeType RuneType) MarshalJSON() ([]byte, error) {
	switch runeType {
	case Attack:
		return []byte("\"Attack\""), nil
	case Defense:
		return []byte("\"Defense\""), nil
	case Magic:
		return []byte("\"Magic\""), nil
	case Trick:
		return []byte("\"Trick\""), nil
	default:
		return nil, errors.New("unknown rune type")
	}
}

var NameToRune = map[RuneName]Rune{
	Copper:  {defense(1), attack(1)},
	Bronze:  {defense(2), attack(1)},
	Golden:  {defense(2), attack(2)},
	Red:     {attack(1), magic(1)},
	Scarlet: {attack(2), magic(1)},
	Green:   {attack(1), trick(1)},
	Olive:   {attack(2), trick(1)},
	Blue:    {magic(1), trick(1)},
	Indigo:  {magic(2), trick(1)},
}

func attack(power int) RuneSide {
	return RuneSide{Attack, power}
}

func defense(power int) RuneSide {
	return RuneSide{Defense, power}
}

func magic(power int) RuneSide {
	return RuneSide{Magic, power}
}

func trick(power int) RuneSide {
	return RuneSide{Trick, power}
}

func Cast(runes []RuneName) []RuneSide {
	result := make([]RuneSide, len(runes))
	for idx, element := range runes {
		result[idx] = NameToRune[element].cast()
	}
	return result
}
