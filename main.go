package main

func main() {
	cards := newCard()

	firstShift, lastShift := handle(cards, 5)

	firstShift.print()
	lastShift.print()

	// cards.print()

}
