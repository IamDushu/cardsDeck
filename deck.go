package main

import "fmt"

// Create a new type of 'deck'
// which is a slice of string
type deck []string

// Func to create new Deck
func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// reciever
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// return multiple values
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}
