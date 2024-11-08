package main

import (
	"github.com/ankylosaurus11/pokedex/internal/pokecache"

	"os"
)

func exit(_ *config, _ *pokecache.Cache, _ ...string) error {
	os.Exit(0)
	return nil
}
