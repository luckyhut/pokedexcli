package main

import (
	"github.com/luckyhut/pokedexcli/internal/pokeapi"
	"time"
)

type config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
	pokedex          map[string]pokeapi.Pokemon
}

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	firstLocationURL := "https://pokeapi.co/api/v2/location-area"
	prevLocationURL := "https://pokeapi.co/api/v2/location-area"
	cfg := &config{
		pokeapiClient:    pokeClient,
		nextLocationsURL: &firstLocationURL,
		prevLocationsURL: &prevLocationURL,
	}
	cfg.pokedex = make(map[string]pokeapi.Pokemon)
	startRepl(cfg)
}
