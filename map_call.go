package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func mapf(cfg *config) error {
	url := ""
	if cfg.NextURL == "" {
		url = "https://pokeapi.co/api/v2/location?limit=20"
	} else {
		url = cfg.NextURL
	}
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
	if err != nil {
		fmt.Println(err)
	}

	for _, location := range apiConfig.Results {
		fmt.Printf("location name: %s\n", location.Name)
	}

	return nil
}

//func mapb() error {

//}
