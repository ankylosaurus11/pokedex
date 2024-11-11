package main

import (
	"fmt"

	"github.com/ankylosaurus11/pokedex/internal/pokecache"
)

func inspect(_ *config, _ *pokecache.Cache, pokemon ...string) error {
	value, ok := caughtPokemon[pokemon[0]]
	if ok {
		fmt.Printf("Name: %v\n", pokemon[0])
		fmt.Printf("Height: %v\n", value.Height)
		fmt.Printf("Weight: %v\n", value.Weight)
		fmt.Println("Stats:")
		for _, stat := range value.Stats {
			fmt.Printf("  -%v: %v\n", stat.Stat.Name, stat.BaseStat)
		}
		fmt.Println("Types:")
		for _, t := range value.Types {
			fmt.Printf("  -%v\n", t.Type.Name)
		}
	} else {
		fmt.Println("You havent caught that Pokemon yet!")
	}
	return nil
}
