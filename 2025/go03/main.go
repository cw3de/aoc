package main

import (
	"aoc/2025/go03/bank"
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

	// strip last empty line
	if len(lines[len(lines)-1]) == 0 {
		lines = lines[:len(lines)-1]
	}
	return lines
}

func task(filename string, numberOfDigits int, verbose bool) {
	lines := ReadLines(filename)
	var result int64 = 0
	for _, line := range lines {
		joltage := bank.FindBest(line, numberOfDigits)
		if verbose {
			fmt.Printf("%s: %d\n", line, joltage)
		}
		result += joltage
	}
	fmt.Println(filename, "task 2:", result)
}

func main() {
	// task("sample.txt", 2, true)
	task("input.txt", 2, false)
	// task("sample.txt", 12, true)
	task("input.txt", 12, false)
}
