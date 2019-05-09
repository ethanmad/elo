package elo

import (
	"math"
)

const (
	deviation uint    = 400
	k         uint    = 32
	winScore  float64 = 1.0
	drawScore float64 = 0.5
	loseScore float64 = 0.0
)

// Player contains the Player's rating.
// Using a struct to support the future addition of other information.
type Player struct {
	rating int
}

// Match contains two *Players: w is the winner, l is the loser.
// If draw is set, the order of *Players doesn't matter.
type Match struct {
	w, l *Player
	draw bool
}

// Play updates the player's ratings based on the match results.
func (m *Match) Play() {
	d := int(m.delta())
	m.w.rating += d
	m.l.rating -= d
}

// ExpectedScore returns the expected score of the Player corresponding to ratingA.
func ExpectedScore(ratingA, ratingB int) float64 {
	return 1 / (1 + math.Pow(10, float64((ratingB-ratingA))/float64(deviation)))
}

// delta returns the absolute value of the rating change each player incurs as a result of the match.
func (m *Match) delta() uint {
	expec := ExpectedScore(m.w.rating, m.l.rating)
	d := 0.0
	if m.draw {
		d = drawScore - expec
	} else {
		d = winScore - expec
	}
	d *= float64(k)
	return uint(math.Abs(math.Round(d)))
}
