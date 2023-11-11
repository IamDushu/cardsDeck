package main

import "fmt"

// 1. define struct
type person struct {
	firstName   string
	lastName    string
	contactInfo //-----> same as contactInfo contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 517501,
		},
	}

	jim.updateName("Jimmy")
	jim.print() //---> still show jim as firstname, why?

}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (p person) updateName(newFirstName string) {
	p.firstName = newFirstName
}
