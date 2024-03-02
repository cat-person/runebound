package main

type Game struct {
	AncestryDeck   Deck[AncestryName]
	ProfessionDeck Deck[ProfessionName]
	FitDeck        Deck[FitName]
}
