package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	
	if len(d) != 16 {
		t.Errorf("Expected deck length of 20, but got %d", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected Ace of Spades but got %v", d[0])
	}

	if(d[len(d)-1] != "Four of Club"){
		t.Errorf("Expected Four of Club but got %v", d[len(d)-1])
	}

}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T){
	os.Remove("_decktesting.txt")

	deck := newDeck()
	deck.saveToFile("_decktesting.txt")
	loadedDeck := newDeckFromFile("_decktesting.txt")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 cards but got %v", len(loadedDeck))
	}
	os.Remove("_decktesting.txt")
}