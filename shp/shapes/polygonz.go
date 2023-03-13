package shapes

type PolygonZ struct {
	box    [2]Point
	parts  []int32
	points []PointZ
}

func (p *PolygonZ) ParseShape(b []byte) {

}

func (p *PolygonZ) GetShapeType() int32 {
	return 15
}

func (p *PolygonZ) String() string {
	return "PolygonZ"
}

func (p *PolygonZ) Copy() Shape {
	return new(PolygonZ)
}
