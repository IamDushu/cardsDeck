package main

import "fmt"

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

// The below 2 functions have same logic :- Lets make it DRY so it works with both bots - INTERFACES!
func printGreeting(eb englishBot) {
	fmt.Println(eb.getGreeting())
}

func printGreeting(tb teluguBot) {
	fmt.Println(tb.getGreeting())
}
