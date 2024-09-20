package main

import "fmt"

func commandExplore(cfg *config, areaName string) error {
	locationResp, err := cfg.pokeapiClient.ListPokemon(cfg.nextLocationURL, areaName)
	if err != nil {
		return err
	}

	fmt.Print(locationResp)

	return nil
}