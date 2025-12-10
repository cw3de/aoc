package puzzle

import (
	"fmt"
	"slices"
)

func SortButtons(m *Machine) {
	// sort button with most light to the beginning
	slices.SortFunc(m.Buttons, func(a, b *Button) int {
		return len(b.Numbers) - len(a.Numbers)
	})
}

func CountButtonsForJoltage(m *Machine) int {
	// maximum buttons: 10
	// maximum joltage: 286

	minCount := FindMaxJoltageLevel(m)

	fmt.Printf("minimum count: %d\n", minCount)

	result := 0
	if TryButtonsForJoltage(m, 0, minCount) {
		for i, but := range m.Buttons {
			fmt.Printf("%d: button %v x %d\n", i, but.Numbers, but.Count)
			result += but.Count
		}
	} else {
		fmt.Printf("no solution\n")
	}
	return result
}

func TryButtonsForJoltage(m *Machine, b, maxCount int) bool {

	if increaseJoltage(m, m.Buttons[b], -maxCount) {
		m.Buttons[b].Count = maxCount
		return true
	}

	for count := maxCount; count >= 0; count-- {

		fmt.Printf("check button %d count %d\n", b, count)

		if b+1 < len(m.Buttons) {
			if TryButtonsForJoltage(m, b+1, maxCount) {
				m.Buttons[b].Count = count
				return true
			}
		}

		if increaseJoltage(m, m.Buttons[b], 1) {
			m.Buttons[b].Count = count
			return true
		}
	}
	return false
}

func FindMaxJoltageLevel(m *Machine) int {
	maxLevel := 0
	for _, j := range m.Joltage {
		if j > maxLevel {
			maxLevel = j
		}
	}
	return maxLevel
}

func increaseJoltage(m *Machine, b *Button, count int) bool {

	allZero := true
	for i, b := range b.Pattern {
		if b {
			m.Joltage[i] = count
			if m.Joltage[i] != 0 {
				allZero = false
			}
		}
	}
	return allZero
}
