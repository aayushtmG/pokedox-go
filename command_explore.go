package main

import (
	"fmt"
)
func commandExplore(cfg *config,params ...string) error{
		if len(params) == 0 {
			return fmt.Errorf("error: include id or name of the location area")
		}
		resp, err := cfg.pokeapiClient.FindPokemons(params[0])

		if err != nil {
			return err
		}

		fmt.Printf("Exploring %s...\n",resp.Name)
		fmt.Println("Found Pokemon:")
		for _, pokeObj := range resp.PokemonsFound {
				fmt.Println(" - ",pokeObj.P.Name)
		}

		return nil
}