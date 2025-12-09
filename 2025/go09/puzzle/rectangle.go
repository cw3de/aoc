package puzzle

import "fmt"

type Rectangle struct {
	Min Point
	Max Point
}

// coordinates are reordered, so that Min.X <= Max.X and Min.Y <= Max.Y
func NewRectangle(a, b Point) Rectangle {
	xMin, xMax := GetMinMax(a.X, b.X)
	yMin, yMax := GetMinMax(a.Y, b.Y)

	return Rectangle{
		Min: NewPoint(xMin, yMin),
		Max: NewPoint(xMax, yMax),
	}
}

func (r Rectangle) Left() int   { return r.Min.X }
func (r Rectangle) Right() int  { return r.Max.X }
func (r Rectangle) Top() int    { return r.Min.Y }
func (r Rectangle) Bottom() int { return r.Max.Y }

// func (r Rectangle) TopLeft() Point     { return NewPoint(r.Min.X, r.Min.Y) }
// func (r Rectangle) TopRight() Point    { return NewPoint(r.Max.X, r.Min.Y) }
// func (r Rectangle) BottomLeft() Point  { return NewPoint(r.Min.X, r.Max.Y) }
// func (r Rectangle) BottomRight() Point { return NewPoint(r.Max.X, r.Max.Y) }

// the bounderies are included, so that a rectangle with Min == Max
// has a size of 1 rather than 0
func (r Rectangle) GetSize() int {
	return (r.Right() - r.Left() + 1) * (r.Bottom() - r.Top() + 1)
}

// true if point is inside (not on the edge)
func (r Rectangle) PointInside(p Point) bool {
	return p.X > r.Min.X && p.X < r.Max.X &&
		p.Y > r.Min.Y && p.Y < r.Max.Y
}

func (r Rectangle) String() string {
	return fmt.Sprintf("%v-%v", r.Min, r.Max)
}
