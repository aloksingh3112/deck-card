package main

import "github.com/aloksingh3112/card-game/deck"

func main() {
	cards := deck.newCard()
	cards.shuffle()
	cards.print()

	// cards := newCard()
	// cards := readFromFile("my_file")
	// cards.print()
	// fmt.Println(cards)

	// cards.saveTodrive("my_file")

	// // firstShift, lastShift := handle(cards, 5)

	// // firstShift.print()
	// // lastShift.print()
	// fmt.Println(cards.toString())

	// // cards.print()

}
