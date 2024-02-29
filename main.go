package main

import (
	"fmt"
	"strconv"
)

func main() {
	hero := Hero{}

	heroName := ChooseHeroName()
	fmt.Printf("You have chosen %s ancestry\n", heroName)
	hero.Name = heroName

	ancestryName := ChooseAncestry([]AncestryName{Elf, Dwarf, Human, Dragonkin}, Ancestries)
	hero.Ancestry = append(hero.Ancestry, ancestryName)

	professionName := ChooseProfession([]ProfessionName{Blacksmith, Herbalist, Fighter, Criminal}, Professions)
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

func ChooseAncestry(choices []AncestryName, ancestries map[AncestryName]Ancestry) AncestryName {
	fmt.Println("What is your ancestry?")

	for ancestryIdx, ancestryName := range choices {
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
			if 0 < ancestryIdx && ancestryIdx <= len(choices) {
				result = choices[ancestryIdx-1]
				break ancestryChoice
			} else {
				fmt.Printf("Please choose option by number between 1 and %d\n", len(choices))
			}
		} else {
			for _, ancestryName := range choices {
				if command == fmt.Sprint(ancestryName) {
					result = ancestryName

					return ancestryName
				}
			}
			fmt.Printf("Please choose one of the ancestries %v\n", choices)
		}
	}
	fmt.Printf("You have chosen %s ancestry\n", result)
	return result
}

func ChooseProfession(choices []ProfessionName, ancestries map[ProfessionName]Profession) ProfessionName {
	fmt.Println("What is your profession?")

	for ancestryIdx, ancestryName := range choices {
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
			if 0 < professionNameIdx && professionNameIdx <= len(choices) {
				result = choices[professionNameIdx-1]
				break professionChoice
			}
			fmt.Printf("Please choose option by number between 1 and %d\n", len(choices))
		} else {
			for _, professionName := range choices {
				if command == fmt.Sprint(professionName) {
					result = professionName
					break professionChoice
				}
			}
			fmt.Printf("Please choose one of the profession %v\n", choices)
		}
	}
	fmt.Printf("You have chosen %s profession\n", result)
	return result
}

func levelUp(heroPtr *Hero, professions map[ProfessionName]Profession, ancestries map[AncestryName]Ancestry) {
	// Hero level should be increased by 1
	// On level 3 choose second profession
	// On level 5 choose advanced ancestry
	// On every other level up add fit from the table of fits
	// On level 10 print "You won"
	heroPtr.Level += 1

	switch heroPtr.Level {
	case 1, 2, 4, 6:
		heroPtr.Fits = append(heroPtr.Fits, ChooseFit([]FitName{Mountain, Stoic, SpiritBound, Scholar, SilverTongue, IronGuts}, Fits))
	case 3:
		professionNames := make([]ProfessionName, len(professions))
		professionNameIdx := 0
		for professionName := range professions {
			professionNames[professionNameIdx] = professionName
			professionNameIdx++
		}
		heroPtr.Profession = append(heroPtr.Profession, ChooseProfession(professionNames, professions))
	case 5:
		heroAncestryName := heroPtr.Ancestry[0]
		heroAncestry := ancestries[heroAncestryName]
		heroPtr.Ancestry = append(heroPtr.Ancestry, ChooseAncestry(heroAncestry.Progression, ancestries))
	case 7, 8, 9:
		heroPtr.Fits = append(heroPtr.Fits, ChooseFit([]FitName{Bully, Prude, Elder, Bookworm, Sensualist, Bear}, Fits))
	default:
		fmt.Println("You won !")
		return
	}
}

func ChooseFit(choices []FitName, fits map[FitName]Fit) FitName {
	fmt.Println("You have Leveled Up !")

	for fitIdx, fitName := range choices {
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
			if 0 < fitIdx && fitIdx <= len(choices) {
				result = choices[fitIdx-1]
				break fitChoice
			}
			fmt.Printf("Please choose option by number between 1 and %d\n", len(choices))
		} else {
			for _, fitName := range choices {
				if command == fmt.Sprint(fitName) {
					result = fitName
					break fitChoice
				}
			}
			fmt.Printf("Please choose one of the following fits %v\n", choices)
		}
	}

	fmt.Printf("You have chosen %s fit\n", result)
	return result
}
