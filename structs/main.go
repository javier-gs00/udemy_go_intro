package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastname  string
	// contact   contactInfo
	contactInfo
}

func main() {
	jim := person{
		firstName: "Jim",
		lastname:  "Party",
		// contact: contactInfo{
		// 	email:   "jim@gmail.com",
		// 	zipCode: 9400,
		// },
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 9400,
		},
	}

	jim.updateName("Jimmy")

	jim.print()
}

func (p person) updateName(newFirstName string) {
	p.firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
