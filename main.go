package main

func main() {
	cards := deck{"hello", alok()}
	cards = append(cards, "ggsipu")

	cards.print()

}

func alok() string {
	return "SS"
}
