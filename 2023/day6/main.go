package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Race struct {
	Time     int
	Distance int
}

func GetNumbers(line string) []int {
	result := []int{}
	words := strings.Split(line, " ")
	for _, word := range words {
		if len(word) > 0 {
			number, err := strconv.Atoi(word)
			if err != nil {
				panic(err)
			}
			result = append(result, number)
		}
	}
	return result
}

func checkTime(speed int, race Race) bool {
	remainingTime := race.Time - speed
	distance := speed * remainingTime
	return distance > race.Distance
}

func checkInterval(minTime, maxTime int, race Race) int {
	possibilities := 0
	if (maxTime - minTime) <= 10 {
		for speed := minTime; speed < maxTime; speed++ {
			if checkTime(speed, race) {
				possibilities++
			}
		}
	} else {
		if checkTime(minTime, race) && checkTime(maxTime, race) {
			possibilities = maxTime - minTime
		} else {
			halfTime := (minTime + maxTime) / 2
			possibilities += checkInterval(minTime, halfTime, race)
			possibilities += checkInterval(halfTime, maxTime, race)
		}
	}
	return possibilities
}
func checkRace(race Race) int {
	return checkInterval(1, race.Time, race)
}

func task1(filename string) {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(data), "\n")
	if len(lines) < 2 {
		panic("too few lines")
	}
	if lines[0][0:5] != "Time:" {
		panic("line 1 must start with 'Time:'")
	}
	times := GetNumbers(lines[0][5:])
	if lines[1][0:9] != "Distance:" {
		panic("line 2 must start with 'Distance:'")
	}
	dists := GetNumbers(lines[1][9:])
	fmt.Println(times)
	fmt.Println(dists)

	product := 1
	for r := 0; r < len(times); r++ {
		race := Race{times[r], dists[r]}
		factor := checkRace(race)
		product *= factor
	}
	fmt.Println("task1 Product:", product)
}

func getOneNumber(line string) int {
	number := 0
	for _, c := range line {
		if c >= '0' && c <= '9' {
			number = number*10 + int(c-'0')
		}
	}
	return number
}

func task2(filename string) {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(data), "\n")
	if len(lines) < 2 {
		panic("too few lines")
	}
	if lines[0][0:5] != "Time:" {
		panic("line 1 must start with 'Time:'")
	}
	maxTime := getOneNumber(lines[0])
	if lines[1][0:9] != "Distance:" {
		panic("line 2 must start with 'Distance:'")
	}
	dist := getOneNumber(lines[1])
	race := Race{maxTime, dist}
	possCount := checkRace(race)
	fmt.Println("task2 possibilities:", possCount)
}

func main() {
	task1("input.txt")
	task2("input.txt")
}
