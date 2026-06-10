package main

import (
	"cbn/greetings"
	"fmt"
	"log"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	names := []string{"Casper", "Malene", "Marty"}

	messages, err := greetings.Hellos(names)


	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(messages)
}
