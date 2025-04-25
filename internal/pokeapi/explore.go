package pokeapi

import (
	"encoding/json"
	"fmt"
	"github.com/luckyhut/pokedexcli/internal/pokecache"
	"io"
	"net/http"
)

func (c *Client) Explore(pageURL *string, cache *pokecache.Cache) (LocationArea, error) {
	req, err := http.NewRequest("GET", *pageURL, nil)
	if err != nil {
		return LocationArea{}, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationArea{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationArea{}, err
	}

	cache.Add(*pageURL, data)

	encountersResp, err := UnmarshalPokemonEncounter(data)
	if err != nil {
		return LocationArea{}, err
	}
	return encountersResp, err
}

func UnmarshalPokemonEncounter(data []byte) (LocationArea, error) {
	encountersResp := LocationArea{}
	err := json.Unmarshal(data, &encountersResp)
	if err != nil {
		fmt.Println("there was an error")
		return LocationArea{}, err
	}
	return encountersResp, err
}
