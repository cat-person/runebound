package main

import (
	"fmt"
	"strconv"
)

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
