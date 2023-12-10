package maze

type QuadMaze struct {
	maze   *Maze
	Block  [][]byte
	Walked [][]bool
}

func getSymbol(pipe byte) (string, string) {

	switch pipe {
	case Ground:
		return "..", ".."
	case Start:
		return "..", "S."
	case Vertical:
		return "x.", "x."
	case Horizontal:
		return "..", "xx"
	case NorthEast:
		return "x.", "xx"
	case NorthWest:
		return "x.", "x."
	case SouthEast:
		return "..", "xx"
	case SouthWest:
		return "..", "x."
	case Outer:
		return "..", "O."
	case Inner:
		return "..", "I."
	}
	panic("unreachable")
}

func NewQuadMaze(maze *Maze) *QuadMaze {
	qh := 2*maze.Height() + 1
	qw := 2*maze.Width() + 1
	// fmt.Println("make new quad size", qw, qh, "from", maze.Width(), maze.Height())
	q := &QuadMaze{
		maze:   maze,
		Block:  make([][]byte, qh),
		Walked: make([][]bool, qh)}

	q.Block[0] = make([]byte, qw)
	q.Walked[0] = make([]bool, qw)

	for x := 0; x < qw; x++ {
		q.Block[0][x] = '.'
	}
	for y, row := range maze.Pipes {
		qy := 1 + y*2
		// fmt.Println("make new quad row", qy, "from", y)
		q.Block[qy+0] = make([]byte, qw)
		q.Block[qy+1] = make([]byte, qw)
		q.Walked[qy+0] = make([]bool, qw)
		q.Walked[qy+1] = make([]bool, qw)
		for x, pipe := range row {
			qx := 1 + x*2
			top, bot := getSymbol(pipe)
			q.Block[qy][qx] = top[0]
			q.Block[qy][qx+1] = top[1]
			q.Block[qy+1][qx] = bot[0]
			q.Block[qy+1][qx+1] = bot[1]
		}
		q.Block[qy][0] = '.'
		q.Block[qy+1][0] = '.'
	}

	return q
}

func (q *QuadMaze) Draw() {
	for _, row := range q.Block {
		// print(string(strconv.Itoa(y) + " "))
		for _, sym := range row {
			print(string(sym))
		}
		println()
	}
}
