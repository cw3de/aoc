package puzzle

import (
	"fmt"
	"slices"
)

type Junction struct {
	Distance int64
	PointA   int
	PointB   int
}

func MakeSortedJunctionList(points []Point, verbose bool) []Junction {

	maxCount := (len(points) - 1) * len(points) / 2
	listOfJunctions := make([]Junction, 0, maxCount)

	for a := 0; a < len(points); a++ {
		for b := a + 1; b < len(points); b++ {

			listOfJunctions = append(listOfJunctions, Junction{
				Distance: FastDistance(points[a], points[b]),
				PointA:   a,
				PointB:   b,
			})
		}
	}

	slices.SortFunc(listOfJunctions, func(a, b Junction) int {
		return int(a.Distance - b.Distance)
	})

	if verbose {
		for i, j := range listOfJunctions {
			fmt.Printf("%d %d:%v - %d:%v = %d\n", i,
				j.PointA, points[j.PointA], j.PointB, points[j.PointB], j.Distance)

			if i > 39 {
				break
			}
		}
	}

	return listOfJunctions
}
