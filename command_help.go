package main

import (
	"fmt"
	"github.com/luckyhut/pokedexcli/internal/pokecache"
)

func commandHelp(c *config, cache *pokecache.Cache) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")

	for k, v := range commands {
		fmt.Printf("%s: %s\n", k, v.description)
	}
	return nil
}
