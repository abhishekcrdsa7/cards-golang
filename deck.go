package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func newDeck() deck {
	cards := deck{}
	suits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	vals := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range suits {
		for _, val := range vals {
			cards = append(cards, val+" of "+suit)
		}
	}
	return cards
}

func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func readFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("ERROR: ", err)
		os.Exit(1)
	}
	d := strings.Split(string(bs), ",")
	return deck(d)
}

func (d deck) shuffle() {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)
	for i := range d {
		n := r.Intn(len(d) - 1)
		d[i], d[n] = d[n], d[i]
	}
}
