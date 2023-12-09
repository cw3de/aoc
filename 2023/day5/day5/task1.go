package day5

import (
	"fmt"
	"slices"
)

// suche die passende RangeMap und wende sie an
func MapSeed(seed int, tranStep *TransformationStep) int {
	for _, rm := range tranStep.ListOfRangeMaps {
		if rm.Source.ContainsValue(seed) {
			return rm.MapValue(seed)
		}
	}
	return seed
}

// wende alle Transformationen der Reihe nach auf den Seed an
func MapSeedToAllMaps(seed int, listOfAllSteps []*TransformationStep) int {
	result := seed
	for _, tranStep := range listOfAllSteps {
		result = MapSeed(result, tranStep)
	}
	return result
}

func Task1(filename string) {
	almanac := LoadAlmanac(filename)
	// ShowAlmanac(almanac)
	result := []int{}
	for _, seed := range almanac.Seeds {
		result = append(result, MapSeedToAllMaps(seed, almanac.ListOfTransformations))
	}

	slices.Sort(result)
	fmt.Println(filename, ":", result[0])
}
