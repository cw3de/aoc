package puzzle

type Requirements struct {
	A, B int
}

func CountPaths(p *Puzzle) int {
	visited := make([]bool, len(p.Num2Name))
	you := p.MustFind("you")
	out := p.MustFind("out")
	return countFrom(p, visited, you, out, nil)
}

func CountViaDacFft(p *Puzzle) int {
	visited := make([]bool, len(p.Num2Name))
	svr := p.MustFind("svr")
	dac := p.MustFind("dac")
	fft := p.MustFind("fft")
	out := p.MustFind("out")

	return countFrom(p, visited, svr, out, &Requirements{A: dac, B: fft})
}

func countFrom(p *Puzzle, visited []bool, rack, endRack int, req *Requirements) int {

	visited[rack] = true
	count := 0
	for _, next := range p.Racks[rack].Outputs {

		if next == endRack {
			if req == nil {
				count++
			} else {
				if visited[req.A] && visited[req.B] {
					count++
				}
			}
		} else {
			if !visited[next] {
				count += countFrom(p, visited, next, endRack, req)
			}
		}
	}

	visited[rack] = false
	return count
}
