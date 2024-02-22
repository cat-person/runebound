package main

import "math/rand"

type RuneType uint16

const (
	Attack RuneType = iota
	Defense
	Magic
	Trick
)

type Rune struct {
	head  string
	tails string
}

type RuneSide struct {
	number   int
	runeType RuneType
}

func (rune Rune) cast() string {
	if rand.Intn(2) == 0 {
		return rune.head
	} else {
		return rune.tails
	}
}
