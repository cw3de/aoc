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

type Direction int

const (
	Increase Direction = iota
	Decrease
)

const Count = 100

func GetDirectionAndCount(line string) (Direction, int) {
	if line[0] == 'R' {
		value, err := strconv.Atoi(line[1:])
		if err != nil {
			panic(err)
		}
		return Increase, value
	}
	if line[0] == 'L' {
		value, err := strconv.Atoi(line[1:])
		if err != nil {
			panic(err)
		}
		return Decrease, value
	}
	panic("bad direction")
}

func task1(filename string) {
	lines := ReadLines(filename)
	result := 0
	position := 50
	// TODO: implement
	for _, line := range lines {
		if len(line) < 2 {
			continue
		}
		dir, val := GetDirectionAndCount(line)
		switch dir {
		case Decrease:
			position = (position + Count - val) % Count
		case Increase:
			position = (position + val) % Count
		}

		if position == 0 {
			result++
		}
	}
	fmt.Println(filename, "task 1:", result)
}

func task2(filename string) {
	lines := ReadLines(filename)
	result := 0
	position := 50
	for _, line := range lines {
		if len(line) < 2 {
			continue
		}
		dir, val := GetDirectionAndCount(line)
		switch dir {
		case Decrease: // left
			for val >= Count {
				result++
				val -= Count
			}
			if val == position {
				result++
				position = 0
			} else if val > position {
				if position > 0 {
					result++
				}
				position = position + Count - val
			} else {
				position = position - val
			}

		case Increase: // right
			for val >= Count {
				result++
				val -= Count
			}
			if position+val >= Count {
				result++
				position = position + val - Count
			} else {
				position = position + val
			}
		}
		fmt.Printf("%s: %d %d\n", line, position, result)
	}
	fmt.Println(filename, "task 2:", result)
}

func main() {
	task1("input.txt")
	// task2("sample.txt")
	task2("input.txt")
}
