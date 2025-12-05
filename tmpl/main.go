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

func task1(filename string, verbose bool) {
	// lines := ReadLines(filename)
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
	// task1("input.txt", false)
	// task2("sample.txt", true)
	// task2("input.txt", false)
}
