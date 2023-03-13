package shapes

type Point struct {
	X float64 // x coordinate
	Y float64 // y coordinate
}

func (p *Point) ParseShape(b []byte) {

}

func (p *Point) GetShapeType() int32 {
	return 1
}

func (p *Point) String() string {
	return "Point"
}

func NewPoint(x float64, y float64) Point {
	return Point{x, y}
}