package main

import (
	"testing"
	"os"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	// Target: Code to make sure that a deck is created with x number of cards
	if len(d) != 16 {
		t.Errorf("Expected deck length of 20, but got: %v", len(d))
	}

	// Target: Code to make sure that the first card is an Ace of Spades
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card is Ace of Spades, but got %v", d[0])
	}

	// Target: Code to make sure that the last card is a Four of Clubs
	if d[len(d) - 1] != "Four of Clubs" {
		t.Errorf("Expected last card is Four of Clubs, but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	// Delete any files in current working directory with the name "_decktesting"
	os.Remove("_decktesting")
	// Create a deck
	deck := newDeck()
	// Save to file "_decktesting"
	deck.saveToFile("_decktesting")
	// Load from file
	loadedDeck := newDeckFromFile("_decktesting")
	// Assert deck length
	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 cards in deck, got %v", len(loadedDeck))
	}
	// Delete any files in current working directory with the name "_decktesting"
	os.Remove("_decktesting")
}