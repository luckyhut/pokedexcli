package main

import (
	"github.com/luckyhut/pokedexcli/internal/pokeapi"
	"time"
)

type config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
}

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	cfg := &config{
		pokeapiClient: pokeClient,
	}
	startRepl(cfg)
}
