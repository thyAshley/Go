package main

func main() {
	cards := newDeck()
	cards.shuffle()
	hand, remainder := deal(cards, 5)
	hand.print()
	remainder.print()
}

func newCard() string {
	return "Five of Diamonds"
}
