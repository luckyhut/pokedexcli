package main

import (
	"fmt"
	"github.com/luckyhut/pokedexcli/internal/pokecache"
)

func commandInspect(c *config, cache *pokecache.Cache, pokemonName *string) error {
	p := c.pokedex[*pokemonName]
	fmt.Println("Name:", p.Name)
	fmt.Println("Height:", p.Height)
	fmt.Println("Weight:", p.Weight)
	fmt.Println("Stats:")
	fmt.Println("  - hp:", p.Stats[0].BaseStat)
	fmt.Println("  - attack::", p.Stats[1].BaseStat)
	fmt.Println("  - defense:", p.Stats[2].BaseStat)
	fmt.Println("  - special attack:", p.Stats[3].BaseStat)
	fmt.Println("  - special defense:", p.Stats[4].BaseStat)
	fmt.Println("  - speed:", p.Stats[5].BaseStat)
	fmt.Println("Types:")
	fmt.Println("  -", p.Types[0].Type.Name)
	if len(p.Types) > 1 {
		fmt.Println("  -", p.Types[1].Type.Name)
	}

	return nil
}
