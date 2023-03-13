package main

func main() {
	cards := newDeckFromFile("my_cards")
	cards.print()

	// if you want to test error handling, attempt to load a file that does not exist.
	newDeckFromFile("asdfasdfasd")
	// Error: open asdfasdfasd: no such file or directory
	// exit status 1
}
