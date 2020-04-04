package main

import "fmt"

type deck []string

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func newCard() deck {
	cards := deck{}
	cardSuits := []string{"spade", "heart", "club", "diamonds"}
	cardValue := []string{"Ace", "one", "two", "three", "four"}

	for _, suit := range cardSuits {
		for _, val := range cardValue {
			cards = append(cards, val+" of "+suit)
		}
	}
	return cards
}

func handle(d deck, shift int) (deck, deck) {
	return d[:shift], d[shift:]
}
