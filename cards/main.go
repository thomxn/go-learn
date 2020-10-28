package main

func main() {
	// var card string = "Ace of spades"

	cards := newDeck()

	// hand, remainingCards := deal(cards, 5)

	// cards.saveToFile("default_cards")

	// fcards := newDeckFromFile("default_cards")
	// fmt.Println(fcards)

	// for i, j := range fcards {
	// 	fmt.Println(i, j)
	// }

	cards.shuffle()
	cards.print()
}
