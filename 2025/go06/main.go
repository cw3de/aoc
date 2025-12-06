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
	lines := strings.Split(string(data), "\n")
	if len(lines) == 0 {
		return lines
	}
	// remove last row if empty
	lastRow := len(lines) - 1
	if len(lines[lastRow]) == 0 {
		lines = lines[:lastRow]
	}
	return lines
}

func task1(filename string, verbose bool) {
	lines := ReadLines(filename)

	var result int64 = 0
	values := make([]int64, 0)
	width := 0
	for y, line := range lines {
		fields := strings.Fields(line)
		if y == 0 {
			width = len(fields)
			if verbose {
				fmt.Printf("width = %d\n", width)
			}
		}

		for x, field := range fields {
			switch field {
			case "+":
				var sum int64 = 0
				for i := x; i < len(values); i += width {
					sum += values[i]
					if verbose {
						fmt.Printf("add (%d) %d\n", i, values[i])
					}
				}
				result += sum
			case "*":
				var prod int64 = 1
				for i := x; i < len(values); i += width {
					prod *= values[i]
					if verbose {
						fmt.Printf("mul (%d) %d\n", i, values[i])
					}
				}
				result += prod
			default:
				val, err := strconv.ParseInt(field, 10, 64)
				if err != nil {
					panic(err)
				}
				values = append(values, val)
			}
		}
	}
	fmt.Println(filename, "task 1:", result)
}

type Calculator struct {
	Values []int64
}

func NewCalculator() Calculator {
	return Calculator{
		Values: make([]int64, 0),
	}
}

func (calc *Calculator) Push(val int64) {
	if val > 0 {
		calc.Values = append(calc.Values, val)
	}
}

func (calc *Calculator) Work(init int64, op func(int64, int64) int64) int64 {
	res := init
	for _, v := range calc.Values {
		res = op(res, v)
	}
	calc.Values = calc.Values[:0]
	return res
}

func (calc *Calculator) Sum(val int64) int64 {
	return calc.Work(val, func(a, b int64) int64 { return a + b })
}

func (calc *Calculator) Mul(val int64) int64 {
	return calc.Work(val, func(a, b int64) int64 { return a * b })
}

func task2(filename string, verbose bool) {
	lines := ReadLines(filename)
	if len(lines) < 2 {
		panic("too few lines")
	}
	lastLine := len(lines) - 1
	symbols := lines[lastLine]
	width := len(symbols)
	var result int64 = 0
	calc := NewCalculator()

	for x := width - 1; x >= 0; x-- {
		var value int64 = 0
		for y := 0; y < lastLine; y++ {
			ch := lines[y][x]
			if ch >= '0' && ch <= '9' {
				value = 10*value + int64(ch-'0')
			} else if ch != ' ' {
				panic(fmt.Errorf("illegal character '%c' in line %d column %d", ch, y, x))
			}
		}
		switch symbols[x] {
		case ' ':
			calc.Push(value)
			if verbose {
				fmt.Printf("value %d\n", value)
			}
		case '+':
			sum := calc.Sum(value)
			if verbose {
				fmt.Printf("value %d -> sum %d\n", value, sum)
			}
			result += sum
			x--
		case '*':
			prod := calc.Mul(value)
			if verbose {
				fmt.Printf("value %d -> product %d\n", value, prod)
			}
			result += prod
			x--
		default:
			panic(fmt.Errorf("illegal charcater '%c' in last line column %d",
				symbols[x], x))
		}
	}

	fmt.Println(filename, "task 2:", result)
}

func main() {
	task1("sample.txt", true)
	task1("input.txt", false)
	task2("sample.txt", true)
	task2("input.txt", false)
}
