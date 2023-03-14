package shapes

import (
	"encoding/binary"
	"go-shp/utils"
)

// todo documentation
type PolyLine struct {
	box          [2]Point // bounding box XMin, YMin, XMax, YMax
	numberParts  int32    // number of distinct line segments
	numberPoints int32    // number of total points
	parts        []int32  // stores the index of first point for every line
	points       []Point  // no delimiter between parts
}

// todo documentation
func (p *PolyLine) ParseShape(r []byte) {
	shapeTypeSlice := r[0:4]
	xMinSlice := r[4:12]
	yMinSlice := r[12:20]
	xMaxSlice := r[20:28]
	yMaxSlice := r[28:36]
	numberPartsSlice := r[36:40]
	numberPointsSlice := r[40:44]

	var shapeType, numberParts, numberPoints int32
	var xMin, yMin, xMax, yMax float64

	utils.ReadBinary(shapeTypeSlice, binary.LittleEndian, &shapeType)
	utils.ReadBinary(xMinSlice, binary.LittleEndian, &xMin)
	utils.ReadBinary(yMinSlice, binary.LittleEndian, &yMin)
	utils.ReadBinary(xMaxSlice, binary.LittleEndian, &xMax)
	utils.ReadBinary(yMaxSlice, binary.LittleEndian, &yMax)
	utils.ReadBinary(numberPartsSlice, binary.LittleEndian, &numberParts)
	utils.ReadBinary(numberPointsSlice, binary.LittleEndian, &numberPoints)

	// build bounding box
	p.box[0] = NewPoint(xMin, yMin)
	p.box[1] = NewPoint(xMax, yMax)
	p.numberParts = numberParts
	p.numberPoints = numberPoints

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

// todo documentation
func (p *PolyLine) GetShapeType() int32 {
	return 3
}

// todo documentation
func (p *PolyLine) String() string {
	return "PolyLine"
}

// todo documentation
func (p *PolyLine) Copy() Shape {
	return new(PolyLine)
}
