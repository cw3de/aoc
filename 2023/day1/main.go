package main

import (
	"fmt"
	"os"
	"strings"
)

func task1() {

	data, err := os.ReadFile("input.txt")
	// data, err := os.ReadFile("sample.txt")
	if err != nil {
		panic(err)
	}

	code := 0
	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}
		first := -1
		last := -1
		for _, c := range line {
			if c >= '0' && c <= '9' {
				if first == -1 {
					first = int(c - '0')
				}
				last = int(c - '0')
			}
		}
		if first == -1 || last == -1 {
			fmt.Println("invalid line")
		} else {
			number := 10*first + last
			code += number
			fmt.Println(line, first, last, number, code)
		}
	}
	fmt.Println(code)
}

var dititWords = []string{
	"zero",
	"one",
	"two",
	"three",
	"four",
	"five",
	"six",
	"seven",
	"eight",
	"nine",
}

func splitDigits(input string) []int {
	result := []int{}
	for i, c := range input {
		if c >= '0' && c <= '9' {
			result = append(result, int(c-'0'))
		} else {
			for d := 1; d <= 9; d++ {
				l := len(dititWords[d])
				if i+l <= len(input) && input[i:i+l] == dititWords[d] {
					result = append(result, d)
				}
			}
		}
	}
	return result
}
func task2() {

	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	code := 0
	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}
		digits := splitDigits(line)
		if len(digits) == 0 {
			fmt.Println("invalid line")
		} else {
			first := digits[0]
			last := digits[len(digits)-1]
			number := 10*first + last
			code += number
			fmt.Println(line, first, last, number, code)
		}
	}
	fmt.Println(code)
}

func main() {
	task2()
}
