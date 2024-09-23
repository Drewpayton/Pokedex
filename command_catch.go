package main

import "errors"

func commandCatch(cfg *config, pokemonName string) error{
	if pokemonName == "" {
		return errors.New("you need to include a pokemon")
	}

	
	return nil
}