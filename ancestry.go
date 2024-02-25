package main

type AncestryName string

const (
	Elf AncestryName = "Elf"
)

type Ancestry struct {
	Name        string
	Description string
	Progression []Ancestry
}
