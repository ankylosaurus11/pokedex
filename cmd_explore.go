package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/ankylosaurus11/pokedex/internal/pokecache"
)

type Pokemon struct {
	Name string `json:"name"`
}

type PokemonEncounter struct {
	Pokemon Pokemon `json:"pokemon"`
}

type LocationArea struct {
	PokemonEncounters []PokemonEncounter `json:"pokemon_encounters"`
}

func explore(cfg *config, cache *pokecache.Cache, locationName ...string) error {
	url := "https://pokeapi.co/api/v2/location-area/" + locationName[0]
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("response failed with status code: %d", res.StatusCode)
	}
	if err != nil {
		log.Fatal(err)
	}
	var locationArea LocationArea

	err = json.NewDecoder(res.Body).Decode(&locationArea)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Found Pokemon:")
	for _, pokemon := range locationArea.PokemonEncounters {
		fmt.Println(" - ", pokemon.Pokemon.Name)
	}

	return nil
}
