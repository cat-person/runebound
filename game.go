package main

type Game struct {
	MonsterDeck   Deck[MonsterName]
	EquipmentDeck Deck[EquipmentName]
	InjuryDeck    Deck[InjuryName]
	FitsDeck      Deck[FitName]
	OriginDeck    Deck[ProfessionName]
	AncestryDeck  Deck[AncestryName]
	Hero          Hero
}

type Deck[T any] struct {
	cards []T
}

func (deck Deck[T]) Draw() T {
	lastIdx := len(deck.cards) - 1
	card := deck.cards[lastIdx]
	deck.cards = deck.cards[:lastIdx]
	return card
}

var game = Game{
	MonsterDeck:   Deck[MonsterName]{},
	EquipmentDeck: Deck[EquipmentName]{},
	InjuryDeck:    Deck[InjuryName]{},
	FitsDeck:      Deck[FitName]{},
	OriginDeck:    Deck[ProfessionName]{},
	AncestryDeck:  Deck[AncestryName]{},
	Hero:          Hero{},
}
