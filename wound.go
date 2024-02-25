package main

type InjuryName string

const (
	JustAScratch InjuryName = "JustAScratch"
)

type Wound struct {
	Injury
}

// Negative effect
type Injury struct {
	Name        string
	Description string
}
