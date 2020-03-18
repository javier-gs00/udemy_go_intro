package main

import "fmt"

type person struct {
	firstName string
	lastname  string
}

func main() {
	// Declare a new type
	// alex := person{"Alex", "Anderson"}
	// alex := person{firstName: "Alex", lastname: "Anderson"}
	var alex person // zero value

	// Update properties
	alex.firstName = "Alex"
	alex.lastname = "Anderson"

	fmt.Println(alex)
	// fmt.Printf("%v", alex)
}
