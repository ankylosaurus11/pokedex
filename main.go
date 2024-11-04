package main

import (
	"time"

	"github.com/ankylosaurus11/pokedex/internal/pokecache"
)

func main() {
	cache := pokecache.NewCache(5 * time.Minute)
	cfg := config{
		NextURL:     "",
		PreviousURL: nil,
	}
	pokedexStart(cfg, &cache)
}
