package shapes

type MultiPointZ struct {
	box          [2]Point
	numberPoints int32
	points       []PointZ
}

func (m *MultiPointZ) Parse(b []byte) {

}

func (m *MultiPointZ) Type() int32 {
	return 18
}

func (m *MultiPointZ) String() string {
	return "MultiPointZ"
}

func (m *MultiPointZ) New() Shape {
	return new(MultiPointZ)
}
