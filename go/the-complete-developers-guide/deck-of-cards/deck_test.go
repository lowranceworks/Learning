package main

import (
	"os"
	"testing"
)

// a test is really nothing more than a go function that is executed when you run `go test` in your terminal.
// run `go test` in your terminal to run all files ending in `_test.go`.

// how do we know what to test?
// well as the developer, what do you care about?
// when a deck is created, it always has 16 cards in it.
// one test we can have is checking the length of the deck and making sure it is 16.
func TestNewDeck(t *testing.T) { // tests will always take an argument of (t *testing.T)
	d := newDeck()

	// test 1
	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

	// test 2
	if d[0] != "Ace of Spades" { // if the first card in the deck is not "Ace of Spades".
		t.Errorf("Expected first card of Ace of Spaces, but got %v", d[0])
	}

	// test 3
	if d[len(d)-1] != "Four of Clubs" { // if the last card in the deck is not "Four of Clubs".
		t.Errorf("Expected last card of Four of Clubs, but got %v", d[len(d)-1])
	}
}

// the reason for the long test name:
// if you make a change to the newDeckFromFile() function inside deck.go,
// you can copy/paste the name and find it in `deck_test.go`.
func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting") // delete any files in the current working directory with the name "_decktesting"

	deck := newDeck()               // create a deck
	deck.saveToFile("_decktesting") // save to file "_decktesting"

	loadedDeck := newDeckFromFile("_decktesting") // load from file

	if len(loadedDeck) != 16 { // assert deck length
		t.Errorf("Expected 16 cards in deck, got %v", len(loadedDeck))
	}
}
