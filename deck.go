package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// Create a new type of 'deck'
// which is a slice of strings

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}

	cardValue := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}
	fmt.Println(cards.toString())
	for _, suit := range cardSuits {
		for _, value := range cardValue {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards

}

// the 'd' represents the 'the of cards' varible we created in main.go
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// this is telling Go that we are going to return two values as deck aka string
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

//takes what we have and turns it into a string
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

//takes the streing we created above and converts it into a byteslice
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)

}
