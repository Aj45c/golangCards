package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 2000 {
		//Video tells me to use Errorf but that wasn't working so I did Error instead
		t.Error("Expected deck length of 20, but got", len(d))
	}
}
