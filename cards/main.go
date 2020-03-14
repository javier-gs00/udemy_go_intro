package main

import "fmt"

func main() {
	cards := newDeck()

	hand, remainingDeck := deal(cards, 5)

	cards.print()
	fmt.Println("\n")
	hand.print()
	fmt.Println("\n")
	remainingDeck.print()
}
