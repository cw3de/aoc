package main

import (
	"aoc/2025/go12/puzzle"
	"fmt"
)

func task1(filename string, verbose bool) {
	p := puzzle.Load(filename)
	if verbose {
		p.Show()
	}
	result := 0
	for i, reg := range p.Regions {
		blockCount := 0
		for s := 0; s < puzzle.ShapeCount; s++ {
			blockCount += p.Shapes[s].Count * reg.Presents[s]
		}
		fieldCount := reg.Height * reg.Width

		var fits bool
		if blockCount <= fieldCount {
			fits = true
			result++
		} else {
			fits = false
		}

		fmt.Printf("%d: need %d have %d: %v\n", i, blockCount, fieldCount, fits)
	}
	// TODO: implement
	fmt.Println(filename, "task 1:", result)
}

func task2(filename string, verbose bool) {
	// lines := ReadLines(filename)
	result := 0
	// TODO: implement
	fmt.Println(filename, "task 2:", result)
}

func main() {
	task1("sample.txt", true)
	task1("input.txt", false) // 403
	// task2("sample.txt", true)
	// task2("input.txt", false)
}
