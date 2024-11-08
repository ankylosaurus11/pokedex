package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/ankylosaurus11/pokedex/internal/pokecache"
)

func mapf(cfg *config, cache *pokecache.Cache, _ ...string) error {
	url := ""

	if cfg.NextURL == "" {
		url = "https://pokeapi.co/api/v2/location-area?limit=20"
	} else {
		url = cfg.NextURL
	}

	if cachedData, ok := cache.Get(url); ok {
		fmt.Println("Cache hit!")
		var apiConfig apiConfig

		err := json.Unmarshal(cachedData, &apiConfig)
		cfg.NextURL = apiConfig.Next
		cfg.PreviousURL = apiConfig.Previous
		if err != nil {
			fmt.Println(err)
		}

		for _, location := range apiConfig.Results {
			fmt.Printf("location name: %s\n", location.Name)
		}

		return nil
	}

	fmt.Println("Cache Miss!")

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("response failed with status code: %d and \nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}
	var apiConfig apiConfig

	err = json.Unmarshal(body, &apiConfig)
	cfg.NextURL = apiConfig.Next
	cfg.PreviousURL = apiConfig.Previous

	cache.Add("locations", body)

	if err != nil {
		fmt.Println(err)
	}

	for _, location := range apiConfig.Results {
		fmt.Printf("location name: %s\n", location.Name)
	}

	return nil
}

func mapb(cfg *config, cache *pokecache.Cache, _ ...string) error {
	url := ""
	if cfg.PreviousURL == nil || *cfg.PreviousURL == "" {
		return errors.New("you are at the start of the list")
	} else {
		url = *cfg.PreviousURL
	}

	if cachedData, ok := cache.Get(url); ok {
		fmt.Println("Cache hit!")
		var apiConfig apiConfig

		err := json.Unmarshal(cachedData, &apiConfig)
		cfg.NextURL = apiConfig.Next
		cfg.PreviousURL = apiConfig.Previous

		if err != nil {
			fmt.Println(err)
		}

		for _, location := range apiConfig.Results {
			fmt.Printf("location name: %s\n", location.Name)
		}

		return nil
	}
	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	var apiConfig apiConfig
	err = json.Unmarshal(body, &apiConfig)
	if err != nil {
		return err
	}
	cfg.NextURL = apiConfig.Next
	cfg.PreviousURL = apiConfig.Previous

	cache.Add("locations", body)

	for _, location := range apiConfig.Results {
		fmt.Printf("location name: %s\n", location.Name)
	}

	return nil
}
