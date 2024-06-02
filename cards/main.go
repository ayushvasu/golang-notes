package main

import "fmt"

func main() {
	cards := newDeck()

	cards.print()
	fmt.Println(cards.toString())
	cards.shuffle2()
	fmt.Println("---------After Shuffle------------")
	fmt.Println(cards.toString())
	fmt.Println("---------------------")
	cards.shuffle()
	hand, remainingCards := deal(cards, 5)

	fmt.Println(hand.toString())
	fmt.Println("---------------------")
	fmt.Println(remainingCards.toString())
	fmt.Println("---------------------")
	fmt.Println("Writing to File")
	hand.saveToFile("Hand1")
	fmt.Println("Done writing")
	fmt.Println("---------------------")
	fmt.Println("Reading from File")
	fmt.Println(newDeckFromFile("Hand1").toString())
}
