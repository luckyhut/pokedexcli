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
	callback    func(*config, *pokecache.Cache) error
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

		// check command
		c := commands[words[0]]
		if c.callback == nil {
			fmt.Println("Unknown command")
			continue
		}
		c.callback(cfg, &cache)
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
}

func cleanInput(text string) []string {
	result := []string{}
	words := strings.Fields(text)
	for _, word := range words {
		result = append(result, strings.ToLower(word))
	}

	return result
}
