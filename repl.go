package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Drewpayton/Pokedex/internal/pokeapi"
)

func startRepl(cfg *config) {
	// Creating a scanner
	reader := bufio.NewScanner(os.Stdin)
	areaName := ""

	// repl for loop that exits when user inputs "exit"
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanWords(reader.Text())

		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		if len(words) > 1 {
			areaName = words[1]
		}
		
		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(cfg, areaName)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown Command")
			continue
		}
	}
}

func cleanWords(text string) []string {
	newText := strings.ToLower(text)
	output := strings.Fields(newText)
	return output

}

type cliCommand struct {
	name 		string
	description string
	callback 	func(cfg *config, input string) error
}

type config struct {
	pokeapiClient    pokeapi.Client
	nextLocationURL  *string
	prevLoactionsURL *string
	caughtPokemon map[string]pokeapi.Pokemon
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help" : {
			name: "help",
			description: "Displays a help message",
			callback: commandHelp,
		},
		"exit" : {
			name: "exit",
			description: "Exit the Pokedex",
			callback: commandExit,
		},
		"map" : {
			name: "map",
			description: "Get the next page of locations",
			callback: commandMapf,
		},
		"mapb" : {
			name: "mapb",
			description: "Get the previous page of locations.",
			callback: commandMapb,
		},
		"explore": {
			name: "explore",
			description: "Explore pokemon in given area.",
			callback: commandExplore,
		},
		"catch": {
			name: "catch",
			description: "Catch a specific pokemon.",
			callback: commandCatch,
		},
		"inspect" : {
			name: "inspect",
			description: "Inspect a caught pokemon.",
			callback: commandInspect,
		},
	}
}

