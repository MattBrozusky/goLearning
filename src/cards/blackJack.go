package main

import "fmt"

func newBlackJackGame(players allPlayers, gameDeck deck) {
	turn := 0
	for {
		turn++
		for i := 0; i < len(players); i++ {

			if turn == 1 {
				gameDeck = dealCards(gameDeck, 2, &players[i])
			} else {
				fmt.Printf("\n\n%v", gameDeck)
				gameDeck = dealCards(gameDeck, 1, &players[i])
			}

			fmt.Printf("\n%v here are your cards:\n", players[i].name)
			for j := 0; j < len(players[i].hand); j++ {
				fmt.Printf("%v ", players[i].hand[j])
			}
			fmt.Printf("\nAnd here is your score: %v\n", players[i].updateScore())
		}
		if turn == 3 {
			break
		}
	}
}

func submitPlayers() allPlayers {
	return allPlayers{
		newPlayer("Matt"),
		newPlayer("Dealer"),
		newPlayer("Tester"),
	}
}

func checkScore(c card) int {
	score := 0
	switch c.Value {
	case "Ace":
		score = 1
	case "Two":
		score = 2
	case "Three":
		score = 3
	case "Four":
		score = 4
	case "Five":
		score = 5
	case "Six":
		score = 6
	case "Seven":
		score = 7
	case "Eight":
		score = 8
	case "Nine":
		score = 9
	case "Ten":
		score = 10
	case "Jack":
		score = 11
	case "Queen":
		score = 11
	case "King":
		score = 11
	}
	return score
}
