package puzzle_test

import (
	"aoc/2025/go09/puzzle"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHasPointInsideRect(t *testing.T) {

	// .......... 0
	// .......... 1
	// ..xxxxxx.. 2
	// ..x    x.. 3
	// ..x    x.. 4
	// ..x    x.. 5
	// ..x    x.. 6
	// ..xxxxxx.. 7
	// .......... 8
	// .......... 9
	// 0123456789 *

	rect := puzzle.NewRectangle(puzzle.NewPoint(2, 2), puzzle.NewPoint(7, 7))

	goodPoints := []puzzle.Point{
		puzzle.NewPoint(2, 1),
		puzzle.NewPoint(7, 1),
		puzzle.NewPoint(7, 7),
		puzzle.NewPoint(2, 7),
	}

	good := puzzle.HasPointInsideRect(goodPoints, rect, false)
	assert.False(t, good)

	badPoints := []puzzle.Point{
		puzzle.NewPoint(2, 1),
		puzzle.NewPoint(6, 1),
		puzzle.NewPoint(6, 7),
		puzzle.NewPoint(2, 7),
	}
	bad := puzzle.HasPointInsideRect(badPoints, rect, false)
	assert.True(t, bad)
}
