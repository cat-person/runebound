package main

type AncestryName string

// Elf Progression
const (
	Elf    AncestryName = "Elf"
	Faerun AncestryName = "Faerun" // Wild elfs
	Sithi  AncestryName = "Sithi"  // Magic
	Drow   AncestryName = "Drow"   // Dark edgy stuff
)

// Dwarf Progression
const (
	Dwarf         AncestryName = "Dwarf"
	GoldDwarves   AncestryName = "GoldDwarves"   // More rewards / More life
	ShieldDwarves AncestryName = "ShieldDwarves" // Stronger and better protected dwarfs
	Duergar       AncestryName = "Duergar"       // Sneaky evil dwarfs
)

// Human progression
const (
	Human     AncestryName = "Human"
	Calishite AncestryName = "Calishite" // Diplomacy + More loot
	Mulan     AncestryName = "Mulan"     // Fighty humans
	Rashemi   AncestryName = "Rashemi"   // Magical humans
)

// Dragonborn
const (
	Dragonborn       AncestryName = "Dragonborn"
	BronzeDragonborn AncestryName = "BronzeDragonborn"
	BlackDragonborn  AncestryName = "BlackDragonborn"
	RedDragonborn    AncestryName = "RedDragonborn"
)

var Ancestries map[AncestryName]Ancestry = map[AncestryName]Ancestry{
	Elf:    Ancestry{"You know, Elf", []AncestryName{Faerun, Sithi, Drow}},
	Faerun: Ancestry{"Wild Elf", []AncestryName{}},
	Sithi:  Ancestry{"Magic Elf", []AncestryName{}},
	Drow:   Ancestry{"Dark, edgy Elf", []AncestryName{}},

	Dwarf:         Ancestry{"5'11 Drunks", []AncestryName{GoldDwarves, ShieldDwarves, Duergar}},
	GoldDwarves:   Ancestry{"More rewards / More life", []AncestryName{}},
	ShieldDwarves: Ancestry{"Stronger and better protected dwarfs", []AncestryName{}},
	Duergar:       Ancestry{"Sneaky evil dwarfs", []AncestryName{}},

	Human:     Ancestry{"6'0 Drunks", []AncestryName{Calishite, Mulan, Rashemi}},
	Calishite: Ancestry{"Diplomacy + More loot", []AncestryName{}},
	Mulan:     Ancestry{"6'0 Drunks", []AncestryName{}},
	Rashemi:   Ancestry{"6'0 Drunks", []AncestryName{}},

	Dragonborn:       Ancestry{"Dragons, you like dragons right?", []AncestryName{Calishite, Mulan, Rashemi}},
	BronzeDragonborn: Ancestry{"Dragonborn with more defence / atack", []AncestryName{}},
	BlackDragonborn:  Ancestry{"Evil sneaky magical Dragonborn", []AncestryName{}},
	RedDragonborn:    Ancestry{"Evil fighty magical Dragonborn", []AncestryName{}},
}

type Ancestry struct {
	Description string
	Progression []AncestryName
}
