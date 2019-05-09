package elo

import (
	"math"
	"testing"
)

func TestPlay(t *testing.T) {
	for _, d := range providePlay() {
		d.m.Play()
		if d.m.w.rating != d.final.w.rating {
			t.Errorf("expecting winner rating %d, got %d", d.final.w.rating, d.m.w.rating)
		}
		if d.m.l.rating != d.final.l.rating {
			t.Errorf("expecting loser rating %d, got %d", d.final.w.rating, d.m.w.rating)
		}
		if d.m.draw != d.final.draw {
			t.Errorf("expecting draw %t, got %t", d.final.draw, d.m.draw)
		}

	}
}

type dataPlay struct {
	m     Match
	final Match
}

func providePlay() []dataPlay {
	data := []dataPlay{
		{ // Win with same ratings
			Match{&Player{rating: 1000}, &Player{rating: 1000}, false},
			Match{&Player{rating: 1016}, &Player{rating: 984}, false},
		},
		{ // Draw with same ratings
			Match{&Player{rating: 1000}, &Player{rating: 1000}, true},
			Match{&Player{rating: 1000}, &Player{rating: 1000}, true},
		},
		{ // Win with different ratings
			Match{&Player{rating: 1100}, &Player{rating: 1000}, false},
			Match{&Player{rating: 1112}, &Player{rating: 988}, false},
		},
		{ // Draw with w=higher rating
			Match{&Player{rating: 1400}, &Player{rating: 1100}, true},
			Match{&Player{rating: 1389}, &Player{rating: 1111}, true},
		},
	}
	return data
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
