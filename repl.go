package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	// "github.com/aayushtmG/pokedexcli/internal"
	"github.com/aayushtmG/pokedexcli/internal/pokeapi"
)

type config struct {
	caughtPokemon map[string]pokeapi.Pokemon
	pokeapiClient        pokeapi.Client
	nextLocationsURL     *string
	previousLocationsURL *string
}

func startRepl(cfg *config) {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedox > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		args := []string{}

		if len(words) > 1 {
			args = words[1:]	
		}

		commandName := words[0]

		command, exists := getCommands()[commandName]
		
		if exists {
			err := command.callback(cfg,args...)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config,...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Get the next page of locations",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Get the previous page of locations",
			callback:    commandMapb,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"explore": {
			name:        "explore <location name>",
			description: "Explore a area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch <pokemon name>",
			description: "Catch a pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect <pokemon name>",
			description: "inspect a catched pokemon",
			callback:    commandInspect,
		},
	}

}
