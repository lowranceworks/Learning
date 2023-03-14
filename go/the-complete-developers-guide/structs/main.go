package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Jones",
		contact: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 78750, // multiline structs must have a comma at the end of the property, even if it is the last property.
		},
	}

	fmt.Printf("%+v", jim) // {firstName:Jim lastName:Jones contact:{email:jim@gmail.com zipCode:78750}}
}
