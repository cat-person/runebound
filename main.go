package main

import "fmt"

func main() {
	hero := Hero{}
	ancestryName := ChooseAncestry([]AncestryName{Elf, Dwarf, Human, Dragonborn}, Ancestries)
	fmt.Printf("You have chosen %s ancestry\n", ancestryName)
	hero.Ancestry = ancestryName

}

func fight(hero Hero) {
	// heroRunes := hero.castRunes()
	// fmt.Println(heroRunes)
	fmt.Println(hero)

	// monsteRunes := monster.castRunes()
	// fmt.Println(monsteRunes)
}

func levelUp() {

}
