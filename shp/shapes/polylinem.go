package shapes

// todo documentation
type PolyLineM struct {
	box          [2]Point
	numberParts  int32
	numberPoints int32
	parts        []int32
	points       []PointM
	measures     []float64
	measureRange [2]float64
}

// Parse todo documentation
func (p *PolyLineM) Parse(b []byte) {

}

// Type todo documentation
func (p *PolyLineM) Type() ShapeType {
	return POLYLINEM
}

// todo documentation
func (p *PolyLineM) String() string {
	return "PolyLineM"
}

// New todo documentation
func (p *PolyLineM) New() Shape {
	return new(PolyLineM)
}
