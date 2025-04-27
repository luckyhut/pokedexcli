package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemonInfo(pokemonName *string) Pokemon {
	fullURL := BaseURL + "pokemon/" + *pokemonName
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		fmt.Println("Error getting Pokemon data")
		return Pokemon{}
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		fmt.Println("Error getting Pokemon data")
		return Pokemon{}
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error getting Pokemon data")
		return Pokemon{}
	}

	pokemon, err := UnmarshalPokemon(data)
	if err != nil {
		fmt.Println("Error getting Pokemon data")
		return Pokemon{}
	}
	return pokemon
}

func UnmarshalPokemon(data []byte) (Pokemon, error) {
	pokemon := Pokemon{}
	err := json.Unmarshal(data, &pokemon)
	if err != nil {
		fmt.Println("error unmarshaling pokemon json data")
		return Pokemon{}, err
	}
	return pokemon, err
}
