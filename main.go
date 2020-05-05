package main

func main() {
	cards := new readFromFile("myCards")
	cards.shuffle()
	cards.print()
}
