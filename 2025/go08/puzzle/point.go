package puzzle

import "fmt"

type Point struct {
	X int64
	Y int64
	Z int64
}

func (p Point) String() string {
	return fmt.Sprintf("(%d,%d,%d)", p.X, p.Y, p.Z)
}
