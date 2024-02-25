package main

type DieName string

const (
	CopperDice  DieName = "Copper"
	BronzeDice  DieName = "Bronze"
	GoldenDice  DieName = "Golden"
	RedDice     DieName = "Red"
	ScarletDice DieName = "Scarlet"
	GreenDice   DieName = "Green"
	OliveDice   DieName = "Olive"
	BlueDice    DieName = "Blue"
	IndigoDice  DieName = "Indigo"
)

type SideType string

const (
	AttackDice  SideType = "Attack"
	DefenseDice SideType = "Defense"
	MagicDice   SideType = "Magic"
	TrickDice   SideType = "Trick"
	None        SideType = "None"
)

type Side struct {
	SideType SideType
	Power    int
}

type Dice struct {
	Name  DieName
	Sides [6]Side
}
