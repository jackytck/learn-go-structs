package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

// embeded struct
type person struct {
	firstName string
	lastName  string
	// contact   contactInfo
	contactInfo
}

// p is pass by value
func (p person) print() {
	fmt.Printf("%+v\n", p)
}

func (p *person) updateName(newFirstName string) {
	p.firstName = newFirstName
}

func main() {
	// // 1. implicit order
	// alex := person{"Alex", "Anderson"}
	// fmt.Printf("%+v\n", alex)

	// // 2. field names
	// alex2 := person{firstName: "Alex", lastName: "Anderson"}
	// fmt.Printf("%+v\n", alex2)

	// // 3. zero value
	// var alex3 person
	// alex3.firstName = "Alex"
	// alex3.lastName = "Anderson"
	// fmt.Printf("%+v\n", alex3)

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}

	jim.updateName("jimmy")
	jim.print()

	mapBasics()

	inter()

	getRequest()
}
