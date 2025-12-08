package puzzle_test

import (
	"aoc/2025/go08/puzzle"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {

	_, err := puzzle.Parse("123,456\n")
	assert.Error(t, err)

	_, err = puzzle.Parse("123,456,x\n")
	assert.Error(t, err)

	p, err := puzzle.Parse("123,456,789\n987,654,321\n\n123456789012345,1,9876543219876\n")
	assert.NoError(t, err)
	assert.Equal(t, []puzzle.Point{
		{123, 456, 789},
		{987, 654, 321},
		{123456789012345, 1, 9876543219876},
	}, p)
}
