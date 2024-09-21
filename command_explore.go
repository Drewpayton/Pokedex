package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, areaName string) error {
	if areaName == "" {
		return errors.New("you must provide a location name")
	} 
	
	location, err := cfg.pokeapiClient.GetLocation(areaName)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", location.Location.Name)
	fmt.Println("Found Pokemon:")
	for _, enc := range location.PokemonEncounter {
		fmt.Printf(" - %s\n", enc.Pokemon.Name)
	}

	return nil
}