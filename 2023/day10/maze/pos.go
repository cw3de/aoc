package maze

import "fmt"

type Pos struct {
	X, Y int
}

func (pos Pos) GoRight() Pos {
	return Pos{pos.X + 1, pos.Y}
}

func (pos Pos) GoLeft() Pos {
	return Pos{pos.X - 1, pos.Y}
}

func (pos Pos) GoUp() Pos {
	return Pos{pos.X, pos.Y - 1}
}

func (pos Pos) GoDown() Pos {
	return Pos{pos.X, pos.Y + 1}
}

func (pos Pos) String() string {
	return fmt.Sprintf("%d,%d", pos.X, pos.Y)
}
