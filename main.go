package main

import (
	"fmt"
	"log"

	"./greetings"
)

func main() {
	log.SetPrefix("greetings :")
	log.SetFlags(0)

	message, err := greetings.Hello("Matthieu")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}

func addOne(value int) int {
	return value + 1
}
