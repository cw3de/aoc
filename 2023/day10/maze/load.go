package maze

import (
	"os"
	"strings"
)

func readLines(filename string) ([]string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return strings.Split(string(data), "\n"), nil
}

func LoadMaze(filename string) (*Maze, error) {
	lines, err := readLines(filename)
	if err != nil {
		return nil, err
	}
	maze := NewMaze()
	for _, line := range lines {
		if line != "" {
			line = strings.TrimSpace(line)
			maze.Pipes = append(maze.Pipes, []byte(line))
			used := make([]bool, len(line))
			maze.Used = append(maze.Used, used)
		}
	}
	return maze, nil
}
