package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Card struct {
	Id      int
	Count   int
	Winners []int
	Number  []int
}

func ParseInt(num string) int {
	s := 0
	for s < len(num) && num[s] == ' ' {
		s++
	}
	val, err := strconv.Atoi(num[s:])
	if err != nil {
		panic(err)
	}
	return val
}

func ParseIntList(s string) []int {
	list := make([]int, 0)
	for _, num := range strings.Split(s, " ") {
		if len(num) == 0 {
			continue
		}

		list = append(list, ParseInt(num))
	}
	return list
}

func readCards(filename string) []Card {
	cards := make([]Card, 0)
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	for _, line := range strings.Split(string(data), "\n") {
		prePost := strings.Split(line, ":")
		if len(prePost) != 2 || line[0:5] != "Card " {
			continue
		}
		card := Card{}
		card.Id = ParseInt(prePost[0][5:])
		card.Count = 1
		winLoss := strings.Split(prePost[1], "|")
		card.Winners = ParseIntList(winLoss[0])
		card.Number = ParseIntList(winLoss[1])
		cards = append(cards, card)
	}
	return cards
}

func GetMatches(winners []int, number []int) int {
	matches := 0
	for _, num := range number {
		for _, winner := range winners {
			if winner == num {
				matches++
			}
		}
	}
	return matches
}

func GetPoints(matches int) int {
	if matches == 0 {
		return 0
	}
	return 1 << (matches - 1)
}

func task1(filename string) {
	cards := readCards(filename)
	sum := 0
	for _, card := range cards {
		matches := GetMatches(card.Winners, card.Number)
		points := GetPoints(matches)
		// fmt.Println(card.Id, card.Winners, card.Number, "->", matches, points)
		sum += points
	}
	fmt.Println(filename, "task 1:", sum)
}

func task2(filename string) {
	cards := readCards(filename)
	total := 0
	for i, card := range cards {
		matches := GetMatches(card.Winners, card.Number)
		for copy := 0; copy < matches; copy++ {
			j := i + 1 + copy
			if j < len(cards) {
				cards[j].Count += card.Count
			}
		}
		total += card.Count
	}
	fmt.Println(filename, "task 2:", total)
}

func main() {
	task1("sample.txt")
	task1("input.txt")
	task2("sample.txt")
	task2("input.txt")
}
