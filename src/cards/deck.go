package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"time"
)

//Create a new type of deck
//Which is a slice of strings
type card struct {
	Value string `json:"value"`
	Suit  string `json:"suit"`
}

type deck []card

func newDeck() deck {
	cardsSlice := deck{}

	cardsSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range cardsSuits {
		for _, values := range cardValues {
			oneCard := card{
				Suit:  suit,
				Value: values,
			}
			cardsSlice = append(cardsSlice, oneCard)
		}
	}
	return cardsSlice
}

func dealCards(d deck, handSize int, p *player) deck {
	gameDeck := d[handSize:]
	cardsForPlayer := d[:handSize]

	p.hand = append(p.hand, cardsForPlayer...)
	return gameDeck
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func (d deck) saveToJSON(fileName string) error {
	data, err := json.Marshal(d)
	errorHandler(err)
	return ioutil.WriteFile(fileName, data, 0666)
}

func newDeckFromJSON(fileName string) deck {
	jsonFile, err := os.Open(fileName)
	errorHandler(err)
	bv, e := ioutil.ReadAll(jsonFile)
	errorHandler(e)
	var deckFromFile deck
	errorHandler(json.Unmarshal(bv, deckFromFile))
	return deckFromFile
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}

// func (d deck) toString() string {
// 	var deckString string
// 	for i, card := range d {
// 		if i == len(d)-1 {
// 			deckString += card.Value + " " + card.Suit
// 		} else {
// 			deckString += card.Value + " " + card.Suit + ", "
// 		}
// 	}
// 	return deckString
// }
