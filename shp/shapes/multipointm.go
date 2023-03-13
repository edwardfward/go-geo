package shapes

type MultiPointM struct {
	box          [2]Point
	numberPoints int32
	points       []PointM
	measures     []float64
	measureRange [2]float64
}

func (m *MultiPointM) ParseShape(b []byte) {

}

func (m *MultiPointM) GetShapeType() int32 {
	return 28
}

func (m *MultiPointM) String() string {
	return "MultiPointM"
}
