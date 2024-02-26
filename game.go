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

var game = Game{
	MonsterDeck:   Deck[MonsterName]{},
	EquipmentDeck: Deck[EquipmentName]{},
	InjuryDeck:    Deck[InjuryName]{},
	FitsDeck:      Deck[FitName]{},
	OriginDeck:    Deck[ProfessionName]{},
	AncestryDeck:  Deck[AncestryName]{},
	Hero:          Hero{},
}
