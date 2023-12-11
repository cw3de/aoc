package galaxy

import (
	"os"
	"strings"
)

type Galaxy struct {
	Width, Height int
	Raster        [][]byte
	ColEmpty      []bool
	RowEmpty      []bool
}

func LoadGalaxy(filename string) *Galaxy {
	g := &Galaxy{}

	lines := readLines(filename)
	for r, line := range lines {
		if len(line) == 0 {
			continue
		}
		if r == 0 {
			g.Width = len(line)
		} else {
			if len(line) != g.Width {
				panic("inconsistent width")
			}
		}
		g.Raster = append(g.Raster, []byte(line))
	}
	g.Height = len(g.Raster)
	g.ColEmpty = make([]bool, g.Width)
	g.RowEmpty = make([]bool, g.Height)

	// find empty columns
	for x := 0; x < g.Width; x++ {
		g.ColEmpty[x] = func(c int) bool {
			for y := 0; y < g.Height; y++ {
				if g.Raster[y][c] == '#' {
					return false
				}
			}
			return true
		}(x)
	}

	for y := 0; y < g.Height; y++ {
		g.RowEmpty[y] = func(r int) bool {
			for x := 0; x < g.Width; x++ {
				if g.Raster[r][x] == '#' {
					return false
				}
			}
			return true
		}(y)
	}
	return g
}

func readLines(filename string) []string {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return strings.Split(string(data), "\n")
}
