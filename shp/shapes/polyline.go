package shapes

import (
	"encoding/binary"
	"go-shp/utils"
)

// PolyLine todo documentation
type PolyLine struct {
	box          [2]Point // bounding box XMin, YMin, XMax, YMax
	numberParts  int32    // number of distinct line segments
	numberPoints int32    // number of total points
	parts        []int32  // stores the index of first point for every line
	points       []Point  // no delimiter between parts
}

// Parse todo documentation
func (p *PolyLine) Parse(r []byte) {
	_shape := r[0:4] // todo not used or needed
	_box := r[4:36]
	_parts := r[36:40]
	_points := r[40:44]

	var shapeType, parts, points int32

	utils.ReadBinary(_shape, binary.LittleEndian, &shapeType)
	utils.ReadBinary(_parts, binary.LittleEndian, &parts)
	utils.ReadBinary(_points, binary.LittleEndian, &points)

	// build bounding box
	p.box = NewBoundaryBox(_box)
	p.numberParts = parts
	p.numberPoints = points

	// build parts array
	index := int32(44)
	for x := int32(0); x < p.numberParts; x++ {
		var part int32
		partSlice := r[index : index+4]
		utils.ReadBinary(partSlice, binary.LittleEndian, &part)
		p.parts = append(p.parts, part)
		index += 4
	}

	// build points array
	for x := int32(0); x < p.numberPoints; x++ {
		pointXSlice := r[index : index+8]
		pointYSlice := r[index+8 : index+16]
		var pointX, pointY float64
		utils.ReadBinary(pointXSlice, binary.LittleEndian, &pointX)
		utils.ReadBinary(pointYSlice, binary.LittleEndian, &pointY)
		p.points = append(p.points, NewPoint(pointX, pointY))
		index += 16
	}
}

// Type todo documentation
func (p *PolyLine) Type() int32 {
	return 3
}

// todo documentation
func (p *PolyLine) String() string {
	return "PolyLine"
}

// New todo documentation
func (p *PolyLine) New() Shape {
	return new(PolyLine)
}
