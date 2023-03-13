package shapes

type Polygon struct {
	box          [2]Point
	numberParts  int32
	numberPoints int32
	parts        []int32
	points       []Point
}

func (p *Polygon) ParseShape(b []byte) {

}

func (p *Polygon) GetShapeType() int32 {
	return 5
}

func (p *Polygon) NumberParts() int32 {
	return int32(len(p.parts))
}

func (p *Polygon) NumberPoints() int32 {
	return int32(len(p.points))
}

func (p *Polygon) String() string {
	return "Polygon"
}
