package shapes

type PointM struct {
	X float64 // x coordinate
	Y float64 // y coordinate
	M float64 // measure
}

func (p *PointM) ParseShape(b []byte) {

}

func (p *PointM) GetShapeType() int32 {
	return 21
}

func (p *PointM) String() string {
	return "PointM"
}

func (p *PointM) Copy() Shape {
	return new(PointM)
}
