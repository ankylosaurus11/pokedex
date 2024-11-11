package main

import (
	"fmt"

	"github.com/ankylosaurus11/pokedex/internal/pokecache"
)

func pokedex(_ *config, _ *pokecache.Cache, _ ...string) error {
	fmt.Println("Your Pokedex:")
	for key, _ := range caughtPokemon {
		fmt.Println("  - ", key)
	}
	return nil
}
