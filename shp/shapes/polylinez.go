package shapes

// todo documentation
type PolyLineZ struct {
	box    [2]Point
	parts  []int32
	points []PointZ
}

// todo documentation
func (p *PolyLineZ) ParseShape(b []byte) {

}

// todo documentation
func (p *PolyLineZ) GetShapeType() int32 {
	return 13
}

// todo documentation
func (p *PolyLineZ) String() string {
	return "PolyLineZ"
}

// todo documentation
func (p *PolyLineZ) Copy() Shape {
	return new(PolyLineZ)
}
