package shapes

type PolygonM struct {
	box          [2]Point
	numberParts  int32
	numberPoints int32
	parts        []int32
	points       []PointM
	measures     []float64
	measureRange [2]float64
}

func (p *PolygonM) Parse(b []byte) {

}

func (p *PolygonM) Type() ShapeType {
	return POLYLINEM
}

func (p *PolygonM) String() string {
	return "PolygonM"
}

func (p *PolygonM) New() Shape {
	return new(PolygonM)
}
