package main

type HelmEq struct {
	Equipment
}

var Norman Equipment = Equipment{
	Name:    "Norman",
	Type:    Helm,
	Stats:   map[StatName]int{},
	Runes:   []RuneName{},
	Actions: []ActionName{},
}
