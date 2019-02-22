package main

type player struct {
	name  string
	hand  deck
	score int
}

type allPlayers []player

func newPlayer(name string) player {
	p := player{
		name: name,
		score: 0,
	}
	return p
}

func (p *player) updateScore() int {
	newScore := 0
	for _, card := range p.hand {
		newScore += checkScore(card)
	}
	p.score = newScore
	return newScore
}
