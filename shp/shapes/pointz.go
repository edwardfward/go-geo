package shapes

type PointZ struct {
	x float64 // x coordinate
	y float64 // y coordinate
	z float64 // z coordinate
	m float64 // m measure (e.g. temperature, pressure)
}

func (p *PointZ) Parse(b []byte) {

}

func (p *PointZ) Type() ShapeType {
	return POINTZ
}

func (p *PointZ) String() string {
	return "PointZ"
}

func (p *PointZ) New() Shape {
	return new(PointZ)
}
