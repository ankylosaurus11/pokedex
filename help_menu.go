package main

import (
	"fmt"
)

func help() error {
	fmt.Println("helpmenu:")
	for _, cmd := range commands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}
