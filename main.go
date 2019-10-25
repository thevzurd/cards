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
	//cards := deck{"A", "B", newCard()}
	// slice index starts from 0
	// range in slice[starting Index including : upto Not including]
	// cards[:3] - all elements till 3
	// cards[3:] - all elements from 3
	cards := newDeck()
	/* 	hand, remainingDeck := deal(cards, 5)
	   	hand.print()
	   	remainingDeck.print() */
	cards.saveToFile("my_cards")
	//newCards := newDeckFromFile("my_card")
	// Type Conversation
	// []byte("Hello World") -> text is converted to byte slice

}
