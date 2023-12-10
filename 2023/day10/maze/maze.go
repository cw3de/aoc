package maze

// Symbols
const (
	Ground     = '.'
	Start      = 'S'
	Vertical   = '|'
	Horizontal = '-'
	NorthEast  = 'L'
	NorthWest  = 'J'
	SouthEast  = 'F'
	SouthWest  = '7'
	Inner      = 'I'
	Outer      = 'O'
)

// Directions
const (
	DirUp    = 1
	DirDown  = 2
	DirLeft  = 3
	DirRight = 4
)

type Maze struct {
	Pipes [][]byte
	Used  [][]bool
}

func NewMaze() *Maze {
	return &Maze{}
}

func (m *Maze) Height() int {
	return len(m.Pipes)
}

func (m *Maze) Width() int {
	return len(m.Pipes[0])
}

func (m *Maze) Get(p Pos) byte {
	return m.Pipes[p.Y][p.X]
}

func (m *Maze) Set(p Pos, pipe byte) {
	m.Pipes[p.Y][p.X] = pipe
}

func (m *Maze) IsUsed(p Pos) bool {
	return m.Used[p.Y][p.X]
}

func (m *Maze) canGoUp(p Pos) bool {
	if p.Y <= 0 {
		return false
	}
	pipe := m.Get(p)
	return pipe == Vertical || pipe == NorthEast || pipe == NorthWest
}

func (m *Maze) canGoDown(p Pos) bool {
	if p.Y+1 >= m.Height() {
		return false
	}
	pipe := m.Get(p)
	return pipe == Vertical || pipe == SouthEast || pipe == SouthWest
}

func (m *Maze) canGoLeft(p Pos) bool {
	if p.X <= 0 {
		return false
	}
	pipe := m.Get(p)
	return pipe == Horizontal || pipe == NorthWest || pipe == SouthWest
}

func (m *Maze) canGoRight(p Pos) bool {
	if p.X+1 >= m.Width() {
		return false
	}
	pipe := m.Get(p)
	return pipe == Horizontal || pipe == NorthEast || pipe == SouthEast
}

func (m *Maze) FindStart() Pos {
	for y, row := range m.Pipes {
		for x, pipe := range row {
			if pipe == Start {
				return Pos{X: x, Y: y}
			}
		}
	}
	return Pos{-1, -1}
}

func (m *Maze) FindLoop(startReplace byte) int {
	p := m.FindStart()
	if p.X == -1 || p.Y == -1 {
		panic("no start found")
	}

	result := 0
	// find first move from the start
	if p.Y > 0 && m.canGoDown(p.GoUp()) {
		result = m.FollowPipe(p.GoUp(), DirUp) + 1
	} else if p.Y+1 < m.Height() && m.canGoUp(p.GoDown()) {
		result = m.FollowPipe(p.GoDown(), DirDown) + 1
	} else if p.X > 0 && m.canGoRight(p.GoLeft()) {
		result = m.FollowPipe(p.GoLeft(), DirLeft) + 1
	} else if p.X+1 < m.Width() && m.canGoLeft(p.GoRight()) {
		result = m.FollowPipe(p.GoRight(), DirRight) + 1
	} else {
		panic("no way out")
	}
	if true {
		m.Set(p, startReplace) // TODO: found out correct pipe
	}
	return result
}

func (m *Maze) FollowPipe(p Pos, lastStep int) int {
	m.Used[p.Y][p.X] = true
	switch m.Get(p) {
	case Start:
		return 0
	case Vertical: // |
		if lastStep == DirUp {
			return m.FollowPipe(p.GoUp(), DirUp) + 1
		} else if lastStep == DirDown {
			return m.FollowPipe(p.GoDown(), DirDown) + 1
		} else {
			panic("invalid direction")
		}
	case Horizontal: // -
		if lastStep == DirLeft {
			return m.FollowPipe(p.GoLeft(), DirLeft) + 1
		} else if lastStep == DirRight {
			return m.FollowPipe(p.GoRight(), DirRight) + 1
		} else {
			panic("invalid direction")
		}
	case NorthEast: // L
		if lastStep == DirDown {
			return m.FollowPipe(p.GoRight(), DirRight) + 1
		} else if lastStep == DirLeft {
			return m.FollowPipe(p.GoUp(), DirUp) + 1
		} else {
			panic("invalid direction")
		}
	case NorthWest: // J
		if lastStep == DirDown {
			return m.FollowPipe(p.GoLeft(), DirLeft) + 1
		} else if lastStep == DirRight {
			return m.FollowPipe(p.GoUp(), DirUp) + 1
		} else {
			panic("invalid direction")
		}
	case SouthEast: // F
		if lastStep == DirUp {
			return m.FollowPipe(p.GoRight(), DirRight) + 1
		} else if lastStep == DirLeft {
			return m.FollowPipe(p.GoDown(), DirDown) + 1
		} else {
			panic("invalid direction")
		}
	case SouthWest: // 7
		if lastStep == DirUp {
			return m.FollowPipe(p.GoLeft(), DirLeft) + 1
		} else if lastStep == DirRight {
			return m.FollowPipe(p.GoDown(), DirDown) + 1
		} else {
			panic("invalid direction")
		}
	default:
		panic("unknown pipe")
	}
}

func (maze *Maze) ClearUnsed() {
	for y, row := range maze.Used {
		for x, used := range row {
			if !used {
				maze.Pipes[y][x] = Ground
			}
		}
	}
}
