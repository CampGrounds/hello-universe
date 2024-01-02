package main

import (
	"fmt"

	"github.com/CampGrounds/hello-universe/greetings"
)

func main() {
	// Get a greeting message and print it.
	message := greetings.Hello("gladies")
	fmt.Println(message)
}
