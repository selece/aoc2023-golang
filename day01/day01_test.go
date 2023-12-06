package day01

import (
	"reflect"
	"testing"
)

func TestReplaceStringsWithNumbers(t *testing.T) {
	var testCases = []struct {
		label  string
		input  string
		output string
	}{
		{"handle simple case of only numbers", "onetwothree", "on1etw2othre3e"},
		{"handle mixed garbage of only numbers", "onextwothree", "on1extw2othre3e"},
		{"handle mixed garbage of numerals and text", "onex2three", "on1ex2thre3e"},
		{"handle double-usage letters", "oneight", "on1eigh8t"},
		{"handle double-usage letters with numerals", "oneightx1nineight", "on1eigh8tx1nin9eigh8t"},

		{"aoc 1", "two1nine", "tw2o1nin9e"},
		{"aoc 2", "eightwothree", "eigh8tw2othre3e"},
		{"aoc 3", "abcone2threexyz", "abcon1e2thre3exyz"},
		{"aoc 4", "xtwone3four", "xtw2on1e3fou4r"},
		{"aoc 5", "4nineeightseven2", "4nin9eeigh8tseve7n2"},
		{"aoc 6", "zoneight234", "zon1eigh8t234"},
		{"aoc 7", "7pqrstsixteen", "7pqrstsi6xteen"},
	}

	for _, c := range testCases {
		t.Run(c.label, func(t *testing.T) {
			res := replaceStringsWithNumbers(c.input)
			if res != c.output {
				t.Errorf("out: %s, expected: %s", res, c.output)
			}
		})
	}
}

func TestFindNumbers(t *testing.T) {
	var testCases = []struct {
		label  string
		input  string
		output []string
	}{
		{"handle simple case of digits only", "1234", []string{"1", "2", "3", "4"}},
		{"handle simple case of digits and garbage", "12xgfd34", []string{"1", "2", "3", "4"}},
		{"handle mixed case of garbage and digits", "one1two2three3four4five", []string{"1", "2", "3", "4"}},

		{"aoc 1", "219", []string{"2", "1", "9"}},
		{"aoc 2", "8wo3", []string{"8", "3"}},
		{"aoc 3", "abc123xyz", []string{"1", "2", "3"}},
		{"aoc 4", "x2ne34", []string{"2", "3", "4"}},
		{"aoc 5", "49872", []string{"4", "9", "8", "7", "2"}},
		{"aoc 6", "z1ight234", []string{"1", "2", "3", "4"}},
		{"aoc 7", "7pqrst6teen", []string{"7", "6"}},
	}

	for _, c := range testCases {
		t.Run(c.label, func(t *testing.T) {
			res := findNumbers(c.input)
			if !reflect.DeepEqual(res, c.output) {
				t.Errorf("out: %q, expected: %q", res, c.output)
			}
		})
	}
}

func TestCrunchNumericStringArrayToInt(t *testing.T) {
	var testCases = []struct {
		label  string
		input  []string
		output int
	}{
		{"handle simple case of two digits", []string{"1", "2"}, 12},
		{"handle simple case of three digits", []string{"1", "2", "3"}, 13},
		{"handle simple case of one digit", []string{"1"}, 11},
	}

	for _, c := range testCases {
		t.Run(c.label, func(t *testing.T) {
			res := crunchNumericStringArrayToInt(c.input)
			if res != c.output {
				t.Errorf("out: %q, expected: %q", res, c.output)
			}
		})
	}
}

func TestProcessNumbersFromString(t *testing.T) {
	var testCases = []struct {
		label  string
		input  string
		output int
	}{
		{"handle simple case of only numbers", "onetwothree", 13},
		{"handle mixed garbage of only numbers", "onextwothree", 13},
		{"handle mixed garbage of numerals and text", "onex2three", 13},
		{"handle double-usage letters", "oneight", 11},
		{"handle double-usage letters with numerals", "oneight1xnineight", 19},

		{"aoc 1", "two1nine", 29},
		{"aoc 2", "eightwothree", 83},
		{"aoc 3", "abcone2threexyz", 13},
		{"aoc 4", "xtwone3four", 24},
		{"aoc 5", "4nineeightseven2", 42},
		{"aoc 6", "zoneight234", 14},
		{"aoc 7", "7pqrstsixteen", 76},
	}

	for _, c := range testCases {
		t.Run(c.label, func(t *testing.T) {
			res := processNumbersFromString(c.input)
			if res != c.output {
				t.Errorf("out: %q, expected: %q", res, c.output)
			}
		})
	}
}

func TestTallyNumbersFromStringArray(t *testing.T) {
	var testCases = []struct {
		label  string
		input  []string
		output int
	}{
		{
			"digits only",
			[]string{
				"1",
				"12",
				"123",
			},
			36,
		},

		{
			"letters only",
			[]string{
				"one",
				"onetwo",
				"onetwothree",
			},
			36,
		},

		{
			"mixed letters and digits only",
			[]string{
				"one",
				"one2",
				"onetwo3",
			},
			36,
		},

		{
			"garbage with letters and digits",
			[]string{
				"yonexx",
				"onebbf2a",
				"fsafonetwoasdadd3dfsf",
			},
			36,
		},

		{
			"overlap usage letters with garbage",
			[]string{
				"oneight",
				"14543fsfddeightwobnbnfdg",
				"zfdfsewttryhfghbvfiveeightghfdgnjklretnui3276dgdfc76fdsfnfiveight",
			},
			88,
		},

		{
			"aoc short",
			[]string{
				"two1nine",
				"eightwothree",
				"abcone2threexyz",
				"xtwone3four",
				"4nineeightseven2",
				"zoneight234",
				"7pqrstsixteen",
			},
			281,
		},
	}

	for _, c := range testCases {
		t.Run(c.label, func(t *testing.T) {
			res := tallyNumbersFromStringArray(c.input)
			if res != c.output {
				t.Errorf("out: %d, expected: %d", res, c.output)
			}
		})
	}
}
