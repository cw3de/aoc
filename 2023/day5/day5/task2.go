package day5

import "fmt"

func MapRanges(seedRanges []Range, listOfRangeMaps []RangeMap) []Range {
	result := []Range{}
	todos := seedRanges
	for len(todos) > 0 {
		seedRange := todos[0]
		todos = todos[1:]
		found := false

		for _, rm := range listOfRangeMaps {
			if rm.Source.OverlapsWith(seedRange) {
				intersection := GetIntersection(rm.Source, seedRange)
				// fmt.Println("intersection ", intersection, " found in ", rm.Source, " and ", seedRange)
				result = append(result, rm.MapRange(intersection))
				found = true

				if intersection.start > seedRange.start {
					// es gibt eine Range vor der intersection
					todos = append(todos, NewRange(seedRange.start, intersection.start))
				}

				if intersection.end < seedRange.end {
					// es gibt eine Range nach der intersection
					todos = append(todos, NewRange(intersection.end, seedRange.end))
				}
				break
			}
		}
		if !found {
			// no mapping found
			result = append(result, seedRange)
		}
	}
	return result
}

func Task2(filename string) {
	almanac := LoadAlmanac(filename)

	ranges := []Range{}
	for i := 0; i < len(almanac.Seeds); i += 2 {
		r0 := NewRangeWithSize(almanac.Seeds[i], almanac.Seeds[i+1])
		ranges = append(ranges, r0)
	}
	// fmt.Println(ranges)

	for _, m := range almanac.ListOfTransformations {
		ranges = MapRanges(ranges, m.ListOfRangeMaps)
	}

	SortRanges(ranges)
	fmt.Println(ranges)
}
