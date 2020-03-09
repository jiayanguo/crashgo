package main

import (
	"os"
	"testing"
)


// test file naming convention, end with _test.go

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %d", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %s", d[0])
	}

	if d[len(d) - 1] != "Four of Clubs" {
		t.Errorf("Expected last card of Four of Clubs, but got %s", d[len(d) - 1])
	}
}

func TestSaveToDeckAndNewDeckFromFile (t *testing.T) {
	file := "_decktesting"

	os.Remove(file)

	deck := newDeck()
	deck.saveToFile(file)

	loadedDeck := newDeckFromFile(file)

	if len(loadedDeck) != 16 {
		t.Errorf("Expected deck length of 16, but got %d", len(loadedDeck))
	}

	os.Remove(file)
}