package main

import (
	"fmt"

	"github.com/cw3de/aoc/maze"
)

func task1(filename string) {
	m, err := maze.LoadMaze(filename)
	if err != nil {
		panic(err)
	}
	result := m.FindLoop(maze.Start)
	// m.ClearUnsed()
	// m.Draw(false)
	fmt.Println(filename, "task 1:", result/2)
}

func task2(filename string, startReplace byte) {
	m, err := maze.LoadMaze(filename)
	if err != nil {
		panic(err)
	}
	m.FindLoop(startReplace)
	m.ClearUnsed()
	// m.Draw(false)
	q := maze.NewQuadMaze(m)
	q.FillOutside()
	result := q.FillInside()
	m.Draw(false)
	fmt.Println(filename, "task 2:", result)
}

func main() {
	// task1("sample1.txt")
	task1("input.txt")
	task2("loop1.txt", maze.SouthEast)
	// task2("loop2.txt", maze.SouthEast)
	// task2("loop3.txt", maze.SouthEast)
	// task2("loop4.txt", maze.SouthWest)
	task2("input.txt", maze.Vertical)
}
