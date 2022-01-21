package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 28 {
		t.Errorf("Expected deck length of 28, but got %v", len(d))
	}

	if d[0] != "7 of Spades" {
		t.Errorf("Expected first card to be 7 of Spades but got %v", d[0])
	}

	if d[len(d)-1] != "K of Clubs" {
		t.Errorf("Expected last card to be K of Clubs but got %v", d[len(d)-1])
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	s := "_decktesting"
	os.Remove(s)

	d := newDeck()
	d.saveToFile(s)

	loadedDeck := newDeckFromFile(s)

	if len(loadedDeck) != 28 {
		t.Errorf("Expected deck length of 28, but got %v", len(d))
	}

	os.Remove(s)
}