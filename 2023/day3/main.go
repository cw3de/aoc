package main

import (
	"bytes"
	"fmt"
	"os"
)

func loadGrid(filename string) [][]byte {
	allData, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	lines := bytes.Split(allData, []byte("\n"))
	grid := make([][]byte, len(lines))
	for i, line := range lines {
		grid[i] = make([]byte, len(line))
		copy(grid[i], line)
	}
	return grid
}

func isDigit(value byte) bool {
	return value >= '0' && value <= '9'
}

func isSymbol(value byte) bool {
	return !isDigit(value) && value != '.'
}

func checkSymbol(grid [][]byte, row, col int) bool {
	max := len(grid[row]) - 1

	return (row > 0 && col > 0 && isSymbol(grid[row-1][col-1])) ||
		(row > 0 && isSymbol(grid[row-1][col])) ||
		(row > 0 && col < max && isSymbol(grid[row-1][col+1])) ||
		(col > 0 && isSymbol(grid[row][col-1])) ||
		(col < max && isSymbol(grid[row][col+1])) ||
		(row < max && col > 0 && isSymbol(grid[row+1][col-1])) ||
		(row < max && isSymbol(grid[row+1][col])) ||
		(row < max && col < max && isSymbol(grid[row+1][col+1]))
}

func checkNumber(grid [][]byte, row, firstDigit, lastDigit int) int {
	number := 0
	haveSymbol := false

	for i := firstDigit; i <= lastDigit; i++ {
		if checkSymbol(grid, row, i) {
			fmt.Println("found symbol", row, i)
			haveSymbol = true
		}
		number = number*10 + int(grid[row][i]-'0')
	}
	fmt.Printf("checkNumber: r=%d c=%d-%d -> n=%d %v\n", row, firstDigit, lastDigit, number, haveSymbol)
	if !haveSymbol {
		return 0
	}
	return number
}

func task1(filename string) {
	grid := loadGrid(filename)
	sum := 0
	for row, line := range grid {
		firstDigit := -1
		lastDigit := -1
		for col, value := range line {
			if isDigit(value) {
				// fmt.Println("digit", row, col)
				if firstDigit == -1 {
					firstDigit = col
				}
				lastDigit = col
			} else {
				// fmt.Println("not digit", row, col)
				if firstDigit != -1 {
					number := checkNumber(grid, row, firstDigit, lastDigit)
					if number > 0 {
						sum += number
						fmt.Println(number)
					}
					firstDigit = -1
					lastDigit = -1
				}
			}
		}

		if firstDigit != -1 {
			number := checkNumber(grid, row, firstDigit, lastDigit)
			if number > 0 {
				sum += number
				fmt.Println(number)
			}
		}
	}
	fmt.Println("sum", sum)
}

func getNumber(grid [][]byte, row, col int) int {

	if row < 0 ||
		row >= len(grid) ||
		col < 0 ||
		col >= len(grid[row]) ||
		!isDigit(grid[row][col]) {
		return 0
	}
	pos := col
	for pos > 0 && isDigit(grid[row][pos-1]) {
		pos--
	}

	number := 0
	for pos < len(grid[row]) && isDigit(grid[row][pos]) {
		number = number*10 + int(grid[row][pos]-'0')
		pos++
	}
	return number
}

// 123
// 4x5
// 678
func findNumbers(grid [][]byte, row, col int) []int {
	numbers := make([]int, 0)

	fmt.Println("findNumbers", row, col)

	if row > 0 {
		if isDigit(grid[row-1][col]) {
			// symbol obove is digit
			numbers = append(numbers, getNumber(grid, row-1, col))
		} else {
			toLeft := getNumber(grid, row-1, col-1)
			if toLeft > 0 {
				numbers = append(numbers, toLeft)
			}
			toRight := getNumber(grid, row-1, col+1)
			if toRight > 0 {
				numbers = append(numbers, toRight)
			}
		}
	}

	if row < len(grid)-1 {
		if isDigit(grid[row+1][col]) {
			// symbol below is digit
			numbers = append(numbers, getNumber(grid, row+1, col))
		} else {
			toLeft := getNumber(grid, row+1, col-1)
			if toLeft > 0 {
				numbers = append(numbers, toLeft)
			}
			toRight := getNumber(grid, row+1, col+1)
			if toRight > 0 {
				numbers = append(numbers, toRight)
			}
		}
	}

	// symbol left is digit
	toLeft := getNumber(grid, row, col-1)
	if toLeft > 0 {
		numbers = append(numbers, toLeft)
	}
	// symbol right is digit
	toRight := getNumber(grid, row, col+1)
	if toRight > 0 {
		numbers = append(numbers, toRight)
	}
	return numbers
}

func checkGear(grid [][]byte, row, col int) int {
	numbers := findNumbers(grid, row, col)
	if len(numbers) == 2 {
		return numbers[0] * numbers[1]
	} else if len(numbers) > 2 {
		fmt.Println("more than 2 numbers @", row, col)
	}

	return 0
}

func task2(filename string) {
	grid := loadGrid(filename)
	sum := 0
	for row, line := range grid {
		for col, value := range line {
			if value == '*' {
				product := checkGear(grid, row, col)
				sum += product
			}
		}
	}
	fmt.Println("sum", sum)
}

func main() {
	// task2("sample1.txt")
	task2("input.txt")
}
