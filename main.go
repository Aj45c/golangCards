package main

import "fmt"

func main() {
	//var card string = "Ace of Spades"
	cards := newDeck()
	fmt.Println(cards.toString())
}
