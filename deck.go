package main

import "math/rand"

type Deck[T any] struct {
	cards []T
}

func (deck *Deck[T]) Draw() T {
	lastIdx := len(deck.cards) - 1
	card := deck.cards[lastIdx]
	deck.cards = deck.cards[:lastIdx]
	return card
}

func (deck *Deck[T]) Shuffle() {
	shuffledCards := make([]T, len(deck.cards))

	copy(shuffledCards, deck.cards)

	for i := range shuffledCards {
		j := rand.Intn(i + 1)
		shuffledCards[i], shuffledCards[j] = shuffledCards[j], shuffledCards[i]
	}

	deck.cards = shuffledCards
}
