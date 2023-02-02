package main

import "fmt"

func main() {
	var firstName string
	var lastName string
	var emailAddress string

	// ask user for their name (this fails because we are NOT using a pointer)
	println("\n--- ask user for their name (this fails because we are NOT using a pointer) ---\n")
	fmt.Scan(firstName)
	fmt.Printf("user %v booked a ticket\n", firstName)

	// what is a pointer?
	// a variable that points to the memory address of another variable.
	// this is a concept in C or C++ but not a concept in other languages such as javascript.

	// ask user for their name (this works because we are now using a pointer)
	fmt.Println("\n--- ask user for their name (this works because we are now using a pointer) ---\n")

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address: ")
	fmt.Scan(&emailAddress)

	fmt.Printf("Thank you %v %v, a confirmation email has been sent to %v\n", firstName, lastName, emailAddress)

	println("\n--- . ---\n")
	// println("\n--- . ---\n")
	// println("\n--- . ---\n")
	// println("\n--- . ---\n")
	// println("\n--- . ---\n")
	// println("\n--- . ---\n")
	// println("\n--- . ---\n")
	// println("\n--- . ---\n")
	// println("\n--- . ---\n")
	// println("\n--- . ---\n")
	// println("\n--- . ---\n")

}
