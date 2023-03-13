package main

import "testing"

// how do we know what to test?
// well as the developer, what do you care about?
// when a deck is created, it always has 16 cards in it.
// one test we can have is checking the length of the deck and making sure it is 16.
func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}
}

// run `go test` in your terminal to run all files ending in `_test.go`.
