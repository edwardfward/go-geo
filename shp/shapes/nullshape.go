package shapes

type NullShape struct {
}

func (n *NullShape) Parse([]byte) {

}

func (n *NullShape) Type() ShapeType {
	return NULLSHAPE
}

func (n *NullShape) String() string {
	return "NullShape"
}

func (n *NullShape) New() Shape {
	return new(NullShape)
}
