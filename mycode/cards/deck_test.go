package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades but got %v", d[0])
	}

	if d[len(d)-1] != "King of Clubs" {
		t.Errorf("Expected first card of King of Clubs but got %v", d[len(d)-1])
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deckForSave := newDeck()
	deckForSave.saveToFile("_decktesting")

	deckLoaded := newDeckFromFile("_decktesting")

	if len(deckLoaded) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(deckLoaded))
	}

	os.Remove("_decktesting")
}
