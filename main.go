package main

import "fmt"

type person struct {
	firstName string
	lastName string
}

func main() {
	alex := person{"Alex", "Anderson"} // one way of creating a person. values are assigned according to order
	// var alex person // creates a var of type person. the fields are assigned zero values ie empty strings here
	// alex.firstName = "Alex"
	// alex.lastName = "Anderson"
	alex = person{firstName: "Alex", lastName: "Anderson"} // better way of creating var's
	fmt.Println(alex)
}