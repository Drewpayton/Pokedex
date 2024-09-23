package main

import (
	"errors"
	"fmt"
)

func commandPokedex(cfg *config, arg string) error {
	if len(cfg.caughtPokemon) == 0 {
		return errors.New("you have no pokemon")
	}

	fmt.Println("Your Pokedex:")
	for k, _ := range cfg.caughtPokemon {
		fmt.Printf(" - %s\n", k)
	}

	return nil
}