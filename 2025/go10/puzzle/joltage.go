package puzzle

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type Result struct {
	Levels []int
	Counts []int
	Total  int
}

func NewResult(m *Machine) *Result {
	levels := make([]int, 0, len(m.Joltage))

	levels = append(levels, m.Joltage...)
	counts := make([]int, len(m.Buttons))
	return &Result{
		Levels: levels,
		Counts: counts,
		Total:  0,
	}
}

func (r *Result) String() string {
	var sb strings.Builder

	for _, l := range r.Levels {
		sb.WriteString(strconv.Itoa(l))
		sb.WriteString(" ")
	}
	sb.WriteString(": ")
	for b, c := range r.Counts {
		if c > 0 {
			sb.WriteString(strconv.Itoa(c))
			sb.WriteString("x")
			sb.WriteString(strconv.Itoa(b))
			sb.WriteString(" ")
		}
	}
	return sb.String()
}

func SortButtons(m *Machine) {
	// sort button with most lights to the beginning
	slices.SortFunc(m.Buttons, func(a, b *Button) int {
		return len(b.Numbers) - len(a.Numbers)
	})
}

func ShowResult(m *Machine, r *Result) {
	for _, l := range m.Joltage {
		fmt.Printf("%3d ", l)
	}
	fmt.Printf("\n")
	for b, but := range m.Buttons {
		for _, l := range but.Pattern {
			if l {
				fmt.Printf("%3d ", r.Counts[b])
			} else {
				fmt.Printf("  . ")
			}
		}
		fmt.Printf(" %v\n", m.Buttons[b].Numbers)
	}
	fmt.Printf("---------------- total: %d\n", r.Total)
}

func CountButtonsForJoltage(m *Machine) *Result {
	// maximum buttons: 10
	// maximum joltage: 286

	r := NewResult(m)
	if tryButtonsForJoltage(m, r, 0) {
		for _, cnt := range r.Counts {
			r.Total += cnt
		}
	} else {
		fmt.Printf("no solution\n")
	}
	return r
}

func tryButtonsForJoltage(m *Machine, r *Result, b int) bool {

	but := m.Buttons[b]
	maxCount := findMaximumCount(r, but)
	// fmt.Printf("%v  maxcount button %d %s is %d\n", r, b, but.String(), maxCount)

	if increaseJoltage(r, but, -maxCount) {
		r.Counts[b] = maxCount
		return true
	}

	for count := maxCount; count >= 0; count-- {

		r.Counts[b] = count
		// fmt.Printf("%v  try %d x button %d %s\n", r, count, b, but.String())

		if b+1 < len(m.Buttons) {
			if tryButtonsForJoltage(m, r, b+1) {
				return true
			}
		}

		if count > 0 {
			if increaseJoltage(r, but, 1) {
				r.Counts[b] = count
				return true
			}
		}
	}

	return false
}

func findMaximumCount(r *Result, b *Button) int {
	maxLevel := 0
	for i, level := range r.Levels {
		if b.Pattern[i] {
			if maxLevel == 0 || level < maxLevel {
				maxLevel = level
			}
		}
	}
	return maxLevel
}

func increaseJoltage(r *Result, b *Button, count int) bool {

	allZero := true
	for i, b := range b.Pattern {
		if b {
			r.Levels[i] += count
		}
		if r.Levels[i] != 0 {
			allZero = false
		}
	}
	return allZero
}
