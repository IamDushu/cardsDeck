package main

import "fmt"

// Any Type which implements this getGreeting() string; is now an honary member of this interface type bot
type bot interface {
	getGreeting() string
}

type englishBot struct{}
type teluguBot struct{}

func main() {

	eb := englishBot{}
	tb := teluguBot{}

	printGreeting(eb)
	printGreeting(tb)

}

func (eb englishBot) getGreeting() string {
	//VERY Custom logic for generating an english greeting
	return "Hello!"
}

func (teluguBot) getGreeting() string { //-----> if we are not using the reciever value; here tb then we can remove it
	//VERY Custom logic
	return "Namaskaram _/\\_"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}
