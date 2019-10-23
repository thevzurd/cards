package main

func main() {
	// var card string = "Ace of Spades"
	// card := "declare variable with string type"
	// card = "Reassign"
	// Array is fixed length. Slice is a dynamic length array
	// All elements should be of the same type
	//cards := []string{newCard(), "Ace of Diamonds"} // A slice of type string
	//cards = append(cards, "6 of Diamond")           // Append does not modify cards, but overwrites
	// i is index, card is the element
	cards := deck{"A", "B", newCard()}
	cards.print()
}
func newCard() string {
	return "Five of Diamonds"
}
