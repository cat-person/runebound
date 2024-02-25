package main

type StatName string

const (
	MaxWound   StatName = "MaxWound"
	Arm        StatName = "Armour"
	Initiative StatName = "Initiative"
)

type Stat struct {
	StatName StatName
	Stat     int
}
