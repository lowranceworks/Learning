package main

import (
	"encoding/json"
	"fmt"
)

type Book struct {
	Title  string `json:"title"`
	Author Author `json:"author"` // "Author string" will no longer work. Declare that you are using the "Author" struct.
}

type Author struct {
	Name      string `json:"name"`
	Age       int    `json:"age"`
	Developer bool   `json:"is_developer"`
}

func main() {
	author := Author{Name: "Josh Lowrance", Age: 28, Developer: true}
	book := Book{Title: "Learning Go", Author: author}

	fmt.Println("book struct:")
	fmt.Printf("%+v\n\n", book) // {Title:Learning Go Author:{Name:Josh Lowrance Age:28 Developer:true}}

	// convert book struct into JSON string
	byteArray, err := json.MarshalIndent(book, "", "    ")
	/*
		JSON string:
		{
			"title": "Learning Go",
			"author": {
				"name": "Josh Lowrance",
				"age": 28,
				"is_developer": true
			}
		}
	*/

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("JSON string:")
	fmt.Println(string(byteArray)) // {"title":"Learning Go","author":{"name":"Josh Lowrance","age":28,"is_developer":true}}
}
