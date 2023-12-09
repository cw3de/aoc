package day5

// kombiniert mehrere RangeMaps zu einer einzigen
type TransformationStep struct {
	Name            string // name der Transformation
	ListOfRangeMaps []RangeMap
}

func NewPlantMap(name string) *TransformationStep {
	return &TransformationStep{Name: name}
}
