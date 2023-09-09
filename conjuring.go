package main

import "fmt"

// Familiar is a mythical creature summoned to serve a witch.
type Familiar struct {
	Name   string
	Health int
}

// SummonFamiliar creates and summons a familiar to assist the witch.
func SummonFamiliar() {
	fmt.Println("Summoning a Familiar...")
	// TODO: Implement familiar summoning logic here
	familiar := Familiar{Name: "Shadow", Health: 100}
	fmt.Printf("Summoned Familiar: %s with %d health\n", familiar.Name, familiar.Health)
}

