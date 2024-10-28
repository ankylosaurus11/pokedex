package main

import (
	"bufio"
	"fmt"
	"os"
)

func pokedexStart() {
	responseCall := bufio.NewReader(os.Stdin)
	fmt.Println("pokedex >")
	sentence, err := responseCall.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	} else if sentence == "help\n" {
		help()
	} else {
		fmt.Println(sentence)
	}
}
