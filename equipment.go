package main

type EquipmentType string

const (
	Helm   EquipmentType = "Helm"
	Armour EquipmentType = "Armour"
	Weapon EquipmentType = "Weapon"
	Book   EquipmentType = "Book"
)

type Equipment struct {
	Name    string
	Type    EquipmentType
	Stats   map[StatName]int
	Runes   []RuneName
	Actions []ActionName
}