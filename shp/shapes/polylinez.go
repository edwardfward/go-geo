package shapes

type PolyLineZ struct {
	box    [2]Point
	parts  []int32
	points []PointZ
}

func (p *PolyLineZ) ParseShape(b []byte) {

}

func (p *PolyLineZ) GetShapeType() int32 {
	return 13
}

func (p *PolyLineZ) String() string {
	return "PolyLineZ"
}
