package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Response struct {
	Success bool    `json:"success"`
	Results Results `json:"results"`
}

type Results struct {
	Locations []Location `json:"locations"`
}

type Location struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	Phone        string `json:"phone"`
	Email        string `json:"email"`
	Url          string `json:"url"`
	Address      string `json:"address"`
	AddressLine1 string `json:"address_line_1"`
	AddressLine2 string `json:"address_line_2"`
	Postcode     string `json:"Postcode"`
	State        string `json:"State"`
	City         string `json:"City"`
	Yelp_id      string `json:"yelp_id"`
	Twitter      string `json:"twitter"`
	Instagram    string `json:"instagram"`
}

func main() {

	jsonString, err := os.ReadFile("complex.json")
	if err != nil {
		log.Fatal(err)
	}

	// JSON structure
	fmt.Println("\n --- JSON ---")
	fmt.Println(string(jsonString))

	var Response Response
	err = json.Unmarshal([]byte(jsonString), &Response)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("\n --- Response struct ---")
	fmt.Printf("%+v\n", Response)

}
