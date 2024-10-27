package main

import (
	"bufio"
	"fmt"
	"os"
)

func help() {
	responseCall := bufio.NewReader(os.Stdin)
	fmt.Println("Help Menu:\nhelp - brings up help menu\nexit - exits program")
	sentence, err := responseCall.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	} else if sentence == "help\n" {
		help()
	} else if sentence == "exit\n" {
		fmt.Println("")
	}
}

func main() {
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
