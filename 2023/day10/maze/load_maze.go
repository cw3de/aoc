package maze

import (
	"os"
	"strings"
)

func LoadMaze(filename string) (*Maze, error) {
	lines, err := readLines(filename)
	if err != nil {
		return nil, err
	}
	m := NewMaze()
	for _, line := range lines {
		if line != "" {
			line = strings.TrimSpace(line)
			m.Pipe = append(m.Pipe, []byte(line))
			used := make([]bool, len(line))
			m.used = append(m.used, used)
		}
	}
	return m, nil
}

func readLines(filename string) ([]string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return strings.Split(string(data), "\n"), nil
}
