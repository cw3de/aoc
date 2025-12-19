package main

import (
	"aoc/2025/go10/puzzle"
	"fmt"
	"sync"
)

func task1(filename string, _ bool) {
	fmt.Printf("=========================== task 1 %s\n", filename)
	puz := puzzle.Load(filename)

	result := 0
	for _, m := range puz {

		buttons := puzzle.CountButtonsForLights(m)
		if buttons == nil {
			panic("no solution found")
		}
		result += len(buttons)
	}
	fmt.Println(filename, "task 1:", result)
}

func task2(filename string, verbose bool) {
	fmt.Printf("=========================== task 2 %s\n", filename)
	puz := puzzle.Load(filename)

	result := 0
	maxThreads := 10
	tasks := make(chan int, maxThreads)
	var wg sync.WaitGroup

	for i, m := range puz {
		// block, if too many threads running
		wg.Add(1)
		tasks <- i
		fmt.Printf("%d: started ...\n", i)

		if verbose {
			puzzle.Show(m)
		}

		go func() {
			defer wg.Done()
			puzzle.SortButtons(m)
			r := puzzle.CountButtonsForJoltage(m)
			if verbose {
				puzzle.ShowResult(m, r)
			}
			result += r.Total
			fmt.Printf("%d: result %d\n", i, r.Total)
			<-tasks
		}()
	}
	wg.Wait()
	fmt.Println(filename, "task 2:", result)
}

func main() {
	task1("sample.txt", true)  // 7
	task1("input.txt", false)  // 375
	task2("sample.txt", false) // 33
	task2("input.txt", false)
}
