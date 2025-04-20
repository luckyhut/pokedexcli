package main

import (
	"fmt"
	"github.com/luckyhut/pokedexcli/internal/pokeapi"
)

func commandMap(c *config) error {
	locations, err := c.pokeapiClient.Locations(c.nextLocationsURL)
	if err != nil {
		return err
	}
	printLocationArea(&locations)
	c.nextLocationsURL = locations.Next
	c.prevLocationsURL = locations.Previous
	return nil
}

func commandMapb(c *config) error {
	locations, err := c.pokeapiClient.Locations(c.prevLocationsURL)
	if err != nil {
		return err
	}
	printLocationArea(&locations)
	c.nextLocationsURL = locations.Next
	c.prevLocationsURL = locations.Previous
	return nil
}

func printLocationArea(l *pokeapi.Locations) {
	for _, location := range l.Results {
		fmt.Println(location.Name)
	}
}
