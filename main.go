package main

func main() {
	cards := deck{newCard(), "Ace of Diamonds"}
	cards = append(cards, "six of spades")

	cards.print()

}

func newCard() string {
	return "Ace of Spades"
}
