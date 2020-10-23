package main

import "fmt"

func main() {
	// var card string = "Ace of spades"

	cards := newDeck()

	hand, remainingCards := deal(cards, 5)

	fmt.Println(hand.toString())
	remainingCards.print()
}
