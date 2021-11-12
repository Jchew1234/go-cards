package main

import "fmt"

func main() {
	// cards := deck{"Ace of Diamonds", newCard()}
	// cards = append(cards, "Six of Spades")
	cards := newDeck()

	hand, remainingDeck := deal(cards, 5)
	fmt.Println("deck in hand:")
	hand.print()
	fmt.Println("remaining deck:")
	remainingDeck.print()
	// cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
