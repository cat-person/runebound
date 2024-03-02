package main

import (
	"math/rand"
)

type Deck[T any] struct {
	cards []T
}

func (deck *Deck[T]) Shuffle() {
	shuffledDeck := make([]T, len(deck.cards))

	copy(shuffledDeck, deck.cards)

	for i := range shuffledDeck {
		j := rand.Intn(i + 1)
		shuffledDeck[i], shuffledDeck[j] = shuffledDeck[j], shuffledDeck[i]
	}

	deck.cards = shuffledDeck
}

func (deck *Deck[T]) Draw() T {
	cardIdx := len(deck.cards) - 1
	card := deck.cards[cardIdx]
	deck.cards = deck.cards[:cardIdx]
	return card
}
