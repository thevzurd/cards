package main

import (
	"os"
	"testing"
)

// newDeck
// Check the length of the deck
// Check the first element and the last element

// Test<function name> test that function
// The function name is Uppercase unlike for normal func
// t (test handler) passed to the fuction
// T denotes the type of value
func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 12 {
		t.Errorf("Expected deck len 12 but got %v", len(d))
	}

	if d[0] != "A of S" {
		t.Errorf("Expected A of S but got %v", d[0])
	}
	if d[len(d)-1] != "3 of C" {
		t.Errorf("Expected 3 of C but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")
	deck := newDeck()
	deck.saveToFile("_decktesting")
	d := newDeckFromFile("_decktesting")
	if len(d) != 12 {
		t.Errorf("Expected deck len 12 but got %v", len(d))
	}

	if d[0] != "A of S" {
		t.Errorf("Expected A of S but got %v", d[0])
	}
	if d[len(d)-1] != "3 of C" {
		t.Errorf("Expected 3 of C but got %v", d[len(d)-1])
	}
}
