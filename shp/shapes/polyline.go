package shapes

type PolyLine struct {
	box          [2]Point // bounding box XMin, YMin, XMax, YMax
	numberParts  int32    // number of distinct line segments
	numberPoints int32    // number of total points
	parts        []int32  // stores the index of first point for every line
	points       []Point  // no delimiter between parts
}

func (p *PolyLine) ParseShape(b []byte) {

}

func (p *PolyLine) GetShapeType() int32 {
	return 3
}

func (p *PolyLine) String() string {
	return "PolyLine"
}
