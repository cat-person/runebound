package main

import "math/rand"

type Rune struct {
	head  string
	tails string
}

func (rune Rune) cast() string {
	if rand.Intn(2) == 0 {
		return rune.head
	} else {
		return rune.tails
	}
}
