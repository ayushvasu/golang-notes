package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 20 {
		t.Errorf("Expected deck length of 20 but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of deck to be Ace of Spades but got %v", d[0])
	}

	if d[len(d)-1] != "Five of Club" {
		t.Errorf("Expected last card of deck to be Five of Club but got %v", d[len(d)-1])
	}

}

func TestSaveDeckToFileAndNewDeckTestFromFile(t *testing.T) {
	os.Remove("_decktesting")

	d := newDeck()

	d.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 20 {
		t.Errorf("Expected value is 20 got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")

}
