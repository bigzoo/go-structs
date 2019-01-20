package main

import "fmt"

type person struct {
	firstName string
	lastName string
}

func main() {
	alex := person{"Alex", "Anderson"} // one way of creating a person. values are assigned according to order
	alex = person{firstName: "Alex", lastName: "Anderson"} // better way of creating var's
	fmt.Println(alex)
}