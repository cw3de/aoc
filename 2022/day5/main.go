package main

import (
	"fmt"
	"os"
	"regexp"
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

type Stack struct {
	Crates []byte
}

func (stack *Stack) Push(crate byte) {
	stack.Crates = append(stack.Crates, crate)
}

func (stack *Stack) Pop() byte {
	if len(stack.Crates) == 0 {
		panic("stack is empty")
	}
	crate := stack.Crates[len(stack.Crates)-1]
	stack.Crates = stack.Crates[:len(stack.Crates)-1]
	return crate
}

type Move struct {
	Amount int
	From   int
	To     int
}

type Game struct {
	Stacks []Stack
	Moves  []Move
}

func (game *Game) MoveCratesOneByOne(move Move) {
	for i := 0; i < move.Amount; i++ {
		crate := game.Stacks[move.From-1].Pop()
		game.Stacks[move.To-1].Push(crate)
	}
}

func (game *Game) MoveCratesAtOnce(move Move) {
	crates := []byte{}
	for i := 0; i < move.Amount; i++ {
		crate := game.Stacks[move.From-1].Pop()
		crates = append(crates, crate)
	}
	for i := len(crates) - 1; i >= 0; i-- {
		game.Stacks[move.To-1].Push(crates[i])
	}
}

func (game *Game) GetTopCrates() string {
	result := ""
	for _, stack := range game.Stacks {
		if len(stack.Crates) > 0 {
			result += string(stack.Crates[len(stack.Crates)-1])
		} else {
			result += " "
		}
	}
	return result
}

func LoadGame(filename string) *Game {
	lines := ReadLines(filename)
	game := &Game{}

	findBaseLine := func() int {
		for i, line := range lines {
			if line[0:3] == " 1 " {
				return i
			}
		}
		panic("no base line found")
	}

	baseLine := findBaseLine()
	fmt.Println(filename, "base line:", baseLine)

	for s := 0; 4*s < len(lines[baseLine]); s++ {
		col := 1 + 4*s
		stack := Stack{}
		for l := baseLine - 1; l >= 0; l-- {
			if lines[l][col] == ' ' {
				break
			}
			stack.Push(lines[l][col])
		}
		game.Stacks = append(game.Stacks, stack)
	}

	re := regexp.MustCompile(`^move (\d+) from (\d+) to (\d+)$`)
	for l := baseLine + 1; l < len(lines); l++ {
		m := re.FindStringSubmatch(lines[l])
		if m != nil {
			amount, _ := strconv.Atoi(m[1])
			from, _ := strconv.Atoi(m[2])
			to, _ := strconv.Atoi(m[3])
			move := Move{amount, from, to}
			game.Moves = append(game.Moves, move)
		}
	}
	return game
}

func ShowGame(game *Game) {
	for l := 10; l >= 0; l-- {
		for _, stack := range game.Stacks {
			if len(stack.Crates) > l {
				fmt.Printf("%c", stack.Crates[l])
			} else {
				fmt.Printf(" ")
			}
			fmt.Printf(" ")
		}
		fmt.Println()
	}
	for s, _ := range game.Stacks {
		fmt.Printf("%d ", s+1)
	}
	fmt.Println()
}

func task1(filename string) {
	game := LoadGame(filename)
	for _, move := range game.Moves {
		game.MoveCratesOneByOne(move)
	}
	ShowGame(game)
	result := game.GetTopCrates()
	fmt.Println(filename, "task 1:", result)
}

func task2(filename string) {
	game := LoadGame(filename)
	for _, move := range game.Moves {
		game.MoveCratesAtOnce(move)
	}
	ShowGame(game)
	result := game.GetTopCrates()
	fmt.Println(filename, "task 2:", result)
}

func main() {
	task1("sample.txt")
	task1("input.txt")
	task2("sample.txt")
	task2("input.txt")
}
