package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of 'deck'
// which is a slice of strings

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}

	cardValue := []string{"Ace", "Two", "Three", "Four"}
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

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	s := strings.Split(string(bs), ",")
	return deck(s)
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range d {
		newPosition := r.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
