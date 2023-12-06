package day02

import (
	"reflect"
	"testing"
)

func TestExtractGameId(t *testing.T) {
	var testCases = []struct {
		label  string
		input  string
		output int
	}{
		{"handle simple case", "Game 1: ", 1},
		{"handle full line case", "Game 2: 3 blue, 4 green, 1 red; 1 blue, 5 green, 52 red", 2},
	}

	for _, c := range testCases {
		t.Run(c.label, func(t *testing.T) {
			res := extractGameId(c.input)
			if res != c.output {
				t.Errorf("out: %d, expected: %d", res, c.output)
			}
		})
	}
}

func TestExtractGameRounds(t *testing.T) {
	var testCases = []struct {
		label  string
		input  string
		output []GameRound
	}{
		{
			"handle full line case",
			"Game 1: 3 blue, 4 green, 1 red; 1 blue, 5 green, 52 red",
			[]GameRound{
				{red: 1, blue: 3, green: 4},
				{red: 52, blue: 1, green: 5},
			},
		},

		{
			"handle spare sets",
			"Game 2: 3 blue, 4 green; 52 red",
			[]GameRound{
				{blue: 3, green: 4},
				{red: 52},
			},
		},
	}

	for _, c := range testCases {
		t.Run(c.label, func(t *testing.T) {
			res := extractGameRounds(c.input)
			if !reflect.DeepEqual(res, c.output) {
				t.Errorf("out: %q, expected: %q", res, c.output)
			}
		})
	}
}

func TestExtractGameFromLine(t *testing.T) {
	var testCases = []struct {
		label  string
		input  string
		output Game
	}{
		{
			"handle full line case",
			"Game 1: 3 blue, 4 green, 1 red; 1 blue, 5 green, 52 red",
			Game{
				id: 1,
				rounds: []GameRound{
					{red: 1, blue: 3, green: 4},
					{red: 52, blue: 1, green: 5},
				},

				maxRed:   52,
				maxGreen: 5,
				maxBlue:  3,
			},
		},

		{
			"handle spare sets",
			"Game 2: 3 blue, 4 green; 52 red",
			Game{
				id: 2,
				rounds: []GameRound{
					{blue: 3, green: 4},
					{red: 52},
				},

				maxRed:   52,
				maxGreen: 4,
				maxBlue:  3,
			},
		},
	}

	for _, c := range testCases {
		t.Run(c.label, func(t *testing.T) {
			res := extractGameFromLine(c.input)
			if !reflect.DeepEqual(*res, c.output) {
				t.Errorf("out: %q, expected: %q", res, c.output)
			}
		})
	}
}
