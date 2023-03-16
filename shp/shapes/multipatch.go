package shapes

type MultiPatch struct {
}

func (m *MultiPatch) Parse(r []byte) {

}

func (m *MultiPatch) Type() ShapeType {
	return MULTIPATCH
}

func (m *MultiPatch) String() string {
	return "MultiPatch"
}

func (m *MultiPatch) New() Shape {
	return new(MultiPatch)
}
