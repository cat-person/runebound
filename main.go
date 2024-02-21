package main

import "fmt"

// Hello returns a greeting for the named person.
func main() {

}

func fight(hero Hero, monster Monster) {
	heroRunes := hero.castRunes()
	fmt.Println(heroRunes)

	monsteRunes := monster.castRunes()
	fmt.Println(monsteRunes)
}
