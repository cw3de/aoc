package grid

import "fmt"

type Grid struct {
	Height   int
	Width    int
	Original [][]bool
	Modified [][]bool
}

func NewGrid(lines []string) Grid {
	var g Grid
	g.Height = len(lines)
	g.Width = len(lines[0])
	g.Original = make([][]bool, g.Height)
	g.Modified = make([][]bool, g.Height)
	for y := range g.Height {
		orig := make([]bool, g.Width)
		modi := make([]bool, g.Width)
		for x, state := range lines[y] {
			if state == '@' {
				orig[x] = true
				modi[x] = true
			} else {
				orig[x] = false
				modi[x] = false
			}
		}
		g.Original[y] = orig
		g.Modified[y] = modi
	}
	return g
}

func (g *Grid) Show() {
	fmt.Printf("================= %d x %d\n", g.Width, g.Height)
	for y, modi := range g.Modified {
		for x, c := range modi {
			if c != g.Original[y][x] {
				if c {
					fmt.Printf("?")
				} else {
					fmt.Printf("x")
				}
			} else {
				if c {
					fmt.Printf("@")
				} else {
					fmt.Printf(".")
				}
			}
		}
		fmt.Printf("\n")
	}
}

func (g *Grid) ValidPos(x, y int) bool {
	return x >= 0 &&
		y >= 0 &&
		y < len(g.Original) &&
		x < len(g.Original[y])
}

func (g *Grid) HasRoll(x, y int) bool {
	return g.ValidPos(x, y) && g.Original[y][x]
}

func (g *Grid) CountAdjacentRolls(x, y int) int {
	count := 0
	if g.HasRoll(x-1, y-1) {
		count++
	}
	if g.HasRoll(x-1, y) {
		count++
	}
	if g.HasRoll(x-1, y+1) {
		count++
	}
	if g.HasRoll(x, y-1) {
		count++
	}
	if g.HasRoll(x, y+1) {
		count++
	}
	if g.HasRoll(x+1, y-1) {
		count++
	}
	if g.HasRoll(x+1, y) {
		count++
	}
	if g.HasRoll(x+1, y+1) {
		count++
	}
	return count
}

func (g *Grid) Set(x, y int, state bool) {
	if !g.ValidPos(x, y) {
		panic("bad position")
	}
	g.Modified[y][x] = state
}

func (g *Grid) RemoveFreeRollsOnce(verbose bool) int {
	count := 0
	for y := range g.Height {
		for x := range g.Width {
			if g.HasRoll(x, y) {
				n := g.CountAdjacentRolls(x, y)
				if verbose {
					fmt.Printf("%d", n)
				}
				if n < 4 {
					count++
					g.Set(x, y, false)
				}
			} else {
				if verbose {
					fmt.Printf(".")
				}
			}
		}
		if verbose {
			fmt.Printf(" (%d)\n", count)
		}
	}
	return count
}

func (g *Grid) ClearModified() {
	for y := range g.Height {
		for x := range g.Width {
			g.Original[y][x] = g.Modified[y][x]
		}
	}
}

func (g *Grid) RemoveFreeRollsOften(verbose bool) int {
	total := 0
	for {
		count := g.RemoveFreeRollsOnce(verbose)
		if count == 0 {
			break
		}
		g.ClearModified()
		total += count
	}
	return total
}
