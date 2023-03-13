package main

import (
	"fmt"
	"strings"
)

type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, suit+" of "+value)
		}
	}

	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) { // (deck, deck) == return two values, both of type deck.
	return d[:handSize], d[handSize:]
}

// see WriteFile in https://pkg.go.dev/io/ioutil
// we need to take our deck type and convert this to a byte slice (slice of bytes).
func (d deck) toString() string {
	// we can do this because deck is of type slice of stirng.
	// converts ["red", "yellow", "blue"] to "red, yellow, blue".
	return strings.Join([]string(d), ",")
}
