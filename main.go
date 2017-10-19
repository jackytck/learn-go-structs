package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	// 1. implicit order
	alex := person{"Alex", "Anderson"}
	fmt.Printf("%+v\n", alex)

	// 2. field names
	alex2 := person{firstName: "Alex", lastName: "Anderson"}
	fmt.Printf("%+v\n", alex2)

	// 3. zero value
	var alex3 person
	alex3.firstName = "Alex"
	alex3.lastName = "Anderson"
	fmt.Printf("%+v\n", alex3)
}
