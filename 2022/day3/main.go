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

func GetValueFor(c byte) int {
	if c >= 'a' && c <= 'z' {
		return 1 + int(c-'a')
	}
	if c >= 'A' && c <= 'Z' {
		return 27 + int(c-'A')
	}
	panic("Invalid character")
}

func findMatch(left, right string) int {
	for l := 0; l < len(left); l++ {
		for r := 0; r < len(right); r++ {
			if left[l] == right[r] {
				// fmt.Println("Match:", left, right, l, r)
				return GetValueFor(left[l])
			}
		}
	}
	panic("No match found")
}

func task1(filename string) {
	lines := ReadLines(filename)
	total := 0
	for _, line := range lines {
		l := len(line)
		if l > 0 {
			v := findMatch(line[0:l/2], line[l/2:l])
			total += v
		}
	}
	fmt.Println(filename, "task1:", total)
}

func findCommon(left, right, middle string) int {
	for l := 0; l < len(left); l++ {
		for r := 0; r < len(right); r++ {
			if left[l] == right[r] {
				for m := 0; m < len(middle); m++ {
					if left[l] == middle[m] {
						v := GetValueFor(left[l])
						// fmt.Println("Match:", left, right, middle, l, r, m, "->", v)
						return v
					}
				}
			}
		}
	}
	panic("No match found")
}

func task2(filename string) {
	lines := ReadLines(filename)
	total := 0
	for l := 0; l+2 < len(lines); l += 3 {
		v := findCommon(lines[l], lines[l+1], lines[l+2])
		total += v
	}
	fmt.Println(filename, "task2:", total)
}

func main() {
	// task1("sample.txt")
	task1("input.txt")
	// task2("sample.txt")
	task2("input.txt")
}
