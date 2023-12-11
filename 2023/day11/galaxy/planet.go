package galaxy

type Planet struct {
	X, Y int
}

func FindPlanets(g *Galaxy) []Planet {
	var planets []Planet
	for y := 0; y < g.Height; y++ {
		for x := 0; x < g.Width; x++ {
			if g.Raster[y][x] == '#' {
				planets = append(planets, Planet{x, y})
			}
		}
	}
	return planets
}

func (g *Galaxy) Distance(p1, p2 Planet, emptyExtraDistance int) int {
	if p1.X > p2.X {
		p1.X, p2.X = p2.X, p1.X
	}
	if p1.Y > p2.Y {
		p1.Y, p2.Y = p2.Y, p1.Y
	}

	dx := p2.X - p1.X
	for x := p1.X; x < p2.X; x++ {
		if g.ColEmpty[x] {
			dx += emptyExtraDistance
		}
	}
	dy := p2.Y - p1.Y
	for y := p1.Y; y < p2.Y; y++ {
		if g.RowEmpty[y] {
			dy += emptyExtraDistance
		}
	}
	return dx + dy
}
