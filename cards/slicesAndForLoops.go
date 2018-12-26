package main

import "fmt"

func main() {
	//arrays are basic collections
	//slices are more sofisticated

	//Arrays have a fixed length, slices can vary in size (enlarge, shrink)

	// Both arrays and slices need to have all of its elements to be of the same type

	// square brackets + type = collection
	cards := []string{"Ace of Diamonds", newCard(), newCard()}
	// read aloud: slice of type string

	//how to add a new element?
	cards = append(cards, "Six of Spades")
	//append does not modify the existing slice
	//instead, it returns a new one!

	//how to iterate over all elements?

	// Note that with every iteration, a new set of (i, card)
	// variables will be created, the old ones will be thrown away!
	for i, card := range cards { //take the slice and iterate over every element
		//the return type is index, value
		fmt.Println(i, card)
	}
}
