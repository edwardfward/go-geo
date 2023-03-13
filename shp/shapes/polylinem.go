package shapes

type PolyLineM struct {
	box          [2]Point
	numberParts  int32
	numberPoints int32
	parts        []int32
	points       []PointM
	measures     []float64
	measureRange [2]float64
}

func (p *PolyLineM) ParseShape(b []byte) {

}

func (p *PolyLineM) GetShapeType() int32 {
	return 23
}

func (p *PolyLineM) String() string {
	return "PolyLineM"
}
