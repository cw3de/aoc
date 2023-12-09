package day5

import "golang.org/x/exp/slices"

type Range struct {
	start int // included
	end   int // not included
}

func NewRange(start, end int) Range {
	return Range{start, end}
}

func NewRangeWithSize(start, count int) Range {
	return Range{start, start + count}
}

func (r Range) IsEmpty() bool {
	return r.end <= r.start
}

func (r Range) Size() int {
	return r.end - r.start
}

func (r Range) ContainsValue(value int) bool {
	return value >= r.start && value < r.end
}

func (r Range) FirstValue() int {
	return r.start
}

func (r Range) LastValue() int {
	return r.end - 1
}

func (r Range) OverlapsWith(other Range) bool {
	// r . [...]
	// other ... [...]
	return r.start < other.end && r.end > other.start
}

// gitb die Schnittmenge zweier Ranges zurück
func GetIntersection(r1, r2 Range) Range {
	// r1 soll immer vor r2 kommen
	if r1.start > r2.start {
		r1, r2 = r2, r1
	}
	if r1.end < r2.start {
		// r1 ended, bevor r2 beginnt, keine Schnittmenge
		return NewRange(0, 0)
	}
	if r1.end < r2.end {
		// r1 liegt unterhalb von r2
		return NewRange(r2.start, r1.end)
	}
	// r2 liegt komplett in r1
	return r2
}

func SortRanges(ranges []Range) []Range {
	slices.SortFunc(ranges, func(a, b Range) int {
		return a.start - b.start
	})
	return ranges
}

func JoinRanges(ranges []Range) []Range {
	result := []Range{}
	ranges = SortRanges(ranges)
	for _, r := range ranges {
		if len(result) == 0 {
			result = append(result, r)
		} else {
			last := result[len(result)-1]
			if last.end >= r.start {
				// concatenate
				result[len(result)-1] = NewRange(last.start, r.end)
			} else {
				result = append(result, r)
			}
		}
	}
	return result

}
