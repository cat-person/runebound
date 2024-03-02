package main

import (
	"fmt"
	"strconv"
)

func main() {
	game := Game{
		Deck[AncestryName]{[]AncestryName{Elf, Dwarf, Human, Dragonkin}},
		Deck[ProfessionName]{[]ProfessionName{Blacksmith, Herbalist, Fighter, Criminal}},
		Deck[FitName]{},
	}

	hero := Hero{}

	heroName := ChooseHeroName()
	fmt.Printf("You have chosen %s ancestry\n", heroName)
	hero.Name = heroName

	ancestryName := ChooseAncestry(game.AncestryDeck, Ancestries)
	hero.Ancestry = append(hero.Ancestry, ancestryName)

	professionName := ChooseProfession(game.ProfessionDeck, Professions)
	hero.Profession = append(hero.Profession, professionName)

	fmt.Println()
	fmt.Printf("Congratulations you are %s %s now!\n", ancestryName, professionName)

	for hero.Level < 10 {
		fmt.Printf("Your stats are: %v\n", hero.getStats())
		fmt.Println()
		fmt.Println("Press Enter to level up")
		fmt.Scanln()
		levelUp(&hero, Professions, Ancestries)
	}

}

func ChooseHeroName() string {
	fmt.Println("What is your name?")
	var name string
	fmt.Scan(&name)
	return name
}

func ChooseAncestry(deck Deck[AncestryName], ancestries map[AncestryName]Ancestry) AncestryName {
	fmt.Println("What is your ancestry?")

	for ancestryIdx, ancestryName := range deck.cards {
		ancestry := ancestries[ancestryName]
		fmt.Printf("%d. %s: %s\n", ancestryIdx+1, ancestryName, ancestry.Description)
	}

	fmt.Println("Choose one of the above")
	var result AncestryName

ancestryChoice:
	for {
		var command string
		fmt.Scan(&command)
		if ancestryIdx, err := strconv.Atoi(command); err == nil {
			if 0 < ancestryIdx && ancestryIdx <= len(deck.cards) {
				result = deck.cards[ancestryIdx-1]
				break ancestryChoice
			} else {
				fmt.Printf("Please choose option by number between 1 and %d\n", len(deck.cards))
			}
		} else {
			for _, ancestryName := range deck.cards {
				if command == fmt.Sprint(ancestryName) {
					result = ancestryName
					return ancestryName
				}
			}
			fmt.Printf("Please choose one of the ancestries %v\n", deck.cards)
		}
	}
	fmt.Printf("You have chosen %s ancestry\n", result)
	return result
}

func ChooseProfession(deck Deck[ProfessionName], ancestries map[ProfessionName]Profession) ProfessionName {
	fmt.Println("What is your profession?")

	for ancestryIdx, ancestryName := range deck.cards {
		ancestry := ancestries[ancestryName]
		fmt.Printf("%d. %s: %s\n", ancestryIdx+1, ancestryName, ancestry.Description)
	}

	fmt.Println("Choose one of the above")

	var result ProfessionName
professionChoice:
	for {
		var command string
		fmt.Scan(&command)
		if professionNameIdx, err := strconv.Atoi(command); err == nil {
			if 0 < professionNameIdx && professionNameIdx <= len(deck.cards) {
				result = deck.cards[professionNameIdx-1]
				break professionChoice
			}
			fmt.Printf("Please choose option by number between 1 and %d\n", len(deck.cards))
		} else {
			for _, professionName := range deck.cards {
				if command == fmt.Sprint(professionName) {
					result = professionName
					break professionChoice
				}
			}
			fmt.Printf("Please choose one of the profession %v\n", deck.cards)
		}
	}
	fmt.Printf("You have chosen %s profession\n", result)
	return result
}

func ChooseFit(deck Deck[FitName], fits map[FitName]Fit) FitName {
	fmt.Println("You have Leveled Up !")

	for fitIdx, fitName := range deck.cards {
		fit := fits[fitName]
		fmt.Printf("%d. %s: %v\n", fitIdx+1, fitName, fit.StatChange)
	}

	fmt.Println("Choose fit from above")

	var result FitName

fitChoice:
	for {
		var command string
		fmt.Scan(&command)
		if fitIdx, err := strconv.Atoi(command); err == nil {
			if 0 < fitIdx && fitIdx <= len(deck.cards) {
				result = deck.cards[fitIdx-1]
				break fitChoice
			}
			fmt.Printf("Please choose option by number between 1 and %d\n", len(deck.cards))
		} else {
			for _, fitName := range deck.cards {
				if command == fmt.Sprint(fitName) {
					result = fitName
					break fitChoice
				}
			}
			fmt.Printf("Please choose one of the following fits %v\n", deck.cards)
		}
	}

	fmt.Printf("You have chosen %s fit\n", result)
	return result
}

func levelUp(heroPtr *Hero, professions map[ProfessionName]Profession, ancestries map[AncestryName]Ancestry) {
	// Hero level should be increased by 1
	// On level 3 choose second profession
	// On level 5 choose advanced ancestry
	// On every other level up add fit from the table of fits
	// On level 10 print "You won"
}
