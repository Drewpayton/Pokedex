package main

import (
	"errors"
	"fmt"
)

// Go to the next page
func commandMapf(cfg *config) error {
	// Next location URL is the next set up locations and it sends that to list locations to get them.
	locationResp, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationURL)
	if err != nil {
		return err
	}

	// Updating the next and prev values so we can keep traversing through locations.
	cfg.nextLocationURL = locationResp.Next
	cfg.prevLoactionsURL = locationResp.Previous

	// This loops through each location and prints them out
	for _, loc := range locationResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

// Go to the previous page.
func commandMapb(cfg *config) error {
	// Checks if there on the "first page"
	if cfg.prevLoactionsURL == nil {
		return errors.New("you're already on the first page")
	}

	// Sends the prev location to list location so it can send a req to it.
	locationResp, err := cfg.pokeapiClient.ListLocations(cfg.prevLoactionsURL)
	if err != nil {
		return err
	}

	// Updating the next and prev values to keep traversing through locations.
	cfg.nextLocationURL = locationResp.Next
	cfg.prevLoactionsURL = locationResp.Previous

	// Loops through each loation and prints them out.
	for _, loc := range locationResp.Results {
		fmt.Println(loc.Name)
	}
	
	return nil
}