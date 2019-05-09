package elo

import (
	"math"
	"testing"
)

func TestPlay(t *testing.T) {
	m := &Match{&Player{rating: 1000}, &Player{rating: 1000}, false}
	m.Play()
	if m.w.rating != 1016 && m.l.rating != 984 {
		t.Errorf("expecting winner rating 1016, got %d", m.w.rating)
	}
}

func TestExpectedScore(t *testing.T) {
	for _, d := range provideExpectedScore() {
		result := ExpectedScore(d.ratingA, d.ratingB)
		rounded := int(math.Round(result * 100))
		if rounded != d.score {
			t.Errorf("expecting %d, got %d", d.score, rounded)
		}
	}
}

type dataExpectedScore struct {
	ratingA, ratingB int
	score            int
}

func provideExpectedScore() []dataExpectedScore {
	data := []dataExpectedScore{
		{
			1000, 1000,
			50,
		},
		{
			1400, 1100,
			85,
		},
		{
			1100, 1000,
			64,
		},
		{
			1000, 1100,
			36,
		},
	}
	return data
}

func TestDelta(t *testing.T) {
	for _, d := range provideDelta() {
		result := d.m.delta()
		if result != d.delta {
			t.Errorf("expecting %d, got %d", d.delta, result)
		}
	}
}

type dataDelta struct {
	m     Match
	delta uint
}

func provideDelta() []dataDelta {
	data := []dataDelta{
		{
			Match{&Player{rating: 1000}, &Player{rating: 1000}, false},
			16,
		},
		{
			Match{&Player{rating: 1000}, &Player{rating: 1000}, true},
			0,
		},
		{
			Match{&Player{rating: 1100}, &Player{rating: 1000}, false},
			12,
		},
		{
			Match{&Player{rating: 1000}, &Player{rating: 1100}, false},
			20,
		},
		{
			Match{&Player{rating: 1100}, &Player{rating: 1400}, false},
			27,
		},
		{
			Match{&Player{rating: 1400}, &Player{rating: 1100}, false},
			5,
		},
	}
	return data
}
