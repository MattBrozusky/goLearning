package main

import (
	"fmt"
	"os"
)

func main() {
	gameDeck := newDeck()
	gameDeck.shuffle()

	players := submitPlayers()

	newBlackJackGame(players, gameDeck)
}

func errorHandler(e error) {
	if e != nil {
		fmt.Println("Error:", e)
		os.Exit(1)
	}
}
