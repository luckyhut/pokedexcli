package pokeapi

import (
	"encoding/json"
	"fmt"
	"github.com/luckyhut/pokedexcli/internal/pokecache"
	"io"
	"net/http"
)

func (c *Client) Locations(pageURL *string, cache *pokecache.Cache) (Locations, error) {
	fmt.Println(*pageURL)
	req, err := http.NewRequest("GET", *pageURL, nil)
	if err != nil {
		return Locations{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Locations{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return Locations{}, err
	}

	cache.Add(*pageURL, data)

	locationsResp, err := UnmarshalLocations(data)
	if err != nil {
		return Locations{}, err
	}
	return locationsResp, err
}

func UnmarshalLocations(data []byte) (Locations, error) {
	locationsResp := Locations{}
	err := json.Unmarshal(data, &locationsResp)
	if err != nil {
		return Locations{}, err
	}
	return locationsResp, err
}
