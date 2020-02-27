package main

func main() {
	cards := readFromFile("myCards")
	cards.shuffle()
	cards.print()
}
