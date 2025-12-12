package puzzle

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// size of shape is always 3x3
const EdgeSize = 3

// we have always 6 different shapes (0-5)
const ShapeCount = 6

type Shape struct {
	Filled [EdgeSize][EdgeSize]bool
	Count  int
}

type Region struct {
	Width    int
	Height   int
	Presents [ShapeCount]int
}

type Puzzle struct {
	Shapes  [ShapeCount]Shape
	Regions []Region
}

func Load(filename string) *Puzzle {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return Parse(string(data))
}

func Parse(text string) *Puzzle {
	p := &Puzzle{
		Regions: make([]Region, 0),
	}

	// "0:"
	reIndex := regexp.MustCompile(`^(\d):$`)
	// "###"
	reShape := regexp.MustCompile(`^[#\.]{3}$`)
	// "12x34: 12 34 56 78 90 12"
	reRegion := regexp.MustCompile(`^(\d+)x(\d+): (\d+) (\d+) (\d+) (\d+) (\d+) (\d+)`)

	currentShape := -1
	currentRow := -1

	lines := strings.Split(text, "\n")
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		if matchIndex := reIndex.FindStringSubmatch(line); matchIndex != nil {
			var err error
			currentShape, err = strconv.Atoi(matchIndex[1])
			if err != nil {
				panic(err)
			}
			if currentShape < 0 || currentShape >= ShapeCount {
				panic("shape index out of bounds")
			}
			currentRow = 0
		} else if matchShape := reShape.FindStringSubmatch(line); matchShape != nil {
			if currentShape < 0 {
				panic("found shape without index")
			}
			for x := 0; x < EdgeSize; x++ {
				if line[x] == '#' {
					p.Shapes[currentShape].Filled[currentRow][x] = true
					p.Shapes[currentShape].Count++
				} else {
					p.Shapes[currentShape].Filled[currentRow][x] = false
				}
			}
			currentRow++
		} else if matchRegion := reRegion.FindStringSubmatch(line); matchRegion != nil {
			w, err := strconv.Atoi(matchRegion[1])
			if err != nil {
				panic(err)
			}
			h, err := strconv.Atoi(matchRegion[2])
			if err != nil {
				panic(err)
			}

			r := Region{
				Width:  w,
				Height: h,
			}
			for c := 0; c < ShapeCount; c++ {
				r.Presents[c], err = strconv.Atoi(matchRegion[3+c])
				if err != nil {
					panic(err)
				}
			}

			p.Regions = append(p.Regions, r)
			//
		}
	}
	return p
}

func (p *Puzzle) Show() {
	for s, shape := range p.Shapes {
		fmt.Printf("%d:\n", s)
		for y := 0; y < EdgeSize; y++ {
			for x := 0; x < EdgeSize; x++ {
				if shape.Filled[y][x] {
					fmt.Printf("#")
				} else {
					fmt.Printf(".")
				}
			}
			fmt.Printf("\n")
		}
		fmt.Printf("\n")
	}
	for _, reg := range p.Regions {
		fmt.Printf("%d x %d:", reg.Width, reg.Height)
		for i := 0; i < ShapeCount; i++ {
			fmt.Printf(" %d", reg.Presents[i])
		}
		fmt.Printf("\n")
	}
}
