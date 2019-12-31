package main

func main() {
	cards := newDeck()
	cards.saveToFile("my_cards")

	cards2 := newDeckFromFile("my_cards")
	cards2.print()
}
