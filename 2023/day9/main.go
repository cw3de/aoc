package main

import (
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
	return strings.Split(string(data), "\n")
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

func getLastDiffs(numbers []int) int {
	allDiffZero := true
	diffs := []int{}
	// fmt.Println("check:", numbers)

	for i := 1; i < len(numbers); i++ {
		diff := numbers[i] - numbers[i-1]
		diffs = append(diffs, diff)
		if diff != 0 {
			allDiffZero = false
		}
	}

	lastNumber := numbers[len(numbers)-1]
	if allDiffZero {
		// fmt.Println("all diff zero")
		return lastNumber
	}
	downDiff := getLastDiffs(diffs)
	// fmt.Println("lastNumber:", lastNumber, "downDiff:", downDiff, "result:", lastNumber+downDiff)
	return lastNumber + downDiff
}

func task1(filename string) {
	lines := ReadLines(filename)
	result := 0
	for _, line := range lines {
		if line == "" {
			continue
		}
		numbers := GetIntsFromLine(line)
		value := getLastDiffs(numbers)
		result += value
	}
	fmt.Println(filename, "task 1:", result)
}

func getFirstDiffs(numbers []int) int {
	allDiffZero := true
	diffs := []int{}
	// fmt.Println("check:", numbers)

	for i := 1; i < len(numbers); i++ {
		diff := numbers[i] - numbers[i-1]
		diffs = append(diffs, diff)
		if diff != 0 {
			allDiffZero = false
		}
	}

	firstNumber := numbers[0]
	if allDiffZero {
		// fmt.Println("all diff zero")
		return firstNumber
	}
	downDiff := getFirstDiffs(diffs)
	// fmt.Println("firstNumber:", firstNumber, "downDiff:", downDiff, "result:", firstNumber-downDiff)
	return firstNumber - downDiff
}

func task2(filename string) {
	lines := ReadLines(filename)
	result := 0
	for _, line := range lines {
		if line == "" {
			continue
		}
		numbers := GetIntsFromLine(line)
		value := getFirstDiffs(numbers)
		result += value
	}
	fmt.Println(filename, "task 2:", result)
}

func main() {
	task1("sample.txt")
	task1("input.txt")
	task2("sample.txt")
	task2("input.txt")
}
