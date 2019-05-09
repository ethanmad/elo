package elo

import (
	"math"
	"testing"
)

func TestPlay(t *testing.T) {
	for _, d := range providePlay() {
		d.m.Play()
		if d.m.W.Rating != d.final.W.Rating {
			t.Errorf("expecting winner rating %d, got %d", d.final.W.Rating, d.m.W.Rating)
		}
		if d.m.L.Rating != d.final.L.Rating {
			t.Errorf("expecting loser rating %d, got %d", d.final.W.Rating, d.m.W.Rating)
		}
		if d.m.Draw != d.final.Draw {
			t.Errorf("expecting draw %t, got %t", d.final.Draw, d.m.Draw)
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
			Match{&Player{Rating: 1000}, &Player{Rating: 1000}, false},
			Match{&Player{Rating: 1016}, &Player{Rating: 984}, false},
		},
		{ // Draw with same ratings
			Match{&Player{Rating: 1000}, &Player{Rating: 1000}, true},
			Match{&Player{Rating: 1000}, &Player{Rating: 1000}, true},
		},
		{ // Win with different ratings
			Match{&Player{Rating: 1100}, &Player{Rating: 1000}, false},
			Match{&Player{Rating: 1112}, &Player{Rating: 988}, false},
		},
		{ // Draw with w=higher rating
			Match{&Player{Rating: 1400}, &Player{Rating: 1100}, true},
			Match{&Player{Rating: 1389}, &Player{Rating: 1111}, true},
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
			Match{&Player{Rating: 1000}, &Player{Rating: 1000}, false},
			16,
		},
		{
			Match{&Player{Rating: 1000}, &Player{Rating: 1000}, true},
			0,
		},
		{
			Match{&Player{Rating: 1100}, &Player{Rating: 1000}, false},
			12,
		},
		{
			Match{&Player{Rating: 1000}, &Player{Rating: 1100}, false},
			20,
		},
		{
			Match{&Player{Rating: 1100}, &Player{Rating: 1400}, false},
			27,
		},
		{
			Match{&Player{Rating: 1400}, &Player{Rating: 1100}, false},
			5,
		},
	}
	return data
}
