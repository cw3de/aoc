package day5

// maps eine Interval von Zahlen auf ein anderes Interval ab
type RangeMap struct {
	DestinationStart int
	Source           Range
}

func NewRangeMapWithSize(destinationStart, sourceStart, count int) RangeMap {
	return RangeMap{destinationStart, NewRangeWithSize(sourceStart, count)}
}

// value muss in rm.Source liegen
func (rm RangeMap) MapValue(value int) int {
	if value < rm.Source.FirstValue() {
		panic("RangeMap.MapValue: value < rm.Source.FirstValue")
	}
	if value > rm.Source.LastValue() {
		panic("RangeMap.MapValue: value > rm.Source.End")
	}
	return value + rm.DestinationStart - rm.Source.FirstValue()
}

// r muss vollständig in rm.Source liegen
func (rm RangeMap) MapRange(r Range) Range {
	if r.FirstValue() < rm.Source.FirstValue() {
		panic("RangeMap.MapRange: r.FirstValue < rm.Source.FirstValue")
	}
	if r.LastValue() > rm.Source.LastValue() {
		panic("RangeMap.MapRange: r.LastValue > rm.Source.LastValue")
	}
	return NewRangeWithSize(rm.MapValue(r.FirstValue()), r.Size())
}
