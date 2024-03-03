package main

import (
	"errors"
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

func (deck *Deck[T]) DrawSingle() (T, error) {
	deckLen := len(deck.cards)
	if deckLen > 0 {
		lastIdx := deckLen - 1
		card := deck.cards[lastIdx]
		deck.cards = deck.cards[:lastIdx]
		return card, nil
	}
	return *new(T), errors.New("deck is empty")
}

func (deck *Deck[T]) Draw(numberOfCards int) ([]T, error) {
	deckLen := len(deck.cards)
	if deckLen >= numberOfCards {
		result := make([]T, numberOfCards)
		copy(result, deck.cards[deckLen-numberOfCards:])
		deck.cards = deck.cards[:deckLen-numberOfCards]
		return result, nil
	} else {
		return *new([]T), errors.New("not enough cards")
	}
}

func (deck *Deck[T]) DrawAll() []T {
	result := make([]T, len(deck.cards))
	copy(result, deck.cards)
	deck.cards = []T{}
	return result
}

func (deck *Deck[T]) AddSingle(card T) {
	deck.cards = append(deck.cards, card)
}

func (deck *Deck[T]) AddAll(cards []T) {
	deck.cards = append(deck.cards, cards...)
}
