package main

import (
	"fmt"

	"github.com/ankylosaurus11/pokedex/internal/pokecache"
)

func help(_ *config, _ *pokecache.Cache) error {
	fmt.Println("helpmenu:")
	for _, cmd := range commands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}
