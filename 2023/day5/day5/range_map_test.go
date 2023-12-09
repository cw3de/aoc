package day5_test

import (
	"testing"

	"github.com/cw3de/aoc/day5"
	"github.com/stretchr/testify/assert"
)

func TestMapValue(t *testing.T) {
	rm := day5.NewRangeMapWithSize(20, 5, 5) // 5,6,7,8,9 - 20,21,22,23,24

	assert.Equal(t, 20, rm.MapValue(5))
	assert.Equal(t, 21, rm.MapValue(6))
	assert.Equal(t, 24, rm.MapValue(9))
}

func TestMapRange(t *testing.T) {
	rm := day5.NewRangeMapWithSize(20, 5, 5) // 5,6,7,8,9 - 20,21,22,23,24

	r := day5.NewRangeWithSize(5, 2) // 5,6 -> 20,21
	assert.Equal(t, 20, rm.MapRange(r).FirstValue())
	assert.Equal(t, 21, rm.MapRange(r).LastValue())

	r = day5.NewRangeWithSize(6, 3) // 6,7,8 -> 21,22,23
	assert.Equal(t, 21, rm.MapRange(r).FirstValue())
	assert.Equal(t, 23, rm.MapRange(r).LastValue())
}
