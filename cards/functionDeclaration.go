package main

import "fmt"

// Explicitly indicate the return type value in a function declaration!
func newCard() string {
	return "Five of Diamons"
}

func createNewCard() {
	//Go inferes the type of the variable based on the return type set to the function
	card := newCard() //type will be string

	fmt.Println(card)
}
