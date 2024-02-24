package main

import "encoding/json"

type ActionName string

type Actions []ActionName

// Standard
const (
	Punch ActionName = "Punch"
)

const (
	Flank ActionName = "Flank"
)

// ?
type ActionTarget uint16

// ??
const (
	Self ActionTarget = iota
	Opponent
)

// ???
type ActionDescriptor struct {
	Name        string
	Description string
	Price       []RuneSide
	action      Action
	target      ActionTarget
}

type Action struct {
	name   string
	params map[string]string
}

// ??????
func (actionDescription ActionDescriptor) String() string {
	result, err := json.MarshalIndent(actionDescription, "", "  ")
	if err == nil {
		return string(result)
	} else {
		return err.Error()
	}
}
