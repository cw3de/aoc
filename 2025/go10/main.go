package main

import (
	"aoc/2025/go10/puzzle"
	"fmt"
)

func task1(filename string, _ bool) {
	fmt.Printf("=========================== task 1 %s\n", filename)
	puz := puzzle.Load(filename)

	result := 0
	for _, m := range puz {

		buttons := puzzle.CountButtonsForLights(m)
		if buttons == nil {
			panic("no solution found")
		}
		result += len(buttons)
	}
	fmt.Println(filename, "task 1:", result)
}

func task2(filename string, _ bool) {
	fmt.Printf("=========================== task 2 %s\n", filename)
	puz := puzzle.Load(filename)

	result := 0
	for _, m := range puz {
		puzzle.SortButtons(m)
		puzzle.Show(m)
		count := puzzle.CountButtonsForJoltage(m)
		result += count
	}
	fmt.Println(filename, "task 2:", result)
}

func main() {
	task1("sample.txt", true) // 7
	task1("input.txt", false) // 375
	// task2("sample.txt", true) // 33
	// task2("input.txt", false)
}
