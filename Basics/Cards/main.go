package main

func main() {
	cards := newDeck()

	hand, remainder := deal(cards, 5)
	hand.print()
	remainder.print()
	remainder.saveToFile("storeDeck")

}

func newCard() string {
	return "Five of Diamonds"
}
