package main

import (
	"bufio"
	"fmt"
	"os"
)

func pokedexStart() {
	responseCall := bufio.NewReader(os.Stdin)
	fmt.Println("pokedex >")
	sentence, err := responseCall.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	} else if sentence == "help\n" {
		help()
	} else {
		fmt.Println(sentence)
	}
}

type cliCommand struct {
	name		string
	description string
	callback 	func() error
}

func commands() map[string]cliCommand {
	return map[string]cliCommand{
		"help" {
			name:		"help",
			description:"provides a list of commands",
			callback: 	commandHelp
		},
		"exit" {
			name:		"exit",
			description:"exits program",
			callback:	commandExit
		}
	}
}
