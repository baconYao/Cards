package main

import (
	"fmt"
	"io/ioutil"
	"os"
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

// Save a list of cards to a file on the local machine
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

// Load a list of cards from the local machine
func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		// log the error and return a call to newDeck()
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// convert []bytes to string type, and then split it by using ",". Finally, it turns to be a slice such as ["Ace of Spades", "Two of Spades", ...]
	s := strings.Split(string(bs), ",")
	// convert slice to deck type
	return deck(s)
}
