package main

type Game struct {
	AncestryDeck   Deck[AncestryName]
	ProfessionDeck Deck[ProfessionName]
	FeatDeck       Deck[FeatName]
	Discard        struct {
		AncestryDeck   Deck[AncestryName]
		ProfessionDeck Deck[ProfessionName]
		FeatDeck       Deck[FeatName]
	}
}
