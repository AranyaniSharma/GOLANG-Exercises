package main

import (
	"fmt"
	"time"
)

var deckSize int = 20

func main() {
	//var card string = "Ace of Spades"
	card1 := "Ace of spades"
	card1 = "Six of Diamonds"
	fmt.Println(card1)

	card2 := newCard()
	fmt.Println(card2)

	deckSize = 52
	fmt.Println(deckSize)

	//slice
	cards1 := []string{newCard(), newCard(), "Ace Of Diamonds"}

	cards1 = append(cards1, "Six Of Spades")
	fmt.Println(cards1)
	for i, card := range cards1 {
		fmt.Println(i, card)
	}

	cards2 := deck{newCard(), "Queen Of Diamonds"}
	cards2.print()

	cards := newDeck()
	cards.print()

	hand, remaingDeck := deal(cards, 5)
	hand.print()
	remaingDeck.print()

	fmt.Println(cards.toString())

	cards.saveToFile("my_cards")

	fmt.Println("DECK PRINT")
	deck1 := newDeckFromFile("my_cards")
	deck1.print()

	cards.shuffle()
	cards.print()

	fmt.Println(time.Now())
	fmt.Println(time.Now().Day())

	fmt.Println(time.Now().Day() + 1)
	//fmt.Println(timestamppb.N ow())

}

func newCard() string {
	return "Five Of Diamonds"
}
