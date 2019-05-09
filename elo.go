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

// Play updates the ratings based on the match results by modifying the Players referenced by m.
func (m *Match) Play() {
	d := int(m.delta())
	// In the event of a draw, the lower rating should increase and the higher rating should decrease.
	if m.draw && m.w.rating > m.l.rating {
		m.w.rating -= d
		m.l.rating += d
	} else {
		m.w.rating += d
		m.l.rating -= d
	}
}

// ExpectedScore returns the expected score (chance of winning) of Player A.
// The expected score of Player B is 1 - ExpectedScore(ratingA, ratingB).
// This function can be used to predict matchups without alone.
// Equation: ExpectedScoreA = 1 / (1 + 10^{(ratingB - ratingA)/deviation}).
func ExpectedScore(ratingA, ratingB int) float64 {
	return 1 / (1 + math.Pow(10, float64((ratingB-ratingA))/float64(deviation)))
}

// delta returns the absolute value of the rating change each player incurs as a result of the match.
// It is the responsibility of the caller function to sign the returned value appropriately;
// i.e., the winner's rating is added to and the loser's rating is subtracted from.
// Equation: Δ = k * (score - expected).
func (m *Match) delta() uint {
	d := -ExpectedScore(m.w.rating, m.l.rating)
	if m.draw {
		d += drawScore
	} else {
		d += winScore
	}
	d *= float64(k)
	return uint(math.Abs(math.Round(d)))
}
