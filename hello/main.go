package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	message := greetings.Hello("666")
	fmt.Println(message)
}
