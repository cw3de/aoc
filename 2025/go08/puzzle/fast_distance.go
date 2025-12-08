package puzzle

func FastDistance(a, b Point) int64 {
	dx := a.X - b.X
	dy := a.Y - b.Y
	dz := a.Z - b.Z

	// we dont need the actual value, just the order to compare
	// so we dont habe to take the square-root
	return dx*dx + dy*dy + dz*dz
}
