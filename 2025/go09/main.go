package main

import (
	"aoc/2025/go09/puzzle"
	"fmt"
)

func task1(filename string, verbose bool) {
	fmt.Printf("================================= task1 %s\n", filename)
	points, err := puzzle.Load(filename)
	if err != nil {
		panic(err)
	}
	result := puzzle.FindLargestRectangle(points, false, verbose)
	fmt.Println(filename, "task 1:", result)
}

func task2(filename string, verbose bool) {
	fmt.Printf("================================= task2 %s\n", filename)
	points, err := puzzle.Load(filename)
	if err != nil {
		panic(err)
	}

	mm := puzzle.FindMinMax(points)
	fmt.Printf("Min: %v Max: %v Size: %d\n", mm.Min, mm.Max,
		(mm.Max.X-mm.Min.X+1)*(mm.Max.Y-mm.Min.Y+1))

	if mm.Max.X < 30 {
		v := puzzle.NewVisual(mm.Max.X+2, mm.Max.Y+2)
		v.DrawLines(points)
		v.Show()
	}
	result := puzzle.FindLargestRectangle(points, true, verbose)
	fmt.Println(filename, "task 2:", result)
}

func main() {
	task1("sample.txt", true) // 50
	task1("input.txt", false) // 4776487744
	task2("sample.txt", true) // 24
	task2("input.txt", false) // TODO
	// 4581960734 to hight
}
