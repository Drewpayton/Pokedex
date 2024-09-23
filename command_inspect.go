package main

import "fmt"

func commandInspect(cfg *config, pokemonName string) error {

	if val, ok := cfg.caughtPokemon[pokemonName]; ok {
		fmt.Printf("Name: %s\n")
		fmt.Printf("Height: %s\n")
		fmt.Printf("Weight: %s\n")
		fmt.Printf("Stats: %s\n")
		fmt.Printf("Types: %s\n")
	} else {

	}
	
}