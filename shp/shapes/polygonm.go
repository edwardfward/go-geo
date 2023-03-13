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

func (p *PolygonM) ParseShape(b []byte) {

}

func (p *PolygonM) GetShapeType() int32 {
	return 25
}

func (p *PolygonM) String() string {
	return "PolygonM"
}

func (p *PolygonM) Copy() Shape {
	return new(PolygonM)
}
