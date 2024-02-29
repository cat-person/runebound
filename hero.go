package main

import (
	"encoding/json"
)

type Hero struct {
	Name       string
	Level      int
	Fits       []FitName
	Ancestry   []AncestryName
	Profession []ProfessionName
}

func (hero *Hero) getStats() map[StatName]int {
	result := ZeroStats()
	for _, ancestryName := range hero.Ancestry {
		ancestry := Ancestries[ancestryName]
		for statName, statValue := range ancestry.StatChange {
			result[statName] += statValue
		}
	}

	for _, professionName := range hero.Profession {
		profession := Professions[professionName]
		for statName, statValue := range profession.StatChange {
			result[statName] += statValue
		}
	}

	for _, fitName := range hero.Fits {
		fit := Fits[fitName]
		for statName, statValue := range fit.StatChange {
			result[statName] += statValue
		}
	}

	return result
}

func (hero Hero) String() string {
	result, err := json.MarshalIndent(hero, "", "  ")
	if err == nil {
		return string(result)
	} else {
		return err.Error()
	}
}
