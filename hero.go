package main

import (
	"encoding/json"
)

type Hero struct {
	Name       string
	Stats      map[string]int
	Actions    []ActionName
	Runes      []RuneName
	Equipement map[EquipmentType]Equipment
}

func getHero(name string) Hero {
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
