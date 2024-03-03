package main

import (
	"fmt"
	"strconv"
)

func main() {
	game := Game{
		AncestryDeck:   Deck[AncestryName]{[]AncestryName{Elf, Dwarf, Human, Dragonkin}},
		ProfessionDeck: Deck[ProfessionName]{[]ProfessionName{Blacksmith, Herbalist, Fighter, Criminal}},
		FitDeck:        Deck[FitName]{[]FitName{Mountain, Stoic, SpiritBound, Scholar, SilverTongue, IronGuts, Bully, Prude, Elder, Bookworm, Sensualist, Bear}},
		Discard: struct {
			AncestryDeck   Deck[AncestryName]
			ProfessionDeck Deck[ProfessionName]
			FitDeck        Deck[FitName]
		}{
			AncestryDeck:   Deck[AncestryName]{[]AncestryName{}},
			ProfessionDeck: Deck[ProfessionName]{[]ProfessionName{}},
			FitDeck:        Deck[FitName]{[]FitName{}},
		},
	}

	game.AncestryDeck.Shuffle()
	game.ProfessionDeck.Shuffle()
	game.FitDeck.Shuffle()

	hero := Hero{}

	heroName := ChooseHeroName()
	fmt.Printf("Your name is %s\n", heroName)
	hero.Name = heroName

	ancestryChoices, error := game.AncestryDeck.Draw(2)
	if error == nil {
		ancestryName, discardedCards := ChooseAncestry(ancestryChoices, ancestries)
		game.Discard.AncestryDeck.AddAll(discardedCards)
		hero.Ancestry = append(hero.Ancestry, ancestryName)
	} else {
		panic(error.Error())
	}

	professionChoices, error := game.ProfessionDeck.Draw(2)
	if error == nil {
		professionName, discardedCards := ChooseProfession(professionChoices, professions)
		game.Discard.ProfessionDeck.AddAll(discardedCards)
		hero.Profession = append(hero.Profession, professionName)
	} else {
		panic(error.Error())
	}

	fmt.Println()
	fmt.Printf("Congratulations you are %s %s now!\n", hero.Ancestry[0], hero.Profession[0])

	for hero.Level < 10 {
		fmt.Printf("Your stats are: %v\n", hero.getStats())
		fmt.Println()
		fmt.Println("Press Enter to level up")
		fmt.Scanln()

		hero.levelUp(&game)
	}
	fmt.Printf("%v\n", hero)
	fmt.Printf("%v\n", game)
}

func ChooseHeroName() string {
	fmt.Println("What is your name?")
	var name string
	fmt.Scan(&name)
	return name
}

func ChooseAncestry(choices []AncestryName, ancestries map[AncestryName]Ancestry) (AncestryName, []AncestryName) {
	fmt.Println("What is your ancestry?")

	for ancestryIdx, ancestryName := range choices {
		ancestry := ancestries[ancestryName]
		fmt.Printf("%d. %s: %s\n", ancestryIdx+1, ancestryName, ancestry.Description)
	}

	fmt.Println("Choose one of the above")
	var result AncestryName
	var resultIdx int

ancestryChoice:
	for {
		var command string
		fmt.Scan(&command)
		if ancestryIdx, err := strconv.Atoi(command); err == nil {
			if 0 < ancestryIdx && ancestryIdx <= len(choices) {
				resultIdx = ancestryIdx - 1
				result = choices[resultIdx]
				break ancestryChoice
			} else {
				fmt.Printf("Please choose option by number between 1 and %d\n", len(choices))
			}
		} else {
			for ancestryIdx, ancestryName := range choices {
				if command == fmt.Sprint(ancestryName) {
					resultIdx = ancestryIdx
					result = ancestryName
					break ancestryChoice
				}
			}
			fmt.Printf("Please choose one of the ancestries %v\n", choices)
		}
	}

	fmt.Printf("You have chosen %s ancestry\n", result)
	return result, append(choices[:resultIdx], choices[resultIdx+1:]...)
}

