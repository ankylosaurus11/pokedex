package main

import (
	"bufio"
	"fmt"
	"os"
)

func pokedexStart() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("pokedex >")
		scanner.Scan()
		userInput := reader.Text()
		if len(userInput) == 0 {
			continue
		}
		command, exists := commands()[userInput]
		if exists {
			err := command.callback()
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

type cliCommand struct {
	name		string
	description string
	callback 	func() error
}

func commands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:		"help",
			description:"provides a list of commands",
			callback: 	commandHelp
		},
		"exit": {
			name:		"exit",
			description:"exits program",
			callback:	commandExit
		}
	}
}
