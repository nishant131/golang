package main

import (
	"fmt"
	"greetings"
	"log"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	names := []string{"Nishant", "Rakesh", "Utkarsh", ""}
	// message, err := greetings.Hello("Nishant")
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	for k, v := range messages {
		fmt.Println(k, v)
	}
	// fmt.Println(messages)
}
