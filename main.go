package main

import "fmt"

type contactInfo struct {
	email string
	zipCode int
}

type person struct {
	firstName string
	lastName string
	contactInfo
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (p person) greet() {
	fmt.Println("Hey " + p.firstName)
}

func (p person) updateName(newFirstName string) person {
	p.firstName = newFirstName
	return p
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName: "Party",
		contactInfo: contactInfo{
			email: "jim@gmail.com",
			zipCode: 94993,
		},
	}
	// jim.print()
	jim = jim.updateName("Chris")
	jim.greet()
}