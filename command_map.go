package main

import (
	"errors"
	"fmt"
)

func commandMapf(cfg *config) error {

	locationsResp, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationsURL)

	if err != nil {
		return err
	}
	
	cfg.nextLocationsURL = locationsResp.Next
	cfg.previousLocationsURL = locationsResp.Previous


	fmt.Println("DISPLAYING MAP LOCATIONS----------------------") 
	for _, location := range locationsResp.Results {
		fmt.Println(location.Name)
	}
	return nil
}

func commandMapb(cfg *config) error {
	if cfg.previousLocationsURL == nil {
		return errors.New("you're on the first page")
	}

	locationsResp, err := cfg.pokeapiClient.ListLocations(cfg.previousLocationsURL)
	if err != nil {
		return err
	}
	cfg.nextLocationsURL = locationsResp.Next
	cfg.previousLocationsURL = locationsResp.Previous


	fmt.Println("DISPLAYING MAP LOCATIONS----------------------") 
	for _, location := range locationsResp.Results {
		fmt.Println(location.Name)
	}
	return nil
}


