package shapes

import (
	"encoding/binary"
	"go-shp/utils"
	"log"
)

// Polygon todo documentation
type Polygon struct {
	box          [2]Point
	numberParts  int32
	numberPoints int32
	parts        []int32
	points       []Point
}

// Parse todo documentation
func (p *Polygon) Parse(r []byte) {
	_shapeType := r[0:4] // shape type int32
	_box := r[4:36]
	_parts := r[36:40]
	_points := r[40:44]

	var parts, points int32
	var shapeType ShapeType

	// verify the shape is correct
	utils.ReadBinary(_shapeType, binary.LittleEndian, &shapeType)
	if p.Type() != shapeType {
		log.Fatalf("polygon.go: incorrect shape parsed")
	}

	p.box = NewBoundaryBox(_box) // set boundary box

	// parse parts
	utils.ReadBinary(_parts, binary.LittleEndian, &parts)
	for x := int32(0); x < parts; x++ {
		offset := x * 4
		_p := r[44+offset : 48+offset]
		var pNum int32
		utils.ReadBinary(_p, binary.LittleEndian, &pNum)
		p.parts = append(p.parts, pNum)
	}
	// parse polygon parts
	utils.ReadBinary(_points, binary.LittleEndian, &points)
	start := 44 + 4*parts // byte offset to start parsing points

	for x := int32(0); x < points; x++ {
		offset := x * 16
		p.points = append(p.points, ParseNewPoint(r[start+offset:start+offset+16]))
	}
}

// Type todo documentation
func (p *Polygon) Type() ShapeType {
	return POLYGON
}

// NumberParts todo documentation
func (p *Polygon) NumberParts() int32 {
	return int32(len(p.parts))
}

// NumberPoints todo documentation
func (p *Polygon) NumberPoints() int32 {
	return int32(len(p.points))
}

func (p *Polygon) String() string {
	return "Polygon"
}

// New todo documentation
func (p *Polygon) New() Shape {
	return new(Polygon)
}
