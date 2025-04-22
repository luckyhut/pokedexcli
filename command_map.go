package main

import (
	"fmt"
	"github.com/luckyhut/pokedexcli/internal/pokeapi"
	"github.com/luckyhut/pokedexcli/internal/pokecache"
)

func commandMap(c *config, cache *pokecache.Cache) error {
	cached, found := cache.Get(*c.nextLocationsURL)
	// it's in the cache
	if found {
		fmt.Println("Cache hit")
		locations, err := pokeapi.UnmarshalLocations(cached)
		if err != nil {
			return err
		}
		printLocationArea(&locations)
		return nil
	}
	// it's not in the cache
	locations, err := c.pokeapiClient.Locations(c.nextLocationsURL, cache)
	if err != nil {
		return err
	}
	printLocationArea(&locations)
	c.nextLocationsURL = locations.Next
	c.prevLocationsURL = locations.Previous
	return nil
}

func commandMapb(c *config, cache *pokecache.Cache) error {
	cached, found := cache.Get(*c.prevLocationsURL)
	// it's in the cache
	if found {
		fmt.Println("Cache hit")

		locations, err := pokeapi.UnmarshalLocations(cached)
		if err != nil {
			return err
		}
		printLocationArea(&locations)
		return nil
	}
	locations, err := c.pokeapiClient.Locations(c.prevLocationsURL, cache)
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
