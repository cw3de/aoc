package main

import (
	"aoc/2025/go04/grid"
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
	lastRow := len(lines) - 1
	if len(lines[lastRow]) == 0 {
		lines = lines[:lastRow]
	}
	return lines
}

func task1(filename string, verbose bool) {
	lines := ReadLines(filename)
	g := grid.NewGrid(lines)
	result := g.RemoveFreeRollsOnce(verbose)
	fmt.Println(filename, "task 1:", result)
}

func task2(filename string, verbose bool) {
	lines := ReadLines(filename)
	g := grid.NewGrid(lines)
	result := g.RemoveFreeRollsOften(verbose)
	fmt.Println(filename, "task 2:", result)
}

func main() {
	task1("sample.txt", false)
	task1("input.txt", false)
	task2("sample.txt", false)
	task2("input.txt", false)
}
