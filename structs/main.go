package main

import (
	"fmt"
)

type contactInfo struct {
	email string
	zipCode int
}

type person struct {
	firstName string
	lastName string
	//contact contactInfo
	contactInfo
}

func main() {
	// alex := person{firstName :"Alex",  lastName : "Anderson"}

	var alex person
	alex.lastName = "Alex"
	alex.firstName = "Anderson"
	alex.contactInfo = contactInfo{
		email: "alex@gmail.com",
		zipCode: 10000,
	}


	// pass by value language

	// alexPointer := &alex  // & memory address
	alex.updateName("Andy")  // alex auto converted to pointer
	alex.print()
}

func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName  // * the value of this address
}

func (p person) print() {
	fmt.Printf("%+v \n", p)
}