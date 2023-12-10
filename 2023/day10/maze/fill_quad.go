package maze

func (qm *QuadMaze) FillOutside() {
	qm.qWalkCell(Pos{0, 0})
}

func (qm *QuadMaze) FillInside() int {
	count := 0
	for y := 2; y < qm.qHeight(); y += 2 {
		for x := 3; x < qm.qWidth(); x += 2 {
			p := Pos{x, y}
			if qm.qGet(p) == Ground {
				qm.qSet(p, Inner)
				qm.maze.Set(Pos{(x - 1) / 2, (y - 1) / 2}, Inner)
				count++
			}
		}
	}
	return count
}

func (qm *QuadMaze) qHeight() int {
	return len(qm.Block)
}

func (qm *QuadMaze) qWidth() int {
	return len(qm.Block[0])
}

func (qm *QuadMaze) qGet(p Pos) byte {
	return qm.Block[p.Y][p.X]
}

func (qm *QuadMaze) qSet(p Pos, sym byte) {
	qm.Block[p.Y][p.X] = sym
}

func (qm *QuadMaze) qNotWalked(p Pos) bool {
	return !qm.Walked[p.Y][p.X]
}

func (qm *QuadMaze) qSetWalked(p Pos) {
	qm.Walked[p.Y][p.X] = true
}

func (qm *QuadMaze) qWalkCell(cur Pos) {
	qm.qSetWalked(cur)

	if qm.qGet(cur) == Ground {
		qm.qSet(cur, Outer)
	}

	// check if we can walk to the right
	if cur.X < qm.qWidth()-1 {
		right := cur.GoRight()
		if qm.qNotWalked(right) {
			if qm.qGet(right) == Ground {
				// fmt.Println("walk right ", cur, "->", right)
				qm.qWalkCell(right)
			}
		}
	}

	// check if we can walk to the left
	if cur.X > 0 {
		left := cur.GoLeft()
		if qm.qNotWalked(left) {
			if qm.qGet(left) == Ground {
				// fmt.Println("walk left ", cur, "->", left)
				qm.qWalkCell(left)
			}
		}
	}

	// check if we can walk down
	if cur.Y < qm.qHeight()-1 {
		down := cur.GoDown()
		if qm.qNotWalked(down) {
			if qm.qGet(down) == Ground {
				// fmt.Println("walk down ", cur, "->", down)
				qm.qWalkCell(down)
			}
		}
	}

	// check if we can walk up
	if cur.Y > 0 {
		up := cur.GoUp()
		if qm.qNotWalked(up) {
			if qm.qGet(up) == Ground {
				// fmt.Println("walk up ", cur, "->", up)
				qm.qWalkCell(up)
			}
		}
	}
}
