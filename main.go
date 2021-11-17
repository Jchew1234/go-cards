package main

func main() {
	// cards := deck{"Ace of Diamonds", newCard()}
	// cards = append(cards, "Six of Spades")

	// cards := newDeckFromFile("my_cards")
	// cards.print()

	cards := newDeck()
	cards.shuffle()
	cards.print()

	// cards := newDeck()
	// cards.saveToFile("my_cards")
	// fmt.Println(cards.toString())

	// hand, remainingDeck := deal(cards, 5)
	// fmt.Println("deck in hand:")
	// hand.print()
	// fmt.Println("remaining deck:")
	// remainingDeck.print()
	// // cards.print()

	// type conversion example of string to byte slice.
	// greeting := "Hi there!"
	// fmt.Println([]byte(greeting))
}

// func newCard() string {
// 	return "Five of Diamonds"
// }
