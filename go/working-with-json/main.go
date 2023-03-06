package main

import (
	"encoding/json"
	"fmt"
)

type Book struct {
	Title  string
	Author string
}

func main() {
	book := Book{Title: "Learning Go", Author: "Josh Lowrance"}

	fmt.Println("book struct:")
	fmt.Printf("%+v\n\n", book)

	// convert book struct into JSON string
	byteArray, err := json.Marshal(book)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("JSON string:")
	fmt.Println(string(byteArray))
}
