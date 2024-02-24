package main

import "fmt"

// Hello returns a greeting for the named person.
func main() {
	fight(getHero("Bob"))
}

func fight(hero Hero) {
	// heroRunes := hero.castRunes()
	// fmt.Println(heroRunes)
	fmt.Println(hero)

	// monsteRunes := monster.castRunes()
	// fmt.Println(monsteRunes)
}
