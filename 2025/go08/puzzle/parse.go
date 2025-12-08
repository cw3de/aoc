package puzzle

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Load(filename string) ([]Point, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return Parse(string(data))
}

func Parse(text string) ([]Point, error) {
	lines := strings.Split(text, "\n")
	points := make([]Point, 0)
	for _, line := range lines {
		if len(line) == 0 {
			// ignore empty line
			continue
		}
		fields := strings.Split(line, ",")
		if len(fields) != 3 {
			return nil, fmt.Errorf("expected 3 fields, but got %d", len(fields))
		}

		var coords [3]int64
		for i := 0; i < 3; i++ {
			var err error
			coords[i], err = strconv.ParseInt(fields[i], 10, 64)
			if err != nil {
				return nil, err
			}
		}
		points = append(points, Point{
			X: coords[0],
			Y: coords[1],
			Z: coords[2],
		})
	}
	return points, nil
}
