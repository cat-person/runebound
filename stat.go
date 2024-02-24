package main

type StatName string

const (
	Health     StatName = "Health"
	Arm        StatName = "Armour"
	Initiative StatName = "Initiative"
)

type Stat struct {
	StatName StatName
	Stat     int
}
