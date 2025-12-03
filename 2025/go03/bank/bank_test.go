package bank_test

import (
	"aoc/2025/go03/bank"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindBest(t *testing.T) {

	var tests = []struct {
		input    string
		length   int
		expected int64
	}{
		{"987654321111111", 2, 98},
		{"811111111111119", 2, 89},
		{"234234234234278", 2, 78},
		{"818181911112111", 2, 92},
		{"987654321111111", 12, 987654321111},
		{"811111111111119", 12, 811111111119},
		{"234234234234278", 12, 434234234278},
		{"818181911112111", 12, 888911112111},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			actual := bank.FindBest(test.input, test.length)
			assert.Equal(t, test.expected, actual)
		})
	}
}
