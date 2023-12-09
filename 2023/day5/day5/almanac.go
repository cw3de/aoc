package day5

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Almanac struct {
	Seeds                 []int
	ListOfTransformations []*TransformationStep
}

func parseNumbers(line string) []int {
	numbers := []int{}
	words := strings.Split(line, " ")
	for _, word := range words {
		if word == "" {
			continue
		}
		number, err := strconv.Atoi(word)
		if err != nil {
			panic(err)
		}
		numbers = append(numbers, number)
	}
	return numbers
}

func LoadAlmanac(filename string) *Almanac {

	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	almanac := &Almanac{}
	var currentMap *TransformationStep

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		} else if line[0:6] == "seeds:" {
			almanac.Seeds = parseNumbers(line[6:])
		} else if line[len(line)-5:] == " map:" {
			currentMap = NewPlantMap(line[:len(line)-5])
			almanac.ListOfTransformations = append(almanac.ListOfTransformations, currentMap)
		} else {
			numbers := parseNumbers(line)
			if len(numbers) != 3 {
				panic("Invalid line: " + line)
			}
			currentMap.ListOfRangeMaps = append(currentMap.ListOfRangeMaps, NewRangeMapWithSize(numbers[0], numbers[1], numbers[2]))
		}
	}

	return almanac
}

func ShowAlmanac(almanac *Almanac) {
	fmt.Println("Seeds:", almanac.Seeds)
	for _, m := range almanac.ListOfTransformations {
		fmt.Println("Map:", m.Name)
		for _, r := range m.ListOfRangeMaps {
			fmt.Println("Range:", r)
		}
	}
}
