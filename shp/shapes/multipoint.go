package shapes

type MultiPoint struct {
	box          [2]Point
	numberPoints int32
	points       []Point
}

func (m *MultiPoint) ParseShape(b []byte) {

}

func (m *MultiPoint) GetShapeType() int32 {
	return 8
}

func (m *MultiPoint) String() string {
	return "MultiPoint"
}

func NewMultiPoint(b []byte) (*MultiPoint, error) {
	return &MultiPoint{}, nil
}

func (m *MultiPoint) Copy() Shape {
	return new(MultiPoint)
}
