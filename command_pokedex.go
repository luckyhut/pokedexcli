package main

import (
	"fmt"
	"github.com/luckyhut/pokedexcli/internal/pokecache"
)

func commandPokedex(c *config, cache *pokecache.Cache, _ *string) error {
	fmt.Println("Your Pokedex:")
	for k, _ := range c.pokedex {
		fmt.Println("  -", k)
	}

	return nil
}
