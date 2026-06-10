package main

import (
	"cbn/greetings"
	"fmt"
	"log"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	message, err := greetings.Hello("Casper")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}
