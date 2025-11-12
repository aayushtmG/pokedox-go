package main

import (
	"fmt"
	"math/rand"
)



func commandCatch(c *config,args ...string) error{
	if len(args) != 1 {
		return fmt.Errorf("give name of the pokemon to catch")
	}

	pokemonName := args[0]
	pokemon , err :=  c.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return fmt.Errorf("error: %w",err)
	}

	res := rand.Intn(pokemon.BaseExperience)
	fmt.Printf("Throwing a Pokeball at %s...\n",pokemon.Name)	
	if res > 40 {
		fmt.Printf("%s escaped!\n",pokemon.Name)
		return nil
	}

	fmt.Printf("%s was caught!\n",pokemon.Name)
fmt.Println("You may now inspect it with the inspect command.")

	c.caughtPokemon[pokemonName] = 	pokemon
	return nil
}