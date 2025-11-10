package main

import (
	"github.com/aayushtmG/pokedexcli/internal"
	"github.com/aayushtmG/pokedexcli/internal/pokeapi"
	"time"
)


func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	pokeClient.Cache = internal.NewCache(5 * time.Second)
	cfg := &config{
		pokeapiClient: pokeClient,
	}
	startRepl(cfg)
}
