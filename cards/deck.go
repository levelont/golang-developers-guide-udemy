package main

import "fmt"

//go is not object oriented
//in the go approach, we define a new type (instead of having a class)
//and define (attach!) new behaviour

// Create a new type of 'deck'
// which is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	//none of the indexes are needed. underscore them!
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// Receiver of the function
// any variable of type deck gets access to the attached method
// A receiver sets up methods on variables we create
// 	  VVVVVV--- receiver!
func (d deck) Print() {
	// Note that d in this function is a copy of the variable

	// CONVENTION: the name of the receiver is the first letter of the type of the receiver
	// a 1-3 abbreviation is also ok
	for i, card := range d {
		fmt.Println(i, card)
	}
}
