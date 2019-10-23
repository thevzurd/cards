package main

import "fmt"

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
