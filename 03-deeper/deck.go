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
// which is a slide of strings
type deck []string

func newDeck() deck {
	cards := deck{}
	
	cardSuits :=[]string{"Spades", "Diamons", "Hearts", "Club"}
	cardValues :=[]string{"Ace", "Two", "Three", "Four"}

	for _, cs := range cardSuits{
		for _, cv := range cardValues{
			cards = append(cards, cv+ " of "+cs)
		}
	}
	return cards
}

// Receiver
func (d deck) print () {
	for _, card := range d {
		fmt.Println(card)
	}
}

func deal (d deck, handSize int) (deck, deck){
	return d[:handSize], d[handSize:]
}

func (d deck)toString() string{
	return strings.Join([]string(d), ",")
}

func (d deck)  saveToFile(filename string) error{
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("error", err)
		os.Exit(1)	
	}
	s := strings.Split(string(bs), ",")
	return deck(s)	
}

func (d deck) shuffle(){
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range d {
		newPos := r.Intn(len(d) - 1)
		d[i], d[newPos] = d[newPos], d[i]
	}
}