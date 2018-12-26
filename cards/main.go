package main

import "fmt"

//Possible to declare variables outside of the function scope without assigning a value to them
var variableWithoutValue string

//Possible to declare variables and assign value to them
var variableWithtValue string = "value"

//not possible to use short hand declaration syntax
//shortHandDeclaredVariable := "shortHandValue"
//produces: syntax error: non-declaration statement outside function body

func main() {
	//Long form declaration
	// var card string = "Ace of Spades"
	//Golang is a statically typed language!

	// Short form declaration with initialization
	card := "Ace of Spades"
	// second time, the variable already exists, no need for the colon
	card /*:*/ = "Five of Diamonds"

	fmt.Println(card)
	fmt.Println(variableWithoutValue)
	fmt.Println(variableWithtValue)
}
