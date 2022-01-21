package main

func main() {
	cards := newDeck()

	// hand, remainingDeck := deal(cards, 5)
	// hand.print()
	// remainingDeck.print()

	// fmt.Println(cards.toString())
	// hand.saveToFile("myHand")

	// cards := newDeckFromFile("myHand")
	cards.shuffle()
	cards.print()	
}