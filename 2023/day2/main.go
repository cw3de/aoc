package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	red   = 0
	green = 1
	blue  = 2
)

func ParseCube(cube string) (int, int, error) {
	parts := strings.Split(cube, " ")
	if len(parts) != 2 {
		fmt.Printf("'%s' -> %v\n", cube, parts)
		return 0, 0, errors.New("invalid cube string (len)")
	}
	count, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, 0, errors.New("invalid cube string (count)")
	}

	if parts[1] == "red" {
		return count, red, nil
	}
	if parts[1] == "green" {
		return count, green, nil
	}
	if parts[1] == "blue" {
		return count, blue, nil
	}
	return 0, 0, errors.New("invalid cube string (color)")
}

func ParseGame(game string, makeProduct bool) (int, error) {

	if game[0:5] != "Game " {
		return 0, errors.New("invalid game string (no Game)")
	}
	posColon := strings.Index(game, ":")
	if posColon == -1 {
		return 0, errors.New("invalid game string (no colon)")
	}
	id, err := strconv.Atoi(game[5:posColon])
	if err != nil {
		return 0, errors.New("invalid game string (bad id)")
	}
	sets := strings.Split(string(game[posColon+1:]), ";")
	var maxRed, maxGreen, maxBlue int

	for _, set := range sets {
		fmt.Printf("set %d: '%s'\n", id, set)
		cubes := strings.Split(set, ",")
		for _, cube := range cubes {
			count, color, err := ParseCube(strings.TrimSpace(cube))
			if err != nil {
				return 0, err
			}

			if makeProduct {
				switch color {
				case red:
					maxRed = max(maxRed, count)
				case green:
					maxGreen = max(maxGreen, count)
				case blue:
					maxBlue = max(maxBlue, count)
				}
			} else {
				switch color {
				case red:
					if count > 12 {
						return 0, errors.New("invalid game string (too many red cubes)")
					}
				case green:
					if count > 13 {
						return 0, errors.New("invalid game string (too many green cubes)")
					}
				case blue:
					if count > 14 {
						return 0, errors.New("invalid game string (too many blue cubes)")
					}
				}
			}
		}
	}
	if makeProduct {
		return maxRed * maxGreen * maxBlue, nil
	}
	return id, nil
}

func task1(filename string, makeProduct bool) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	sum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		game := scanner.Text()
		result, err := ParseGame(game, makeProduct)
		if err != nil {
			fmt.Println(err, game)
		} else {
			fmt.Println(result)
			sum += result
		}
	}
	fmt.Printf("sum: %d\n", sum)
}

func main() {
	task1("input.txt", true)
}
