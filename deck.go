package main

import (
	"fmt"
	"strings"
)

// Create a new type of 'deck'
// which is a slice of strings
type deck []string

// Create and return a list of playing cards. Essentially an array of strings.
func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// Any variable type of "deck" now gets access to the "print" method
// Log out the contents of a deck of cards
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// Create a 'hand' of cards.
func deal(d deck, handSize int) (deck, deck) {
	// return two slices which are type of deck
	// 1: index 0 ~ handSize-1
	// 2: handSize ~ the end of slice
	return d[:handSize], d[handSize:]
}

// convert 'deck' type to 'string' type
func (d deck) toString() string {
	// 將 d (deck type) 轉回 []string type，並用 ","連接slices內的字串
	return strings.Join([]string(d), ",")
}