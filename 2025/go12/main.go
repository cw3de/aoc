package main

import (
	"aoc/2025/go12/puzzle"
	"fmt"
)

func task1(filename string, verbose bool) {
	p := puzzle.Load(filename)
	p.Show()
	result := 0
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
	task1("input.txt", false)
	// task2("sample.txt", true)
	// task2("input.txt", false)
}
