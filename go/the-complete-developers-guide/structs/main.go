package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	// basic usage of a struct
	alex := person{"Alex", "Anderson"}
	// if the order of the properties are swapped,  Alex's firstName will be set to his lastName.

	// the verbose approach to using a struct:
	bob := person{firstName: "Bob", lastName: "Bones"}
	charlie := person{lastName: "Cobb", firstName: "Charlie"}
	// when you use this approach you can change the order of the properties without consequence.

	fmt.Println(alex)
	fmt.Println(bob)
	fmt.Println(charlie)

	// there is one more way to declare a struct.
	var josh person
	fmt.Println(josh)
	// returns zero values:
	// {  }

	fmt.Printf("%+v\n", josh)
	// returns properties and zero values:
	// {firstName: lastName:}

	// each empty attribute defaults to a zero value if you do not use it.

	// zero values for respective types:
	// string -> ""
	// int -> 0
	// float -> 0
	// bool -> false

	// you may want to set the default values to something else such as nil.

	// you can update the properties of an object using a struct:
	josh.firstName = "Josh"
	josh.lastName = "Lowrance"
	fmt.Println(josh)
}
