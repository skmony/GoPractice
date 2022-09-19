package main

import "fmt"

type contactInfo struct {
	email   string
	pinCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	//shashikant := person{"Shashikant", "Mony"} way 1

	// shashikant := person{firstName: "Shashikant", lastName: "Mony"} way 2

	var shashikant person

	shashikant.firstName = "Shashikant"
	shashikant.lastName = "Mony"
	shashikant.contactInfo.email = "xyz@hotmail.com"
	shashikant.contactInfo.pinCode = 870001

	shashikant.updateName("John")
	shashikant.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (p *person) updateName(newFirstName string) {
	p.firstName = newFirstName
}
