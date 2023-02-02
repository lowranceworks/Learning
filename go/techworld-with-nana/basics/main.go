package main

import "fmt"

func main() {

	// print to terminal
	println("\n--- print to terminal ---\n")
	fmt.Println("Hello World!")

	// variables
	var conferenceName = "Go Conference" // this value is allowed to change.
	const conferenceTickets = 50         // this value is NOT allowed to change.
	var remainingTickets = conferenceTickets

	println("\n--- use variables with strings ---\n")
	fmt.Println("Welcome to", conferenceName, "booking application")
	fmt.Println("We have a total of", conferenceTickets, "tickets and", remainingTickets, "are still available")
	fmt.Println("Get your tickets here to attend")

	// print formatted data
	println("\n--- print formatted data ---\n")
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available. \n", conferenceTickets, remainingTickets)

	// see documentation on this here: https://pkg.go.dev/fmt

	// .
	println("\n--- type declarations ---\n")
	var userName string // Go will infer data types based on the value provided. In this case, we didn't provide a value so we have to define a data type.
	var userTickets int

	userName = "Tom"
	userTickets = 2
	fmt.Printf("User %v booked %v tickets. \n", userName, userTickets)

	println("\n--- print the types of these variables ---\n")
	fmt.Printf("userName is of type %T and userTickets is of type %T\n", userName, userTickets)

	// syntactic sugar is a term to describe a feature in a language that let's you do smth more easily. Makes the language "sweeter" for human use.
	// This doesn't add any new functionality that it didn't already have.
	// This does not work with constants.
	println("\n--- syntactic sugar. ---\n")
	conferenceName1 := "Go Conference"
	println(conferenceName1)

	// println("\n--- . ---\n")
	// println("\n--- . ---\n")

}
