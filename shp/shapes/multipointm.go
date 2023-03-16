package shapes

// MultiPointM todo documentation
type MultiPointM struct {
	box          [2]Point
	numberPoints int32
	points       []PointM
	measures     []float64
	measureRange [2]float64
}

// Parse todo documentation
func (m *MultiPointM) Parse(r []byte) {

}

// Type todo documentation
func (m *MultiPointM) Type() ShapeType {
	return MULTIPOINTM
}

func (m *MultiPointM) String() string {
	return "MultiPointM"
}

// New todo documentation
func (m *MultiPointM) New() Shape {
	return new(MultiPointM)
}
