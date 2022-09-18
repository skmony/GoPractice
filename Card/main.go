package main

func main() {

	cards := readFromFile("my_cards")
	cards.shuffle()
	cards.print()

}
