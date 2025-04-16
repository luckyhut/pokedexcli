package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	next_url     string
	previous_url string
}

type locationArea struct {
	Count    int
	Next     string
	Previous string
	Results  []struct {
		Name string
		URL  string
	}
}

var commands = map[string]cliCommand{}

func main() {
	initCommands()
	cfg := config{next_url: "https://pokeapi.co/api/v2/location-area/?limit=%d&offset=%d", previous_url: ""}
	s := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		s.Scan()
		input := s.Text()
		words := strings.Fields(input)

		// check command
		c := commands[words[0]]
		c.callback(&cfg)
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

func commandExit(c *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(c *config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")

	for k, v := range commands {
		fmt.Printf("%s: %s\n", k, v.description)
	}
	return nil
}

func commandMap(c *config) error {
	l := new(locationArea)
	resp, err := http.Get(c.next_url)
	if err != nil {
		return err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if resp.StatusCode > 299 {
		return err
	}
	err = json.Unmarshal(body, &l)
	if err != nil {
		return err
	}
	printLocationArea(l)
	c.next_url = l.Next
	c.previous_url = l.Previous
	return nil
}

func commandMapb(c *config) error {
	l := new(locationArea)
	resp, err := http.Get(c.previous_url)
	if err != nil {
		return err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if resp.StatusCode > 299 {
		return err
	}
	err = json.Unmarshal(body, &l)
	if err != nil {
		return err
	}
	printLocationArea(l)
	c.next_url = l.Next
	c.previous_url = l.Previous
	return nil
}

func printLocationArea(l *locationArea) {
	for _, location := range l.Results {
		fmt.Println(location.Name)
	}
}
