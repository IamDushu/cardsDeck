package main

import "fmt"

// 1. define struct
type person struct {
	firstName string
	lastName  string
}

func main() {
	// 2. Create new value of type person
	// dushu := person{firstName: "Dushyanth", lastName: "Gali"}
	//or
	// dushu := person{"Dushyanth", "Gali"}

	var dushu person

	dushu.firstName = "Dushyanth"
	dushu.lastName = "Gali"

	fmt.Println(dushu)
	fmt.Printf("%+v", dushu)
}
