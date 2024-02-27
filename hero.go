package main

import (
	"encoding/json"
)

type Hero struct {
	Name       string
	Level      int
	Stats      map[string]int
	Actions    []ActionName
	Runes      []RuneName
	Fits       []Fit
	Ancestry   AncestryName
	Equipement map[EquipmentType]Equipment
}

func getHero(name string, ancestry Ancestry) Hero {
	return Hero{
		Name:       name,
		Stats:      map[string]int{"health": 5, "armor": 1},
		Actions:    []ActionName{Punch},
		Runes:      Runes{Copper, Red, Green},
		Equipement: map[EquipmentType]Equipment{},
	}
}

func (hero Hero) String() string {
	result, err := json.MarshalIndent(hero, "", "  ")
	if err == nil {
		return string(result)
	} else {
		return err.Error()
	}
}

// func (hero Hero) getActions() map[ActionName][]RuneSide {
// 	result := make(map[ActionName][]RuneSide)

// 	for _, action := range hero.Actions {
// 		result[action] = Actions[action].Price
// 	}

// 	return result
// }

func (hero *Hero) levelUp() {

}
