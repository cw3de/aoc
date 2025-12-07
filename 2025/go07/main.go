package main

import (
	"aoc/2025/go07/puzzle"
	"fmt"
	"os"
	"strings"
)

func ReadLines(filename string) []string {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(data), "\n")
	if len(lines) == 0 {
		return lines
	}
	lastRow := len(lines) - 1
	if len(lines[lastRow]) == 0 {
		lines = lines[:lastRow]
	}
	return lines
}

func task1(filename string, verbose bool) {
	lines := ReadLines(filename)
	puz := puzzle.NewPuzzle(lines)
	result := puz.Run1(verbose)
	fmt.Println(filename, "task 1:", result)
}

func task2(filename string, verbose bool) {
	lines := ReadLines(filename)
	puz := puzzle.NewPuzzle(lines)
	result := puz.Run2(verbose)
	fmt.Println(filename, "task 2:", result)
}

func main() {
	task1("sample.txt", true) // 21
	task1("input.txt", false)
	task2("sample.txt", true) // 40
	task2("input.txt", false)
}
