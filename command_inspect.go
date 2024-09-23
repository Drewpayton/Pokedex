package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, pokemonName string) error {


	if val, ok := cfg.caughtPokemon[pokemonName]; ok {
		fmt.Printf("Name: %s\n", val.Name)
		fmt.Printf("Height: %d\n", val.Height)
		fmt.Printf("Weight: %d\n", val.Weight)
		fmt.Printf("Stats:\n")
		for _, pokeStats := range val.Stats {
			fmt.Printf("	-%s: %d\n", pokeStats.Stat.Name, pokeStats.BaseStat)
		}
		fmt.Printf("Types:\n")
		for _, pokeTypes := range val.Types {
			fmt.Printf("	- %s", pokeTypes.Type.Name)
		}
	} else {
		return errors.New("you have not caught that pokemon")
	}
	return nil
}