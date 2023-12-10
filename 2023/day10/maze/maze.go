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

type Maze struct {
	Pipe [][]byte
	used [][]bool // pipe has beed used
}

func NewMaze() *Maze {
	return &Maze{}
}

func (m *Maze) Height() int {
	return len(m.Pipe)
}

func (m *Maze) Width() int {
	return len(m.Pipe[0])
}

func (m *Maze) Get(p Pos) byte {
	return m.Pipe[p.Y][p.X]
}

func (m *Maze) Set(p Pos, pipe byte) {
	m.Pipe[p.Y][p.X] = pipe
}
