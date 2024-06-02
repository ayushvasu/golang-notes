package main

import (
	"fmt"
	"strconv"
)

var deckSize int64

func main() {
	var card string = "Ace of spades"
	card1 := "Ace of card"
	card2 := "Ace of card 2"
	deckSize = 52
	fmt.Println(card1 + card2 + " : " + card + " : " + strconv.FormatInt(deckSize, 10))
	fmt.Printf("Output : %s %s %d \n", card1, card2, deckSize)

	card3 := getValue()

	fmt.Println(card3)
	cards := []string{"Ace of Diamonds", newCard()}
	cards = append(cards, "Six of Spades")

	fmt.Println(cards)

	for i, card := range cards {
		fmt.Println(i, card)
	}

	str := "ayush vasu"
	saveToFile(str)

}
