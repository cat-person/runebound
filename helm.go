package main

type HelmEq struct {
	Equipment
}

var Norman Equipment = Equipment{
	Name:    "Norman",
	Type:    Helm,
	Stats:   map[StatName]int{Health: 1, Arm: 1},
	Runes:   []RuneName{},
	Actions: []ActionName{},
}
