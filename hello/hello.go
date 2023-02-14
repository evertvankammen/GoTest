package main

import (
	"evertlevert.com/greetings"
	"fmt"
)

func main() {
	// Get a greeting message and print it.
	message := greetings.Hello("Gladys")
	fmt.Println(message)
}
