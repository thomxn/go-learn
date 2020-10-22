package main

func main() {
	// var card string = "Ace of spades"

	cards := deck{newCard(), "Ace of Spades"}

	cards = append(cards, "Ace of Diamonds")

	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
