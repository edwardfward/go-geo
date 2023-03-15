package shapes

type NullShape struct {
}

func (n *NullShape) Parse([]byte) {

}

func (n *NullShape) Type() int32 {
	return 0
}

func (n *NullShape) String() string {
	return "NullShape"
}

func (n *NullShape) New() Shape {
	return new(NullShape)
}
