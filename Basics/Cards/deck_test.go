package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	deck := newDeck()

	if len(deck) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(deck))
	}

	if deck[0] != "Ace of Spades" {
		t.Errorf("Expected first card to be Ace of Spades but got %v", deck[0])
	}

	if deck[len(deck)-1] != "King of Hearts" {
		t.Errorf("Expected first card to be King of Clubs but got %v", deck[len(deck)-1])
	}
}

func TestSaveAndReadDeck(t *testing.T) {
	filename := "_decktestfile"
	os.Remove(filename)

	deck := newDeck()
	deck.saveToFile(filename)

	fileDeck := getFromFile(filename)

	if len(fileDeck) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(deck))
	}
	if deck[0] != fileDeck[0] {
		t.Errorf("Mismatch deck")
	}

	os.Remove(filename)

}
