package main

import (
	"fmt"

	"example.com/greetings"

	"rsc.io/quote"
)

func main() {
	fmt.Printf(quote.Opt())

	message := greetings.Hello("Gladys")
	fmt.Println(message)
}
