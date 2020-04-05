package deck

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
)

type deck []string

func (d deck) Print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func NewCard() deck {
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

func (d deck) toString() string {
	return strings.Join([]string(d), ",")

}

func (d deck) saveTodrive(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func readFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}

	s := string(bs)
	data := strings.Split(s, ",")
	return deck(data)
}

func (d deck) shuffle() {
	for i := range d {
		newPosition := rand.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
