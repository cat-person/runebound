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
	fmt.Printf("You have chosen %s ancestry\n", ancestryName)
	hero.Ancestry = append(hero.Ancestry, ancestryName)

	professionName := ChooseProfession([]ProfessionName{Blacksmith, Herbalist, Fighter, Criminal}, Professions)
	fmt.Printf("You have chosen %s profession\n", professionName)
	hero.Profession = append(hero.Profession, professionName)

	fmt.Printf("Congratulations you are %s %s now!\n", ancestryName, professionName)

	fmt.Printf("Your stats are: %v\n", hero.getStats())
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
	for {
		var command string
		fmt.Scan(&command)
		if ancestryIdx, err := strconv.Atoi(command); err == nil {
			if 0 < ancestryIdx && ancestryIdx <= len(choices) {
				return choices[ancestryIdx-1]
			} else {
				fmt.Printf("Please choose option by number between 1 and %d\n", len(choices))
			}
		} else {
			for _, ancestryName := range choices {
				if command == fmt.Sprint(ancestryName) {
					return ancestryName
				}
			}
			fmt.Printf("Please choose one of the ancestries %v\n", choices)
		}
	}
}

func ChooseProfession(choices []ProfessionName, ancestries map[ProfessionName]Profession) ProfessionName {
	fmt.Println("What is your profession?")

	for ancestryIdx, ancestryName := range choices {
		ancestry := ancestries[ancestryName]
		fmt.Printf("%d. %s: %s\n", ancestryIdx+1, ancestryName, ancestry.Description)
	}

	fmt.Println("Choose one of the above")
	for {
		var command string
		fmt.Scan(&command)
		if ancestryIdx, err := strconv.Atoi(command); err == nil {
			if 0 < ancestryIdx && ancestryIdx <= len(choices) {
				return choices[ancestryIdx-1]
			} else {
				fmt.Printf("Please choose option by number between 1 and %d\n", len(choices))
			}
		} else {
			for _, ancestryName := range choices {
				if command == fmt.Sprint(ancestryName) {
					return ancestryName
				}
			}
			fmt.Printf("Please choose one of the profession %v\n", choices)
		}
	}
}

func levelUp(heroPtr *Hero) {
	for {
		fmt.Println("Press Enter to level up")
		var command string
		fmt.Scan(&command)

		heroPtr.Level += 1

		switch heroPtr.Level {

		case 1, 2, 4, 6:
			heroPtr.Fits = append(heroPtr.Fits, ChooseFit([]FitName{Mountain, Stoic, SpiritBound, Scholar, SilverTongue, IronGuts}, Fits))

		case 7, 8, 9:
			heroPtr.Fits = append(heroPtr.Fits, ChooseFit([]FitName{Bully, Prude, Elder, Bookworm, Sensualist, Bear}, Fits))
		default:
			fmt.Println("You won !")
			return
		}
	}
	// On level 3 choose second profession
	// On level 5 choose advanced ancestry
	// On every other level up add fit from the table of fits
	// On level 10 print "You won"
}

func ChooseFit(choices []FitName, fits map[FitName]Fit) FitName {
	fmt.Println("What is your profession?")

	for fitIdx, fitName := range choices {
		fit := fits[fitName]
		fmt.Printf("%d. %s: %v\n", fitIdx+1, fitName, fit.StatChange)
	}

	fmt.Println("Choose one of the above")
	for {
		var command string
		fmt.Scan(&command)
		if fitIdx, err := strconv.Atoi(command); err == nil {
			if 0 < fitIdx && fitIdx <= len(choices) {
				return choices[fitIdx-1]
			} else {
				fmt.Printf("Please choose option by number between 1 and %d\n", len(choices))
			}
		} else {
			for _, fitName := range choices {
				if command == fmt.Sprint(fitName) {
					return fitName
				}
			}
			fmt.Printf("Please choose one of the profession %v\n", choices)
		}
	}
}
