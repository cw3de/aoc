package puzzle

import (
	"fmt"
	"sort"
)

type MinMax struct {
	Min int64
	Max int64
}

type Puzzle struct {
	FreshRanges []MinMax
}

func NewPuzzle() Puzzle {
	var p Puzzle
	p.FreshRanges = make([]MinMax, 0)
	return p
}

func (p *Puzzle) AddRange(min, max int64) {
	if min > max {
		min, max = max, min
	}

	p.FreshRanges = append(p.FreshRanges, MinMax{
		Min: min,
		Max: max,
	})
}

func (p *Puzzle) SortAndJoin(verbose bool) {
	sort.Slice(p.FreshRanges, func(i, j int) bool {
		return p.FreshRanges[i].Min < p.FreshRanges[j].Min
	})

	n := make([]MinMax, 0, len(p.FreshRanges))
	var last *MinMax = nil

	for _, mm := range p.FreshRanges {
		if last == nil || mm.Min > last.Max {
			// dont overlap
			n = append(n, mm)
			last = &n[len(n)-1]
			if verbose {
				fmt.Printf("%15d - %15d (new)\n", mm.Min, mm.Max)
			}
		} else if mm.Max > last.Max {
			// extend
			last.Max = mm.Max
			if verbose {
				fmt.Printf("%15d - %15d (extend)\n", mm.Min, mm.Max)
			}
		} else {
			// otherwise mm is included by last
			if verbose {
				fmt.Printf("%15d - %15d (ignored)\n", mm.Min, mm.Max)
			}
		}
	}
	p.FreshRanges = n
}

func (p *Puzzle) Contains(id int64) bool {
	for _, mm := range p.FreshRanges {
		if id >= mm.Min && id <= mm.Max {
			return true
		}
	}
	return false
}
