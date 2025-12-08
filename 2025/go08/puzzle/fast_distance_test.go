package puzzle_test

import (
	"aoc/2025/go08/puzzle"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFastDistance(t *testing.T) {

	tests := []struct {
		name     string
		a        puzzle.Point
		b        puzzle.Point
		expected int64
	}{
		{"alle zero", puzzle.Point{0, 0, 0}, puzzle.Point{0, 0, 0}, 0},
		{"same values", puzzle.Point{1, 2, 3}, puzzle.Point{1, 2, 3}, 0},
		{"max distance", puzzle.Point{99999, 99999, 99999}, puzzle.Point{0, 0, 0}, 29999400003},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			actual := puzzle.FastDistance(test.a, test.b)
			assert.Equal(t, test.expected, actual)
			//
			reverse := puzzle.FastDistance(test.b, test.a)
			assert.Equal(t, test.expected, reverse)
		})
	}
}
