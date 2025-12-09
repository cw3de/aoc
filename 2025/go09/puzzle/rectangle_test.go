package puzzle_test

import (
	"aoc/2025/go09/puzzle"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRectangle(t *testing.T) {

	tests := []struct {
		Name   string
		A      puzzle.Point
		B      puzzle.Point
		Top    int
		Bottom int
		Left   int
		Right  int
		Size   int
	}{
		{
			Name: "Minimal",
			A:    puzzle.NewPoint(1, 1), B: puzzle.NewPoint(1, 1),
			Top: 1, Bottom: 1, Left: 1, Right: 1, Size: 1,
		},
		{
			Name: "Forward: first point is left/above second point",
			A:    puzzle.NewPoint(5, 12), B: puzzle.NewPoint(7, 17),
			Top: 12, Bottom: 17, Left: 5, Right: 7, Size: 18,
		},
		{
			Name: "Backword: first point is right/under second point",
			A:    puzzle.NewPoint(7, 17), B: puzzle.NewPoint(5, 12),
			Top: 12, Bottom: 17, Left: 5, Right: 7, Size: 18,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			r := puzzle.NewRectangle(test.A, test.B)
			assert.Equal(t, test.Top, r.Top(), "Top")
			assert.Equal(t, test.Bottom, r.Bottom(), "Bottom")
			assert.Equal(t, test.Left, r.Left(), "Left")
			assert.Equal(t, test.Right, r.Right(), "Right")
			assert.Equal(t, test.Size, r.GetSize(), "Size")
		})
	}
}

func TestPointInside(t *testing.T) {

	rect := puzzle.NewRectangle(
		puzzle.NewPoint(3, 3),
		puzzle.NewPoint(5, 5),
	)

	tests := []struct {
		Name   string
		P      puzzle.Point
		Inside bool
	}{
		{
			Name:   "Inside",
			P:      puzzle.NewPoint(4, 4),
			Inside: true,
		},
		{
			Name:   "Left",
			P:      puzzle.NewPoint(2, 4),
			Inside: false,
		},
		{
			Name:   "Right",
			P:      puzzle.NewPoint(6, 4),
			Inside: false,
		},
		{
			Name:   "Above",
			P:      puzzle.NewPoint(4, 2),
			Inside: false,
		},
		{
			Name:   "Below",
			P:      puzzle.NewPoint(4, 6),
			Inside: false,
		},
		{
			Name:   "Top Edge",
			P:      puzzle.NewPoint(4, 3),
			Inside: false,
		},
		{
			Name:   "Bottom Edge",
			P:      puzzle.NewPoint(5, 3),
			Inside: false,
		},
		{
			Name:   "Left Edge",
			P:      puzzle.NewPoint(3, 4),
			Inside: false,
		},
		{
			Name:   "Right Edge",
			P:      puzzle.NewPoint(5, 4),
			Inside: false,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			assert.Equal(t, test.Inside, rect.PointInside(test.P))
		})
	}
}
