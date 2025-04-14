package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

var commands = map[string]cliCommand{}

func main() {
	initCommands()
	s := bufio.NewScanner(os.Stdin)
	//fmt.Println(s)
	for {
		fmt.Print("Pokedex > ")
		s.Scan()
		input := s.Text()
		words := strings.Fields(input)

		// check command
		c := commands[words[0]]
		c.callback()
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
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")

	for k, v := range commands {
		fmt.Printf("%s: %s\n", k, v.description)
	}
	return nil
}
