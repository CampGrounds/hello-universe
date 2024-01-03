package main

import (
	"fmt"

	"github.com/CampGrounds/hello-universe/greetings"
)

func main() {
	message := greetings.Hello("Gladius")
	fmt.Println(message)
}
