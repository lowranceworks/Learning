package main

import (
	"fmt"
	"io/ioutil" // this is a sub-package inside the io package.
	"math/rand"
	"os"
	"strings"
)

type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
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

// we want to run deck.saveToFile so we will add a receiver of type deck.
// the WriteFile() function requires a string for the fileName.
// we should return the error that might be produced, error is a type.
// anyone can read or write to this file (0666).
// return the whole expression.
func (d deck) saveToFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
}

// this does not need to be a receiver function.
// initialize and define two new variables: bs (byteSlice) or error.
// nil is equivalent to null, it is a value with no value.

func newDeckFromFile(fileName string) deck {
	bs, err := ioutil.ReadFile(fileName)

	// we use conditional logic to handle the possibility of an error.
	// if something goes wrong here, what do you want to do here?..
	// 1. log the error and return a call to neckDeck().
	// 2. log the error and entirely quit the program.
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// turn the byteSlice into a deck
	// []byte -> string -> []string -> deck
	// use the Split() function: https://pkg.go.dev/strings#Split
	s := strings.Split(string(bs), ",")

	// now that is a slice of type string, you can easily convert this into a deck.
	// see line 10, deck is a slice of type string with additional functions.
	// hover over line 77 to see the functions available, these are the receiver functions created from earlier.
	return deck(s)
}

// the go standard library does not have a function to shuffle a slice.
// this should be a receiver function so we can call deck.shuffle.
// nothing is returned from this function.
func (d deck) shuffle() {
	for i := range d {
		newPosition := rand.Intn(len(d) - 1)        // new position is a random int that is in the range of the size of the deck.
		d[i], d[newPosition] = d[newPosition], d[i] // swaps current position with new position.
	}
}
