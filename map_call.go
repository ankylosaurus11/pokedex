package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func mapf() error {
	res, err := http.Get("https://pokeapi.co/api/v2/location?limit=20")
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

	if err != nil {
		fmt.Println(err)
	}

	for _, location := range apiConfig.Results {
		fmt.Printf("name: %s\n", location.Name)
	}

	return nil
}

//func mapb() error {

//}
