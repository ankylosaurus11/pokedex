package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func mapf() error {
	res, err := http.Get("https://pokeapi.co/api/v2/location?limit=5")
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
	var result interface{}
	err = json.Unmarshal(body, &result)
	/*if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		log.Fatalf("error decoding json: %v", err)
	}*/
	if err != nil {
		fmt.Println(err)
	}

	/*data, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		log.Fatalf("error marshalling json: %v", err)
	}*/

	fmt.Println(result)

	return nil
}

//func mapb() error {

//}
