package puzzle

import "fmt"

type Field rune

const (
	Illegal  Field = '?'
	Empty    Field = '.'
	Start    Field = 'S'
	Splitter Field = '^'
	Beam     Field = '|'
	Splitted Field = ':'
)

type Puzzle struct {
	Rows     [][]Field
	WayCount [][]int
	Height   int
	Width    int
}

func NewPuzzle(lines []string) *Puzzle {
	h := len(lines)
	w := len(lines[0])
	p := &Puzzle{
		Rows:     make([][]Field, h),
		WayCount: make([][]int, h),
		Height:   h,
		Width:    w,
	}
	for y, line := range lines {
		if len(line) != w {
			panic(fmt.Errorf("expected length %d, but got %d", w, len(line)))
		}
		r := make([]Field, w)
		wc := make([]int, w)
		for x, c := range line {
			r[x] = Field(c)
			wc[x] = 0
		}
		p.Rows[y] = r
		p.WayCount[y] = wc
	}

	return p
}

func (p *Puzzle) Show() {
	for _, row := range p.Rows {
		for _, c := range row {
			fmt.Printf("%c", c)
		}
		fmt.Printf("\n")
	}
}

func (p *Puzzle) Get(x, y int) Field {
	if y < 0 || y >= len(p.Rows) {
		return Illegal
	}
	row := p.Rows[y]
	if x < 0 || x >= len(row) {
		return Illegal
	}
	return row[x]
}

func (p *Puzzle) Set(x, y int, c Field) bool {
	if y < 0 || y >= len(p.Rows) {
		fmt.Printf("illegal row %d\n", y)
		return false
	}
	if x < 0 || x >= len(p.Rows[y]) {
		fmt.Printf("illegal col %d\n", x)
		return false
	}
	if p.Rows[y][x] != Empty {
		return false
	}
	p.Rows[y][x] = c
	return true
}

func (p *Puzzle) Run1(verbose bool) int {
	count := 0

	for y, row := range p.Rows {
		for x, c := range row {
			above := p.Get(x, y-1)
			switch c {
			case Start:
			case Empty:
				if above == Beam || above == Start || above == Splitted {
					p.Set(x, y, Beam)
				}
			case Splitter:
				if above == Beam {
					leftSplit := p.Set(x-1, y, Splitted)
					rightSplit := p.Set(x+1, y, Splitted)

					if leftSplit || rightSplit {
						// count only if we got two options
						count++
					}
				}

			}
		}
		if verbose {
			for _, c := range row {
				fmt.Printf("%c", c)
			}
			fmt.Printf("  (%d)\n", count)
		}
	}

	return count
}

func (p *Puzzle) Run2(verbose bool) int {
	lastRow := p.Width - 1
	for y, row := range p.Rows {
		sum := 0
		for x, c := range row {
			count := 0
			switch c {
			case Start:
				count = 1
			case Empty:
				if y > 0 {
					count = p.WayCount[y-1][x] // was von oben durchfällt
					if x > 0 && p.Get(x-1, y) == Splitter {
						count += p.WayCount[y-1][x-1]
					}
					if x < lastRow && p.Get(x+1, y) == Splitter {
						count += p.WayCount[y-1][x+1]
					}
				}
			}
			p.WayCount[y][x] = count
			sum += count
			if verbose {
				if c == Empty {
					fmt.Printf("%2d ", count)
				} else {
					fmt.Printf(" %c ", c)
				}
			}
		}
		if verbose {
			fmt.Printf("  (%d)\n", sum)
		}
	}

	total := 0
	for _, c := range p.WayCount[p.Height-1] {
		total += c
	}
	return total
}
