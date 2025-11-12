package main

import (
	"fmt"
)

func commandInspect(c *config,args ...string) error{
	if len(args) != 1 {
		return fmt.Errorf("give name of the pokemon to inspect")
	}

	name := args[0]
	pokemon , ok := c.caughtPokemon[name]
	if !ok {
		return fmt.Errorf("you have not caught that pokemon")
	}

	fmt.Printf("Name: %s\n",pokemon.Name)
	fmt.Printf("Height: %d\n",pokemon.Height)
	fmt.Printf("Weight: %d\n",pokemon.Weight)
	fmt.Println("Stats:")
	for _, val := range pokemon.Stats {
		fmt.Printf(" -%s: %d\n",val.Stat.Name,val.BaseStat)
	}

	fmt.Println("Types:")
	for _, val := range pokemon.Types {
		fmt.Printf(" -%s\n",val.Type.Name)
	}

	return nil
}