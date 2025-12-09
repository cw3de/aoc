package puzzle

import "fmt"

func FindLargestRectangle(points []Point, checkPoints bool, verbose bool) int {

	bestSize := 0

	for i, a := range points {
		for j := i + 1; j < len(points); j++ {
			b := points[j]
			r := NewRectangle(a, b)
			size := r.GetSize()
			if size > bestSize {
				if checkPoints && HasPointInsideRect(points, r, verbose) {
					if verbose {
						fmt.Printf("%d (%d,%d) x %d (%d,%d) = %d has edges inside\n",
							i, a.X, a.Y, j, b.X, b.Y, size)
					}
					continue
				}
				if verbose {
					fmt.Printf("%d (%d,%d) x %d (%d,%d) = %d\n",
						i, a.X, a.Y, j, b.X, b.Y, size)
				}
				bestSize = size
			}
		}
	}
	return bestSize
}

func GetEdge(points []Point, index int) Rectangle {
	if index >= len(points) {
		panic("index out of bounds")
	}
	if index == 0 {
		return NewRectangle(points[len(points)-1], points[0])
	}
	return NewRectangle(points[index-1], points[index])
}

func HasPointInsideRect(points []Point, rect Rectangle, verbose bool) bool {
	for i := range points {
		edge := GetEdge(points, i)

		if edge.Min.X == edge.Max.X {
			// top to down
			for y := edge.Min.Y; y <= edge.Max.Y; y++ {
				pt := NewPoint(edge.Min.X, y)
				if rect.PointInside(pt) {
					if verbose {
						fmt.Printf("found point %v inside %v\n", pt, rect)
					}
					return true
				}
			}
		} else {
			// left to right
			for x := edge.Min.X; x < edge.Max.X; x++ {
				pt := NewPoint(x, edge.Min.Y)
				if rect.PointInside(pt) {
					if verbose {
						fmt.Printf("found point %v inside %v\n", pt, rect)
					}
					return true
				}
			}
		}
	}
	return false
}
