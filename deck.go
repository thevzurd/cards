package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// GO is not OO languange.
// Instead we are going to use type and reciever functions like so

// Create a new type of 'deck which is a slice of string
type deck []string

// Create a reciever function for deck, making it act similar to methods in a class
// Any variable of type deck gets access
// to function print
// By convention we name the reciever with the one letter
// from the type. So deck's reciever is d
func (d deck) print() {
	// := is used because each iteration both i and card
	// are recreated with the new value
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// This is not a reciever, its a helper function
func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"S", "D", "H", "C"}
	cardValues := []string{"A", "2", "3"}
	// _ tells us that there was a variable there which
	// we are not planning to use, but are still need to show
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// This is not a reciever, its a helper function
// example for returning multiple value from a function
func deal(d deck, handSize int) (deck, deck) {
	// Everything up till handsize
	// from handsize afterwards
	return d[:handSize], d[handSize:]
}

// This is a reciever
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	// 0666 is a file mode that means anyone can read and write this file
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}
