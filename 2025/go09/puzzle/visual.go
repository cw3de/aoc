package puzzle

import "fmt"

type Symbol byte

const (
	Empty Symbol = '.'
	Red   Symbol = '#'
	Green Symbol = 'x'
	Used  Symbol = 'O'
)

type Visual struct {
	Rows [][]Symbol
}

func NewVisual(width, height int) *Visual {
	fmt.Printf("NewVisual(%d,%d)\n", width, height)
	g := make([][]Symbol, 0, height)
	for y := 0; y < height; y++ {
		row := make([]Symbol, 0, width)
		for x := 0; x < width; x++ {
			row = append(row, Empty)
		}
		g = append(g, row)
	}
	return &Visual{Rows: g}
}

func (v *Visual) Show() {
	for y, row := range v.Rows {
		for _, field := range row {
			fmt.Printf("%c", field)
		}
		fmt.Printf("  (%d)\n", y)
	}
	fmt.Printf("0123456789abcdef\n\n")
}

func (v *Visual) Set(p Point, s Symbol) {
	v.Rows[p.Y][p.X] = s
}

func (v *Visual) SetList(points []Point, s Symbol) {
	for _, p := range points {
		v.Set(p, s)
	}
}

func (v *Visual) DrawLines(points []Point) {

	var last Point

	for i, p := range points {
		v.Set(p, Red)
		if i > 0 {
			v.DrawLine(p, last)
		}
		last = p
	}
	v.DrawLine(last, points[0])
}

func (v *Visual) DrawLine(a, b Point) {
	if a.X == b.X {
		// vertical line
		y1, y2 := GetMinMax(a.Y, b.Y)
		for y := y1 + 1; y < y2; y++ {
			v.Rows[y][a.X] = Green
		}
	} else if a.Y == b.Y {
		// horizontal line
		x1, x2 := GetMinMax(a.X, b.X)
		for x := x1 + 1; x < x2; x++ {
			v.Rows[a.Y][x] = Green
		}
	} else {
		panic("diaginal lines not permited")
	}
}

func GetMinMax(a, b int) (int, int) {
	if a > b {
		return b, a
	}
	return a, b
}
