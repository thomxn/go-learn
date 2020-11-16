package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	jim := person{
		firstName: "Jimi",
		lastName:  "Hendrix",
		contact: contactInfo{
			email:   "jim@google.com",
			zipCode: 98000,
		},
	}

	jim.print()

	// jimPointer := &jim
	// jimPointer.updateName("Jimmy")

	jim.updateName("Jimmy")
	jim.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}
