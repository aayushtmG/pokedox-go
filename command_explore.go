package main

import (
	"errors"
	"fmt"
)
func commandExplore(cfg *config,args ...string) error{
		if len(args) != 1 {
			return errors.New("you must provide at least one location name and only one")
		}

		location, err := cfg.pokeapiClient.GetLocation(args[0])
		if err != nil {
			return err
		}

		fmt.Printf("Exploring %s...\n",location.Name)
		fmt.Println("Found Pokemon:")
		for _, enc := range location.PokemonEncounters {
				fmt.Println(" - ",enc.Pokemon.Name)
		}

		return nil
}