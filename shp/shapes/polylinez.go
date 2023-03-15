package shapes

// PolyLineZ todo documentation
type PolyLineZ struct {
	box    [2]Point
	parts  []int32
	points []PointZ
}

// Parse todo documentation
func (p *PolyLineZ) Parse(b []byte) {

}

// Type todo documentation
func (p *PolyLineZ) Type() int32 {
	return 13
}

// todo documentation
func (p *PolyLineZ) String() string {
	return "PolyLineZ"
}

// New todo documentation
func (p *PolyLineZ) New() Shape {
	return new(PolyLineZ)
}
