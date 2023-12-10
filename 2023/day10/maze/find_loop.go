package maze

// Directions
const (
	DirUp    = 1
	DirDown  = 2
	DirLeft  = 3
	DirRight = 4
)

func (m *Maze) FindLoop(startReplace byte) int {
	p := m.findStart()
	if p.X == -1 || p.Y == -1 {
		panic("no start found")
	}

	result := 0
	// find first move from the start
	if p.Y > 0 && m.canGoDown(p.GoUp()) {
		result = m.followPipe(p.GoUp(), DirUp) + 1
	} else if p.Y+1 < m.Height() && m.canGoUp(p.GoDown()) {
		result = m.followPipe(p.GoDown(), DirDown) + 1
	} else if p.X > 0 && m.canGoRight(p.GoLeft()) {
		result = m.followPipe(p.GoLeft(), DirLeft) + 1
	} else if p.X+1 < m.Width() && m.canGoLeft(p.GoRight()) {
		result = m.followPipe(p.GoRight(), DirRight) + 1
	} else {
		panic("no way out")
	}
	if true {
		m.Set(p, startReplace) // TODO: found out correct pipe
	}
	return result
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

func (m *Maze) findStart() Pos {
	for y, row := range m.Pipe {
		for x, pipe := range row {
			if pipe == Start {
				return Pos{X: x, Y: y}
			}
		}
	}
	return Pos{-1, -1}
}

func (m *Maze) IsUsed(p Pos) bool {
	return m.used[p.Y][p.X]
}
func (m *Maze) ClearUnsed() {
	for y, row := range m.used {
		for x, used := range row {
			if !used {
				m.Pipe[y][x] = Ground
			}
		}
	}
}
func (m *Maze) followPipe(p Pos, lastStep int) int {
	m.used[p.Y][p.X] = true
	switch m.Get(p) {
	case Start:
		return 0
	case Vertical: // |
		if lastStep == DirUp {
			return m.followPipe(p.GoUp(), DirUp) + 1
		} else if lastStep == DirDown {
			return m.followPipe(p.GoDown(), DirDown) + 1
		} else {
			panic("invalid direction")
		}
	case Horizontal: // -
		if lastStep == DirLeft {
			return m.followPipe(p.GoLeft(), DirLeft) + 1
		} else if lastStep == DirRight {
			return m.followPipe(p.GoRight(), DirRight) + 1
		} else {
			panic("invalid direction")
		}
	case NorthEast: // L
		if lastStep == DirDown {
			return m.followPipe(p.GoRight(), DirRight) + 1
		} else if lastStep == DirLeft {
			return m.followPipe(p.GoUp(), DirUp) + 1
		} else {
			panic("invalid direction")
		}
	case NorthWest: // J
		if lastStep == DirDown {
			return m.followPipe(p.GoLeft(), DirLeft) + 1
		} else if lastStep == DirRight {
			return m.followPipe(p.GoUp(), DirUp) + 1
		} else {
			panic("invalid direction")
		}
	case SouthEast: // F
		if lastStep == DirUp {
			return m.followPipe(p.GoRight(), DirRight) + 1
		} else if lastStep == DirLeft {
			return m.followPipe(p.GoDown(), DirDown) + 1
		} else {
			panic("invalid direction")
		}
	case SouthWest: // 7
		if lastStep == DirUp {
			return m.followPipe(p.GoLeft(), DirLeft) + 1
		} else if lastStep == DirRight {
			return m.followPipe(p.GoDown(), DirDown) + 1
		} else {
			panic("invalid direction")
		}
	default:
		panic("unknown pipe")
	}
}
