package main

// comment
import (
	"fmt"
)

// you have to have a main function
func main() {
	printHello()
	variables()
}

func printHello() {
	fmt.Println("Hello world")
}

func variables() {

	// two ways to declare variables:
	var person = "Josh" // does not require a value, can be used inside and outside of functions
	number := 3         // infers the value type, can only be used inside functions

	var child string // you can assign a value after it is declared. Helpful if not initially known.
	child = "Lily"

	fmt.Println("Hello")
	fmt.Println(person)
	fmt.Println(number)
	fmt.Println(child)

	var a string // is ""
	var b int    // is 0
	var c bool   // is false

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}

/*
multiline
comment
*/
