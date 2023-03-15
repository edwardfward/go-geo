package shapes

type PolygonZ struct {
	box    [2]Point
	parts  []int32
	points []PointZ
}

func (p *PolygonZ) Parse(b []byte) {

}

func (p *PolygonZ) Type() int32 {
	return 15
}

func (p *PolygonZ) String() string {
	return "PolygonZ"
}

func (p *PolygonZ) New() Shape {
	return new(PolygonZ)
}
