package main

import (
	"fmt"
	"os"
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

func task1(filename string) {
  lines := ReadLines(filename)
  result := 0
  // TODO: implement
  fmt.Println(filename, "task 1:", result)
}

func task2(filename string) {
  lines := ReadLines(filename)
  result := 0
  // TODO: implement
  fmt.Println(filename, "task 2:", result)
}

func main() {
	task1("sample.txt")
	// task1("input.txt")
	// task2("sample.txt")
	// task2("input.txt")
}
