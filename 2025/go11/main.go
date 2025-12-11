package main

import (
	"aoc/2025/go11/puzzle"
	"fmt"
)

func task1(filename string, verbose bool) {
	p := puzzle.Load(filename)
	// p.Show()
	result := puzzle.CountPaths(p)
	fmt.Println(filename, "task 1:", result)
}

func task2(filename string, verbose bool) {
	p := puzzle.Load(filename)
	result := puzzle.CountViaDacFft(p)
	fmt.Println(filename, "task 2:", result)
}

func main() {
	task1("sample1.txt", true)
	task1("input.txt", false)
	task2("sample2.txt", true)
	task2("input.txt", false)
}
