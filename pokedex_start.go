package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/ankylosaurus11/pokedex/internal/pokecache"
)

func pokedexStart(cfg config, cache *pokecache.Cache) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("pokedex >")
		scanner.Scan()
		userInput := scanner.Text()
		if len(userInput) == 0 {
			continue
		}
		command, exists := commands()[userInput]
		if exists {
			err := command.callback(&cfg, cache)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("not a valid command, please try 'help'")
			continue
		}
	}
}

type config struct {
	NextURL     string
	PreviousURL *string
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, *pokecache.Cache) error
}

type apiConfig struct {
	Count    int        `json:"count"`
	Next     string     `json:"next"`
	Previous *string    `json:"previous"`
	Results  []location `json:"results"`
}

type location struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func commands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "provides a list of commands",
			callback:    help,
		},
		"exit": {
			name:        "exit",
			description: "exits program",
			callback:    exit,
		},
		"map": {
			name:        "map",
			description: "shows first 20 locations",
			callback:    mapf,
		},
		"mapb": {
			name:        "mapb",
			description: "shows previous 20 locations",
			callback:    mapb,
		},
	}
}
