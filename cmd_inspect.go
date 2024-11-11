package main

import (
	"fmt"

	"github.com/ankylosaurus11/pokedex/internal/pokecache"
)

func inspect(_ *config, _ *pokecache.Cache, pokemon ...string) error {
	value, ok := caughtPokemon[pokemon[0]]
	if ok {
		fmt.Println(value.Stats[0])
		fmt.Println(value.Stats[1])
		fmt.Println(value.Stats[2])
	} else {
		fmt.Println("You havent caught that Pokemon yet!")
	}
	return nil
}
