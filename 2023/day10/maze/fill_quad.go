package maze

func (qm *QuadMaze) Height() int {
	// return qm.maze.Height()
	return len(qm.Block)
}

func (qm *QuadMaze) Width() int {
	// return qm.maze.Height()
	return len(qm.Block[0])
}

func (qm *QuadMaze) Get(p Pos) byte {
	// return qm.Block[2*p.Y][2*p.X]
	return qm.Block[p.Y][p.X]
}

func (qm *QuadMaze) Set(p Pos, sym byte) {
	// qm.Block[2*p.Y][2*p.X] = sym
	qm.Block[p.Y][p.X] = sym
}

func (qm *QuadMaze) NotWalked(p Pos) bool {
	// return !qm.Walked[2*p.Y][2*p.X]
	return !qm.Walked[p.Y][p.X]
}

func (qm *QuadMaze) SetWalked(p Pos) {
	// qm.Walked[2*p.Y][2*p.X] = true
	qm.Walked[p.Y][p.X] = true
}

func walkCell(qm *QuadMaze, cur Pos) {
	qm.SetWalked(cur)

	if qm.Get(cur) == Ground {
		qm.Set(cur, Outer)
	}

	// check if we can walk to the right
	if cur.X < qm.Width()-1 {
		right := cur.GoRight()
		if qm.NotWalked(right) {
			if qm.Get(right) == Ground {
				// fmt.Println("walk right ", cur, "->", right)
				walkCell(qm, right)
			}
		}
	}

	// check if we can walk to the left
	if cur.X > 0 {
		left := cur.GoLeft()
		if qm.NotWalked(left) {
			if qm.Get(left) == Ground {
				// fmt.Println("walk left ", cur, "->", left)
				walkCell(qm, left)
			}
		}
	}

	// check if we can walk down
	if cur.Y < qm.Height()-1 {
		down := cur.GoDown()
		if qm.NotWalked(down) {
			if qm.Get(down) == Ground {
				// fmt.Println("walk down ", cur, "->", down)
				walkCell(qm, down)
			}
		}
	}

	// check if we can walk up
	if cur.Y > 0 {
		up := cur.GoUp()
		if qm.NotWalked(up) {
			if qm.Get(up) == Ground {
				// fmt.Println("walk up ", cur, "->", up)
				walkCell(qm, up)
			}
		}
	}
}

func (qm *QuadMaze) FillOutside() {
	walkCell(qm, Pos{0, 0})
}

func (qm *QuadMaze) FillInside() int {
	count := 0
	for y := 2; y < qm.Height(); y += 2 {
		for x := 3; x < qm.Width(); x += 2 {
			p := Pos{x, y}
			if qm.Get(p) == Ground {
				qm.Set(p, Inner)
				qm.maze.Set(Pos{(x - 1) / 2, (y - 1) / 2}, Inner)
				count++
			}
		}
	}
	return count
}
