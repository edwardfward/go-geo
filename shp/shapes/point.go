package shapes

import (
	"encoding/binary"
	"fmt"
	"go-shp/utils"
	"log"
)

type Point struct {
	X float64 // x coordinate
	Y float64 // y coordinate
}

// Parse todo complete documentation
func (p *Point) Parse(r []byte) {
	x_ := r[4:12]  // x coordinate float64
	y_ := r[12:20] // y coordinate float64
	var x, y float64
	utils.ReadBinary(x_, binary.LittleEndian, &x)
	utils.ReadBinary(y_, binary.LittleEndian, &y)
	p.X = x
	p.Y = y
}

// Type todo add documentation
func (p *Point) Type() ShapeType {
	return POINT
}

func (p *Point) String() string {
	return fmt.Sprintf("Point:{%f, %f", p.X, p.Y)
}

// New todo add documentation
func (p *Point) New() Shape {
	return new(Point)
}

// NewPoint todo add documentation
func NewPoint(x float64, y float64) Point {
	return Point{x, y}
}

func ParseNewPoint(p []byte) Point {
	if len(p) != 16 { // x,y float 64
		log.Fatalf("point.go : incorrect byte length to parse point")
	}
	_x := p[0:8]
	_y := p[8:16]
	var x, y float64
	utils.ReadBinary(_x, binary.LittleEndian, &x)
	utils.ReadBinary(_y, binary.LittleEndian, &y)

	return NewPoint(x, y)
}
