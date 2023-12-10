package maze

const (
	DoubleVertical          = "\u2551"
	DoubleHorizontal        = "\u2550"
	DoubleTopLeftCorner     = "\u2554"
	DoubleTopRightCorner    = "\u2557"
	DoubleBottomLeftCorner  = "\u255a"
	DoubleBottomRightCorner = "\u255d"

	SingleVertical          = "\u2502"
	SingleHorizontal        = "\u2500"
	SingleTopLeftCorner     = "\u250c"
	SingleTopRightCorner    = "\u2510"
	SingleBottomLeftCorner  = "\u2514"
	SingleBottomRightCorner = "\u2518"

	FullBlock = "\u2588"
	MiddleDot = "\u00b7"
)

func (m *Maze) Draw() {
	print(DoubleTopLeftCorner)
	for x := 0; x < m.Width(); x++ {
		print(DoubleHorizontal)
	}
	println(DoubleTopRightCorner)

	for _, row := range m.Pipe {
		print(DoubleVertical)
		for _, pipe := range row {
			switch pipe {
			case Ground:
				print(MiddleDot)
			case Start:
				print("S")
			case Outer:
				print("o")
			case Inner:
				print(FullBlock)
			case Vertical:
				print(SingleVertical)
			case Horizontal:
				print(SingleHorizontal)
			case NorthEast:
				print(SingleBottomLeftCorner)
			case NorthWest:
				print(SingleBottomRightCorner)
			case SouthEast:
				print(SingleTopLeftCorner)
			case SouthWest:
				print(SingleTopRightCorner)
			default:
				print(string(pipe))
			}
		}
		print(DoubleVertical)
		println()
	}
	print(DoubleBottomLeftCorner)
	for x := 0; x < m.Width(); x++ {
		print(DoubleHorizontal)
	}
	println(DoubleBottomRightCorner)
}
