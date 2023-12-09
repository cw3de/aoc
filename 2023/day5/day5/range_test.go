package day5_test

import (
	"testing"

	"github.com/cw3de/aoc/day5"
	"github.com/stretchr/testify/assert"
)

func TestNewRange(t *testing.T) {
	r0 := day5.NewRangeWithSize(0, 0)
	assert.True(t, r0.IsEmpty())

	r5 := day5.NewRangeWithSize(5, 5)
	assert.Equal(t, 5, r5.FirstValue())
	assert.Equal(t, 9, r5.LastValue())
	assert.False(t, r5.IsEmpty())
}

func TestContainsValue(t *testing.T) {
	r := day5.NewRangeWithSize(5, 5)
	assert.False(t, r.ContainsValue(4))
	assert.True(t, r.ContainsValue(5))
	assert.True(t, r.ContainsValue(9))
	assert.False(t, r.ContainsValue(10))

}

func TestOverlapsWith(t *testing.T) {
	r1 := day5.NewRangeWithSize(5, 5)
	r2 := day5.NewRangeWithSize(0, 5)
	r3 := day5.NewRangeWithSize(10, 5)
	r4 := day5.NewRangeWithSize(7, 5)
	r5 := day5.NewRangeWithSize(3, 5)

	assert.False(t, r1.OverlapsWith(r2))
	assert.False(t, r1.OverlapsWith(r3))
	assert.True(t, r1.OverlapsWith(r4))
	assert.True(t, r1.OverlapsWith(r5))
}

func TestGetIntersection(t *testing.T) {
	r1 := day5.NewRangeWithSize(7, 5)  // 7,8,9,10,11
	r2 := day5.NewRangeWithSize(5, 5)  // 5,6,7,8,9
	r3 := day5.GetIntersection(r1, r2) // 7,8,9
	assert.Equal(t, 3, r3.Size())
	assert.Equal(t, 7, r3.FirstValue())
}

func TestSortRanges(t *testing.T) {
	ranges := []day5.Range{
		day5.NewRangeWithSize(5, 5),
		day5.NewRangeWithSize(7, 5),
		day5.NewRangeWithSize(4, 5),
	}
	assert.Equal(t, 3, len(ranges))

	ranges = day5.SortRanges(ranges)
	assert.Equal(t, 3, len(ranges))
	assert.Equal(t, 4, ranges[0].FirstValue())
	assert.Equal(t, 5, ranges[1].FirstValue())
	assert.Equal(t, 7, ranges[2].FirstValue())
}

func TestJoinRanges(t *testing.T) {
	ranges := []day5.Range{
		day5.NewRangeWithSize(14, 5), // 14,15,16,17,18
		day5.NewRangeWithSize(7, 5),  // 7,8,9,10,11
		day5.NewRangeWithSize(5, 5),  // 5,6,7,8,9
	}
	assert.Equal(t, 3, len(ranges))

	ranges = day5.JoinRanges(ranges)
	assert.Equal(t, 2, len(ranges))
	assert.Equal(t, 5, ranges[0].FirstValue()) // 5..11
	assert.Equal(t, 7, ranges[0].Size())
	assert.Equal(t, 14, ranges[1].FirstValue()) // 14..18
}
