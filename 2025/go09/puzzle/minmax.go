package puzzle

type MinMax struct {
	Min Point
	Max Point
}

func FindMinMax(points []Point) MinMax {

	mm := MinMax{
		Min: NewPoint(0, 0),
		Max: NewPoint(0, 0),
	}

	getMin := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	getMax := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	for i, p := range points {
		if i == 0 {
			mm.Min = p
			mm.Max = p
		} else {
			mm.Min.X = getMin(mm.Min.X, p.X)
			mm.Min.Y = getMin(mm.Min.Y, p.Y)

			mm.Max.X = getMax(mm.Max.X, p.X)
			mm.Max.Y = getMax(mm.Max.Y, p.Y)
		}
	}

	return mm
}
