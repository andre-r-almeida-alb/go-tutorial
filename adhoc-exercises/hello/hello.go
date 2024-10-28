package main

import (
	"fmt"
	"log"

	"github.com/andre-r-almeida-alb/go-tutorial/adhoc-exercises/greetings"
)

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// Get a greeting message and print it.
	message, err := greetings.Hello("Gladys")
	handleError(err)
	fmt.Println(message)

	// A slice of names.
	names := []string{"Gladys", "Samantha", "Darrin"}
	messages, err := greetings.Hellos(names)
	handleError(err)
	fmt.Println(messages)

	// Get a greeting messager with error
	message, err = greetings.Hello("")
	handleError(err)
	fmt.Println(message)
}

func handleError(err error) {
	// If an error was returned, print it to the console and
	// exit the program.
	if err != nil {
		log.Fatal(err)
	}
}
