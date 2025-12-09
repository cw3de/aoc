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
		panic(err)
	}
	return Parse(strings.Split(string(data), "\n"))
}

func Parse(lines []string) ([]Point, error) {
	list := make([]Point, 0)
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		fields := strings.Split(line, ",")
		if len(fields) != 2 {
			return nil, fmt.Errorf("expected 2 fields but got %d", len(fields))
		}
		x, err := strconv.Atoi(fields[0])
		if err != nil {
			return nil, err
		}
		y, err := strconv.Atoi(fields[1])
		if err != nil {
			return nil, err
		}
		list = append(list, Point{X: x, Y: y})
	}
	return list, nil
}
