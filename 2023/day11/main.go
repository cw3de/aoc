package main

import (
	"fmt"

	"github.com/cw3de/aoc/galaxy"
)

func task1(filename string) {
	g := galaxy.LoadGalaxy(filename)
	ps := galaxy.FindPlanets(g)
	result := 0

	for i, p := range ps {
		for j := i + 1; j < len(ps); j++ {
			d := g.Distance(p, ps[j], 1)

			result += d
		}
	}
	fmt.Println(filename, "task 1:", result)
}

func task2(filename string) {
	g := galaxy.LoadGalaxy(filename)
	ps := galaxy.FindPlanets(g)
	result := 0

	const emptyDistance = 1000000 - 1
	for i, p := range ps {
		for j := i + 1; j < len(ps); j++ {
			d := g.Distance(p, ps[j], emptyDistance)

			result += d
		}
	}
	fmt.Println(filename, "task 2:", result)
}

func main() {
	// task1("sample.txt")
	task1("input.txt")
	// task2("sample.txt")
	task2("input.txt")
}
