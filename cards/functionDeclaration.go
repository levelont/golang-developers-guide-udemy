package main

import "fmt"

// Explicitly indicate the return type value in a function declaration!
func createNewCard() string {
	return "Five of Diamonds"
}

func createAndPrintNewCard() {
	//Go inferes the type of the variable based on the return type set to the function
	card := createNewCard() //type will be string

	fmt.Println(card)
}
