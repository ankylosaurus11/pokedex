package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ankylosaurus11/pokedex/internal/pokecache"
)

type Pokemon struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type PokemonEncounter struct {
	Pokemon Pokemon `json:"pokemon"`
}

type LocationArea struct {
	PokemonEncounters []PokemonEncounter `json:"pokemon_encounters"`
}

func explore(cfg *config, cache *pokecache.Cache, locationName ...string) error {
	url := "https://pokeapi.co/api/v2/location-area/" + locationName[0]
	if cachedData, ok := cache.Get(url); ok {
		fmt.Println("Cache hit!")
		var locationArea LocationArea

		err := json.Unmarshal(cachedData, &locationArea)
		if err != nil {
			return err
		}
		fmt.Println("Found Pokemon:")
		for _, pokemon := range locationArea.PokemonEncounters {
			fmt.Println(" - ", pokemon.Pokemon.Name)
		}

		return nil
	}
	fmt.Println("Cache Miss!")
	res, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("location area '%s' not found - please try a valid location", locationName[0])
	}

	defer res.Body.Close()
	if res.StatusCode > 299 {
		return fmt.Errorf("response failed with status code: %d -  please try a valid location", res.StatusCode)
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

	jsonData, err := json.Marshal(locationArea)
	if err != nil {
		return err
	}

	cache.Add(url, jsonData)

	return nil
}
