package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"net/http"

	"github.com/ankylosaurus11/pokedex/internal/pokecache"
)

type PokeSuite struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous any    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type PokemonData struct {
	BaseExperience int `json:"base_experience"`
}

func catch(cfg *config, cache *pokecache.Cache, pokemon ...string) error {
	url := "https://pokeapi.co/api/v2/pokemon/?limit=1302"
	fmt.Println(pokemon[0])
	if cacheData, ok := cache.Get(url); ok {
		fmt.Println("Cache hit!")
		var pokeSuite PokeSuite

		err := json.Unmarshal(cacheData, &pokeSuite)
		if err != nil {
			return err
		}

		for _, pokemonCatch := range pokeSuite.Results {
			if pokemonCatch.Name != pokemon[0] {
				continue
			}
			fmt.Println("Pokemon found!")
			res, err := http.Get(pokemonCatch.URL)
			if err != nil {
				return err
			}
			defer res.Body.Close()
			var pokemonData PokemonData
			err = json.NewDecoder(res.Body).Decode(&pokemonData)
			if err != nil {
				return err
			}

			fmt.Println(pokemonData.BaseExperience)

			throwPokeBall := rand.Intn(301)
			fmt.Println(throwPokeBall)
			if throwPokeBall >= pokemonData.BaseExperience {
				fmt.Println("Throwing a Pokeball at " + pokemon[0] + "...")
				fmt.Println(pokemon[0] + " was caught!")
			} else {
				fmt.Println("Throwing a Pokeball at " + pokemon[0] + "...")
				fmt.Println(pokemon[0] + " escaped!")
			}
			return nil
		}
		return errors.New("Pokemon does not exist, check spelling and try again")
	}
	fmt.Println("Cache miss!")
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	var pokeSuite PokeSuite

	err = json.NewDecoder(res.Body).Decode(&pokeSuite)
	if err != nil {
		return err
	}

	for _, pokemonCatch := range pokeSuite.Results {
		if pokemonCatch.Name != pokemon[0] {
			continue
		}
		fmt.Println("Pokemon found!")
	}

	jsonData, err := json.Marshal(pokeSuite)
	if err != nil {
		return err
	}

	cache.Add(url, jsonData)

	return nil
}
