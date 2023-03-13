package main

func main() {
	cards := newDeck()
	hand, remainingCards := deal(cards, 5) // the second return value will be assigned to remainingDeck.

	// we are able to call a print() function because they're both of type deck.
	// see deck.go line 21.
	hand.print()
	remainingCards.print()
}
