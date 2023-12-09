package main

import (
	"fmt"
	"os"
	"strings"
)

const (
	Rock     = 1
	Paper    = 2
	Scissors = 3
)

const (
	Lose = 1
	Draw = 2
	Win  = 3
)

// Rock defeats Scissors, Scissors defeats Paper, and Paper defeats Rock.
// If both players choose the same shape, the round instead ends in a draw.
func Beats(a, b int) bool {
	return (a == Rock && b == Scissors) ||
		(a == Scissors && b == Paper) ||
		(a == Paper && b == Rock)
}

// was muss ich verwenden, um gegen \a s zu gewinnen?
func WinAgainst(s int) int {
	switch s {
	case Scissors:
		return Rock
	case Paper:
		return Scissors
	case Rock:
		return Paper
	}
	panic("bad choice")
}

// was muss ich verwenden, um gegen \a s zu verlieren?
func LoseAgainst(s int) int {
	switch s {
	case Rock:
		return Scissors
	case Scissors:
		return Paper
	case Paper:
		return Rock
	}
	panic("bad choice")
}

// A for Rock, B for Paper, and C for Scissors
// X for Rock, Y for Paper, and Z for Scissors (Vermutung für Aufgabe 1)

func GetSelection(c byte) int {
	switch c {
	case 'A':
		return Rock
	case 'B':
		return Paper
	case 'C':
		return Scissors
	case 'X':
		return Rock
	case 'Y':
		return Paper
	case 'Z':
		return Scissors
	}
	panic("bad selection input")
}

func GetOutcome(c byte) int {
	switch c {
	case 'X':
		return Lose
	case 'Y':
		return Draw
	case 'Z':
		return Win
	}
	panic("bad outcome input")
}

// Score for selection: 1 for Rock, 2 for Paper, and 3 for Scissors
// Score for outcome: 0 if you lost, 3 if the round was a draw, and 6 if you won
func GetScoreFor(s int) int {
	switch s {
	case Rock:
		return 1
	case Paper:
		return 2
	case Scissors:
		return 3
	}
	panic("bad input")
}

func task1(filename string) {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	total := 0
	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}
		if len(line) == 3 && line[1] == ' ' {
			op := GetSelection(line[0])
			me := GetSelection(line[2])
			score := GetScoreFor(me)
			if Beats(me, op) {
				score += 6 // win
			} else if Beats(op, me) {
				score += 0 // loss
			} else {
				score += 3 // draw
			}
			// fmt.Println(line, score)
			total += score
		} else {
			panic("bad line")
		}
	}
	fmt.Println(filename, "task1 Total:", total)
}

func task2(filename string) {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	total := 0
	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}
		if len(line) == 3 && line[1] == ' ' {
			op := GetSelection(line[0])
			out := GetOutcome(line[2])

			score := 0
			me := 0

			switch out {
			case Lose:
				me = LoseAgainst(op)
				score += 0
			case Draw:
				me = op
				score += 3
			case Win:
				me = WinAgainst(op)
				score += 6
			}
			score += GetScoreFor(me)
			// fmt.Println(line, score)
			total += score
		} else {
			panic("bad line")
		}
	}
	fmt.Println(filename, "task2 Total:", total)
}

func main() {
	task1("input.txt")
	task2("input.txt")
}
