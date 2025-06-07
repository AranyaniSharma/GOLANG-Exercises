package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {

	d := newDeck()

	if len(d) != 52 {
		t.Errorf("expected Deck length 52 but got, %v", len(d))
	}
	if d[0] != "Ace of Diamond" {
		t.Errorf("Expected first card of Ace Of Diamond but got , %v", d[0])
	}
	if d[len(d)-1] != "King of clubs" {
		t.Errorf("Expected first card of King Of club but got , %v", d[len(d)-1])
	}

}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 52 {
		t.Errorf("expected Deck length 52 but got, %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
