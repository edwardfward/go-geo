package shapes

import (
	"encoding/binary"
	"go-shp/utils"
	"log"
)

// MultiPoint todo documentation
type MultiPoint struct {
	box          [2]Point
	numberPoints int32
	points       []Point
}

// Parse todo documentation
func (m *MultiPoint) Parse(r []byte) {
	_shape := r[0:4] // shape type not needed
	_box := r[4:36]
	_points := r[36:40]

	var shape, numberPoints int32

	// verify the shape is correct
	utils.ReadBinary(_shape, binary.LittleEndian, &shape)
	if m.Type() != shape {
		log.Fatalf("multipoint.go: incorrect shape parsed")
	}

	// set boundary box
	m.box = NewBoundaryBox(_box)

	// set number of points
	utils.ReadBinary(_points, binary.LittleEndian, &numberPoints)
	m.numberPoints = numberPoints

	// create points
	for x := int32(0); x < numberPoints; x++ {
		offset := x * 16 // each point(x,y) is 16-bytes 2 x float64
		_point := r[40+offset : 56+offset]
		m.points = append(m.points, ParseNewPoint(_point))
	}
}

// Type todo documentation
func (m *MultiPoint) Type() int32 {
	return 8
}

func (m *MultiPoint) String() string {
	return "MultiPoint"
}

// New todo documentation
func (m *MultiPoint) New() Shape {
	return new(MultiPoint)
}
