package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

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

func deal(deck deckType, handSize int) (deckType, deckType) {
	return deck[:handSize], deck[handSize:]
}

func (deck deckType) toString() string {
	return strings.Join(deck, ",")
}

func (deck deckType) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(deck.toString()), 0666)
}
