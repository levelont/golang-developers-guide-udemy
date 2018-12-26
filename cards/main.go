package main

import "fmt"

func main() {
	//Long form declaration
	// var card string = "Ace of Spades"
	//Golang is a statically typed language!

	// Short form declaration with initialization
	card := "Ace of Spades"
	// second time, the variable already exists, no need for the colon
	card /*:*/ = "Five of Diamonds"

	fmt.Println(card)
}
