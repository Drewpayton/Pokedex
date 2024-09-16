package main

import (
	"time"

	"github.com/Drewpayton/Pokedex/internal/pokeapi"
)

func main() {
	// An instance of http.Client with a timeout of 5 seconds
	pokeClient := pokeapi.NewClient(5 * time.Second)
	cfg := &config{
		pokeapiClient: pokeClient,
	}
	startRepl(cfg)
}