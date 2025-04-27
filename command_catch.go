package main

import (
	"fmt"
	"github.com/luckyhut/pokedexcli/internal/pokecache"
	"math/rand"
)

const CatchRateScalar = 45.0

func commandCatch(c *config, cache *pokecache.Cache, pokemon_name *string) error {
	// get its info, we need to get base experience
	pokemon := c.pokeapiClient.GetPokemonInfo(pokemon_name)
	// use math/rand to calculate a catch rate
	rate := CatchRateScalar * (1.0 / float64(pokemon.BaseExperience))
	randNum := rand.Float64()
	fmt.Println(randNum, rate)
	if randNum < rate {
		fmt.Println(pokemon.Name, "was caught")
		c.pokedex[pokemon.Name] = pokemon
		return nil
	}

	fmt.Println("failed to catch ", pokemon.Name)

	return nil
}
