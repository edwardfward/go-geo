package shapes

type PointZ struct {
	x float64 // x coordinate
	y float64 // y coordinate
	z float64 // z coordinate
	m float64 // m measure (e.g. temperature, pressure)
}

func (p *PointZ) ParseShape(b []byte) {

}

func (p *PointZ) GetShapeType() int32 {
	return 11
}

func (p *PointZ) String() string {
	return "PointZ"
}
