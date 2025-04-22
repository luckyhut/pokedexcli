package main

import (
	"fmt"
	"github.com/luckyhut/pokedexcli/internal/pokecache"
	"os"
)

func commandExit(c *config, cache *pokecache.Cache) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
