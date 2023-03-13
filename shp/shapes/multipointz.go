package shapes

type MultiPointZ struct {
	box          [2]Point
	numberPoints int32
	points       []PointZ
}

func (m *MultiPointZ) ParseShape(b []byte) {

}

func (m *MultiPointZ) GetShapeType() int32 {
	return 18
}

func (m *MultiPointZ) String() string {
	return "MultiPointZ"
}

func (m *MultiPointZ) Copy() Shape {
	return new(MultiPointZ)
}
