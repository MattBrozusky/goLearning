package main

import (
	"fmt"
	"os"
)

func main() {
	gameDeck := newDeck()

	deckOne, gameDeck := dealCards(gameDeck, 10)

	deckOne.shuffle()
	deckOne.print()
	deckOne.saveToJSON("deckOne")

	deckFromFile := newDeckFromJSON("deckOne")

	deckFromFile.print()

}

func errorHandler(e error) {
	if e != nil {
		fmt.Println("Error:", e)
		os.Exit(1)
	}
}