func ChooseProfession(choices []ProfessionName, ancestries map[ProfessionName]Profession) (ProfessionName, []ProfessionName) {
	fmt.Println("What is your profession?")

	for ancestryIdx, ancestryName := range choices {
		ancestry := ancestries[ancestryName]
		fmt.Printf("%d. %s: %s\n", ancestryIdx+1, ancestryName, ancestry.Description)
	}

	fmt.Println("Choose one of the above")

	var resultIdx int
	var result ProfessionName
professionChoice:
	for {
		var command string
		fmt.Scan(&command)
		if professionNameIdx, err := strconv.Atoi(command); err == nil {
			if 0 < professionNameIdx && professionNameIdx <= len(choices) {
				resultIdx = professionNameIdx - 1
				result = choices[resultIdx]
				break professionChoice
			}
			fmt.Printf("Please choose option by number between 1 and %d\n", len(choices))
		} else {
			for professionNameIdx, professionName := range choices {
				if command == fmt.Sprint(professionName) {
					resultIdx = professionNameIdx
					result = professionName
					break professionChoice
				}
			}
			fmt.Printf("Please choose one of the profession %v\n", choices)
		}
	}
	fmt.Printf("You have chosen %s profession\n", result)
	return result, append(choices[:resultIdx], choices[resultIdx+1:]...)
}

func ChooseFit(choices []FitName, fits map[FitName]Fit) (FitName, []FitName) {
	fmt.Println("You have Leveled Up !")

	for fitIdx, fitName := range choices {
		fit := fits[fitName]
		fmt.Printf("%d. %s: %v\n", fitIdx+1, fitName, fit.StatChange)
	}

	fmt.Println("Choose fit from above")

	var resultIdx int
	var result FitName
fitChoice:
	for {
		var command string
		fmt.Scan(&command)
		if fitIdx, err := strconv.Atoi(command); err == nil {
			if 0 < fitIdx && fitIdx <= len(choices) {
				resultIdx = fitIdx - 1
				result = choices[resultIdx]
				break fitChoice
			}
			fmt.Printf("Please choose option by number between 1 and %d\n", len(choices))
		} else {
			for fitIdx, fitName := range choices {
				if command == fmt.Sprint(fitName) {
					resultIdx = fitIdx
					result = fitName
					break fitChoice
				}
			}
			fmt.Printf("Please choose one of the following fits %v\n", choices)
		}
	}

	fmt.Printf("You have chosen %s fit\n", result)
	return result, append(choices[:resultIdx], choices[resultIdx+1:]...)
}

func (heroPtr *Hero) levelUp(game *Game) {
	// Hero level should be increased by 1
	// On level 3 choose second profession
	// On level 5 choose advanced ancestry
	// On every other level up add fit from the table of fits
	// On level 10 print "You won"
	heroPtr.Level += 1

	switch heroPtr.Level {
	case 1, 2, 4, 6, 7, 8, 9:
		fitChoices, error := game.FitDeck.Draw(2)
		if error == nil {
			fitName, discardedCards := ChooseFit(fitChoices, Fits)
			game.Discard.FitDeck.AddAll(discardedCards)
			heroPtr.Fits = append(heroPtr.Fits, fitName)

		}
	case 3:
		professionChoices, error := game.ProfessionDeck.Draw(2)
		if error == nil {
			professionName, discardedCards := ChooseProfession(professionChoices, professions)
			game.Discard.ProfessionDeck.AddAll(discardedCards)
			heroPtr.Profession = append(heroPtr.Profession, professionName)
		} else {
			panic(error.Error())
		}
	case 5:
		heroAncestryName := heroPtr.Ancestry[0]
		heroAncestry := ancestries[heroAncestryName]
		ancestryName, discardedCards := ChooseAncestry(heroAncestry.Progression, ancestries)
		game.Discard.AncestryDeck.AddAll(discardedCards)
		heroPtr.Ancestry = append(heroPtr.Ancestry, ancestryName)
	default:
		fmt.Println("You won !")
		return
	}
}
