package main

import (
	"aoc/2025/go05/puzzle"
	"fmt"
	"os"
	"strconv"
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

func GetIntsFromLine(line string) []int {
	var result []int
	for _, s := range strings.Split(line, " ") {
		if s == "" {
			continue
		}
		n, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		result = append(result, n)
	}
	return result
}

// 187179043703356
// 123456789012345

func task1(filename string, verbose bool) {
	lines := ReadLines(filename)
	result := 0
	puz := puzzle.NewPuzzle()
	for _, line := range lines {
		if len(line) == 0 {
			puz.SortAndJoin(verbose)
			continue
		}
		pos := strings.Index(line, "-")
		if pos > 0 {
			min, err := strconv.ParseInt(line[:pos], 10, 64)
			if err != nil {
				panic(err)
			}
			max, err := strconv.ParseInt(line[pos+1:], 10, 64)
			if err != nil {
				panic(err)
			}
			puz.AddRange(min, max)
			continue
		}

		id, err := strconv.ParseInt(line, 10, 64)
		if err != nil {
			panic(err)
		}
		if puz.Contains(id) {
			result++
		}
	}
	fmt.Println(filename, "task 1:", result)
}

func task2(filename string, verbose bool) {
	lines := ReadLines(filename)
	puz := puzzle.NewPuzzle()
	for _, line := range lines {
		if len(line) == 0 {
			break
		}
		pos := strings.Index(line, "-")
		if pos > 0 {
			min, err := strconv.ParseInt(line[:pos], 10, 64)
			if err != nil {
				panic(err)
			}
			max, err := strconv.ParseInt(line[pos+1:], 10, 64)
			if err != nil {
				panic(err)
			}
			puz.AddRange(min, max)
			continue
		}
	}
	puz.SortAndJoin(verbose)

	var result int64 = 0
	for _, fr := range puz.FreshRanges {
		count := fr.Max - fr.Min + 1
		result += count
		if verbose {
			fmt.Printf("%15d - %15d : %15d\n", fr.Min, fr.Max, count)
		}
	}

	fmt.Println(filename, "task 2:", result)
}

func main() {
	task1("sample.txt", true)
	task1("input.txt", false)
	task2("sample.txt", true)
	task2("input.txt", false)
}
