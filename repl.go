package main

import (
	"bufio"
	"fmt"
	"github.com/luckyhut/pokedexcli/internal/pokecache"
	"os"
	"strings"
	"time"
)

var commands = map[string]cliCommand{}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, *pokecache.Cache, *string) error
}

func startRepl(cfg *config) {
	initCommands()
	cache := pokecache.NewCache(5 * time.Second)
	s := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		s.Scan()
		input := s.Text()
		words := strings.Fields(input)
		location := ""
		pokemon := ""

		// check command
		if len(words) == 0 {
			continue
		}
		c := commands[words[0]]
		if c.callback == nil {
			fmt.Println("Unknown command")
			continue
		}
		if words[0] == "explore" {
			location = words[1]
			c.callback(cfg, &cache, &location)
			continue
		}
		if words[0] == "catch" {
			pokemon = words[1]
			fmt.Printf("Throwing a Pokeball at %s...\n", pokemon)
			c.callback(cfg, &cache, &pokemon)
			continue
		}
		if words[0] == "inspect" {
			pokemon = words[1]
			c.callback(cfg, &cache, &pokemon)
			continue
		}

		c.callback(cfg, &cache, &location)
	}
}

func initCommands() {
	commands["exit"] = cliCommand{
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commandExit,
	}
	commands["help"] = cliCommand{
		name:        "help",
		description: "Displays a help message",
		callback:    commandHelp,
	}
	commands["map"] = cliCommand{
		name:        "map",
		description: "Displays map locations",
		callback:    commandMap,
	}
	commands["mapb"] = cliCommand{
		name:        "mapb",
		description: "Displays previous map locations",
		callback:    commandMapb,
	}
	commands["explore"] = cliCommand{
		name:        "explore",
		description: "Shows available Pokemon in a given area",
		callback:    commandExplore,
	}
	commands["catch"] = cliCommand{
		name:        "catch",
		description: "Attempt to catch a Pokemon",
		callback:    commandCatch,
	}
	commands["inspect"] = cliCommand{
		name:        "inspect",
		description: "View details for a captured Pokemon",
		callback:    commandInspect,
	}
}

func cleanInput(text string) []string {
	result := []string{}
	words := strings.Fields(text)
	for _, word := range words {
		result = append(result, strings.ToLower(word))
	}

	return result
}
