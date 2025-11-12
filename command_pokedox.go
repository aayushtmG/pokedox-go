package main

import "fmt"

func commandPokedox(c *config,args ...string) error {
	
	if len(c.caughtPokemon) == 0 {
		fmt.Println("Your pokedox is empty, Go Hunt!!!")
		return nil
	}
	
	fmt.Println("Your pokedox:")
	for _, pokemon := range c.caughtPokemon {
		fmt.Printf(" - %s\n",pokemon.Name)
	}

	return nil
}