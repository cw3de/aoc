package maze

func (maze *Maze) Draw(usedOnly bool) {
	print("\u2554")
	for x := 0; x < maze.Width(); x++ {
		print("\u2550")
	}
	println("\u2557")

	for y, row := range maze.Pipes {
		print("\u2551")
		for x, pipe := range row {
			if usedOnly && !maze.Used[y][x] {
				print(" ")
			} else {
				switch pipe {
				case Ground:
					print(".")
				case Start:
					print("S")
				case Inner:
					print("\u2588")
				case Vertical:
					print("\u2502")
				case Horizontal:
					print("\u2500")
				case NorthEast:
					print("\u2514")
				case NorthWest:
					print("\u2518")
				case SouthEast:
					print("\u250c")
				case SouthWest:
					print("\u2510")
				default:
					print(string(pipe))
				}
			}
		}
		print("\u2551")
		println()
	}
	print("\u255a")
	for x := 0; x < maze.Width(); x++ {
		print("\u2550")
	}
	println("\u255d")
}
