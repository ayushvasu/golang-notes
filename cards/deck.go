package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func newDeck() deck {
	cards := deck{}

	cardsSuites := []string{"Spades", "Diamond", "Heart", "Club"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five"}

	for _, suite := range cardsSuites {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suite)
		}
	}
	return cards
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ", ")
}

func (d deck) saveToFile(fileName string) error {
	return os.WriteFile(fileName, []byte(d.toString()), 0644)
}

func newDeckFromFile(fileName string) deck {
	bs, err := os.ReadFile(fileName)

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	return deck(strings.Split(string(bs), ", "))

}

func (d deck) shuffle() {
	for index := range d {
		newPosition := rand.Intn(len(d) - 1)
		d[index], d[newPosition] = d[newPosition], d[index]
	}
}

func (d deck) shuffle2() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for index := range d {
		newPosition := r.Intn(len(d) - 1)
		d[index], d[newPosition] = d[newPosition], d[index]
	}
}
