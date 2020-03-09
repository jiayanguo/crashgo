package main

func main() {
	//var card string = "Ace of Spades"
	// card := "Ace of Spades"
	// card = "Five of Diamonds"
	// card := newCard()

	cards := newDeck()
	// cards = append(cards, "Six of Spades")

	// cards.print()

	// hand, remainingCards := deal(cards, 5)

	// hand.print()
	// remainingCards.print()

	cards.saveToFile("mycards.txt")
	
	newcards := newDeckFromFile("mycards.txt")
	newcards.shuffle()

	newcards.print()

	// for i, card := range cards {
	// 	fmt.Println(i, card)
	// }

	// fmt.Println(cards)
}

/**
package note:
Excutable package -- package main
Reusable package -- defines a package that can be used as a dependency, can be any name.
*/
