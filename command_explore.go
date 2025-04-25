package main

import (
	"fmt"
	"github.com/luckyhut/pokedexcli/internal/pokeapi"
	"github.com/luckyhut/pokedexcli/internal/pokecache"
)

func commandExplore(c *config, cache *pokecache.Cache, location *string) error {
	fullURL := pokeapi.BaseURL + "location-area/" + *location
	cached, found := cache.Get(fullURL)
	// it's in the cache
	if found {
		fmt.Println("Cache hit")
		encounters, err := pokeapi.UnmarshalPokemonEncounter(cached)
		if err != nil {
			return err
		}
		printEncounters(encounters)
	}

	// it's not in the cache
	encounters, err := c.pokeapiClient.Explore(&fullURL, cache)
	if err != nil {
		return err
	}
	printEncounters(encounters)
	return nil
}

func printEncounters(encounters pokeapi.LocationArea) error {
	if len(encounters.PokemonEncounters) == 0 {
		fmt.Println("No Pokemon were found here")
		return nil
	}
	fmt.Println("Found Pokemon:")
	for _, e := range encounters.PokemonEncounters {
		fmt.Println("- ", e.Pokemon.Name)
	}
	return nil
}
