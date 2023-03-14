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

// todo documentation
func (p *PolyLineM) ParseShape(b []byte) {

}

// todo documentation
func (p *PolyLineM) GetShapeType() int32 {
	return 23
}

// todo documentation
func (p *PolyLineM) String() string {
	return "PolyLineM"
}

// todo documentation
func (p *PolyLineM) Copy() Shape {
	return new(PolyLineM)
}
