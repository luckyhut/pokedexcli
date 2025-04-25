package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) Explore(pageURL string) (LocationArea, error) {
	req, err := http.NewRequest("GET", pageURL, nil)
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
