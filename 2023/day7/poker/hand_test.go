package poker_test

import (
	"testing"

	"github.com/cw3de/aoc/poker"
	"github.com/stretchr/testify/assert"
)

func TestValueOfCard(t *testing.T) {
	var tests = []struct {
		input    rune
		expected int
	}{
		{'2', 2},
		{'3', 3},
		{'4', 4},
		{'5', 5},
		{'6', 6},
		{'7', 7},
		{'8', 8},
		{'9', 9},
		{'T', 10},
		{'J', 11},
		{'Q', 12},
		{'K', 13},
		{'A', 14},
	}
	for _, test := range tests {
		t.Run(string(test.input), func(t *testing.T) {
			assert.Equal(t, test.expected, poker.ValueOfCard(test.input, false))
		})
	}

	assert.Equal(t, 1, poker.ValueOfCard('J', true))
}

func TestValueOfHande(t *testing.T) {
	var tests = []struct {
		input    string
		expected int
		joker    bool
		name     string
	}{
		{"A2T3K", 11402100313, false, "Highcard A"},
		{"32T3K", 20302100313, false, "Pair of 3"},
		{"JJ677", 31111060707, false, "Two pairs"},
		{"55567", 40505050607, false, "Three of a kind"},
		{"KKKQQ", 51313131212, false, "Full house"},
		{"KKKK7", 61313131307, false, "Four of a kind"},
		{"KKKKK", 71313131313, false, "Five of a kind"},
	}
	for _, test := range tests {
		t.Run(string(test.input), func(t *testing.T) {
			actual := poker.ValueOfHand(test.input, false)
			assert.Equal(t, test.expected, actual)
		})
	}

}
