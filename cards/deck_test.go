package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	const deckSize = 40
	deck := newDeck()

	if len(deck) != deckSize {
		t.Errorf("Expected deck length of %d, but got %d", deckSize, len(deck))
	}

	if deck[0] != "Ace of Spades" {
		t.Errorf("Expected Ace of Spades, but got %v", deck[0])
	}

	if deck[len(deck)-1] != "Ten of Clubs" {
		t.Errorf("Expected Ten of Clubs, but got %v", deck[len(deck)-1])
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	const fileName = "_decktesting"
	os.Remove(fileName)

	deck := newDeck()
	deck.saveToFile(fileName)

	loadedDeck := newDeckFromFile(fileName)

	const deckSize = 40

	if len(loadedDeck) != deckSize {
		t.Errorf("Expected deck length of %d, but got %d", deckSize, len(loadedDeck))
	}

	os.Remove(fileName)
}
