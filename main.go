package main

import (
	"fmt"
	"log"

	"github.com/rawkalnapus/greetings"
)

func main() {
	log.SetPrefix("Greetings: ")
	log.SetFlags(0)

	message, err := greetings.Hello("Steve")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}
