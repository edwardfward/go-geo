package shapes

type NullShape struct {
}

func (n *NullShape) ParseShape([]byte) {

}

func (n *NullShape) GetShapeType() int32 {
	return 0
}

func (n *NullShape) String() string {
	return "NullShape"
}
