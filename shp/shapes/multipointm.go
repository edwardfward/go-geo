package shapes

import (
	"encoding/binary"
	"go-shp/utils"
)

// MultiPointM todo documentation
type MultiPointM struct {
	box          [2]Point
	numberPoints int32
	points       []Point
	measureRange [2]float64
	measures     []float64
}

// Parse todo documentation
func (m *MultiPointM) Parse(r []byte) {
	_shape := r[0:4]    // int32 little endian
	_box := r[4:36]     // xMin, yMin, xMax, yMax float64 little endian
	_points := r[36:40] // number of points int32 little endian

	var shapeType ShapeType
	var points int32
	utils.ReadBinary(_shape, binary.LittleEndian, &shapeType)
	utils.ReadBinary(_points, binary.LittleEndian, &points)

	// set boundary box
	m.box = NewBoundaryBox(_box)

	for x := int32(0); x < points; x++ {
		offset := 40 + 16*x
		_point := r[40+offset : 56+offset]
		m.points = append(m.points, ParseNewPoint(_point))
	}

	// check if measure data is present
	offset := int(40 + 16*points)
	if len(r) > offset {
		_mMin := r[offset : offset+8]
		_mMax := r[offset+8 : offset+16]

		var mMin, mMax float64
		utils.ReadBinary(_mMin, binary.LittleEndian, &mMin)
		utils.ReadBinary(_mMax, binary.LittleEndian, &mMax)
		m.measureRange = [2]float64{mMin, mMax}
		offset = offset + 16

		for x := int32(0); x < points; x++ {
			offset += 8 * int(x) // shape header + points + mMin + mMax
			_m := r[offset : offset+8]
			var measure float64
			utils.ReadBinary(_m, binary.LittleEndian, &m)
			m.measures = append(m.measures, measure)
		}

	}

}

// Type todo documentation
func (m *MultiPointM) Type() ShapeType {
	return MULTIPOINTM
}

func (m *MultiPointM) String() string {
	return "MultiPointM"
}

// New todo documentation
func (m *MultiPointM) New() Shape {
	return new(MultiPointM)
}
