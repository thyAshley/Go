package main

import "fmt"

type deckType []string

func newDeck() deckType {
	cards := deckType{}
	cardSuites := []string{"Spades", "Hearts", "Diamonds", "Hearts"}
	cardValue := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}
	for _, suit := range cardSuites {
		for _, value := range cardValue {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func (deck deckType) print() {
	for i, card := range deck {
		fmt.Println(i, card)
	}
}

func (deck deckType) shuffle() {
	for i, card := range deck {
		fmt.Println(i, card)
	}
}
