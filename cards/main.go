package main

func newCard() string {
	return "Five of Diamonds"
}

func main() {
	cards := deck{"Ace of Diamonds", newCard()}

	cards = append(cards, "Six of Spades")
	cards.Print()
}
