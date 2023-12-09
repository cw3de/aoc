package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Elf struct {
	index    int
	sum      int
	supplies []int
}

func NewElf(i int) Elf {
	return Elf{
		index:    i,
		sum:      0,
		supplies: make([]int, 0),
	}
}

func ReadElvesFromFile(filename string) []Elf {
	fileData, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	elves := make([]Elf, 0)
	blocks := strings.Split(string(fileData), "\n\n")
	for _, block := range blocks {
		elf := NewElf(len(elves) + 1)
		lines := strings.Split(block, "\n")
		for _, line := range lines {
			if line == "" {
				continue
			}
			number, err := strconv.Atoi(line)
			if err != nil {
				fmt.Printf("Error converting %s to int\n", line)
			} else {
				elf.supplies = append(elf.supplies, number)
				elf.sum += number

			}
		}
		elves = append(elves, elf)
	}

	sort.Slice(elves, func(i, j int) bool {
		return elves[i].sum > elves[j].sum
	})

	return elves
}

func showElves(elves []Elf) {
	for _, elf := range elves {
		fmt.Println(elf)
	}
}

func task1(filename string) {
	elves := ReadElvesFromFile(filename)
	fmt.Println(filename, "Max calories:", elves[0].sum)
}

func task2(filename string) {
	elves := ReadElvesFromFile(filename)
	fmt.Println(filename, "Max 3 calories:", elves[0].sum+elves[1].sum+elves[2].sum)
}

func main() {
	// task1("sample.txt")
	task1("input.txt")
	task2("input.txt")
}
