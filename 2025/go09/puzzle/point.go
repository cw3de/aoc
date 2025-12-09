package puzzle

import "fmt"

type Point struct {
	X int
	Y int
}

func NullPoint() Point {
	return Point{X: 0, Y: 0}
}

func NewPoint(x, y int) Point {
	return Point{X: x, Y: y}
}

func (p *Point) String() string {
	return fmt.Sprintf("[%d,%d]", p.X, p.Y)
}
