package shapes

type PointM struct {
	X float64 // x coordinate
	Y float64 // y coordinate
	M float64 // measure
}

func (p *PointM) Parse(b []byte) {

}

func (p *PointM) Type() int32 {
	return 21
}

func (p *PointM) String() string {
	return "PointM"
}

func (p *PointM) New() Shape {
	return new(PointM)
}
