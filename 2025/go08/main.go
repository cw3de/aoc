package main

import (
	"aoc/2025/go08/puzzle"
	"fmt"
)

func task1(filename string, maxConnections int, verbose bool) {

	points, err := puzzle.Load(filename)
	if err != nil {
		panic(err)
	}
	junctions := puzzle.MakeSortedJunctionList(points, verbose)
	circuits, _ := puzzle.MakeCircuitList(junctions, maxConnections, 0, verbose)
	result := circuits[0].Count() * circuits[1].Count() * circuits[2].Count()
	fmt.Println(filename, "task 1:", result)
}

func task2(filename string, verbose bool) {
	points, err := puzzle.Load(filename)
	if err != nil {
		panic(err)
	}
	junctions := puzzle.MakeSortedJunctionList(points, verbose)
	_, junction := puzzle.MakeCircuitList(junctions, 0, len(points), verbose)
	result := points[junction.PointA].X * points[junction.PointB].X
	fmt.Println(filename, "task 2:", result)
}

func main() {
	task1("sample.txt", 10, false)  // 40
	task1("input.txt", 1000, false) // 68112
	task2("sample.txt", false)      // 25272
	task2("input.txt", false)       // 44543856
}
