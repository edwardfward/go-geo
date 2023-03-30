package records

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type Polygon struct {
	shape     ShapeType
	box       BoundaryBox
	numParts  int32
	numPoints int32
	parts     []int32
	points    []Point
}

func (p *Polygon) Parse(record []byte) error {
	// get the number of parts and points to set up parsing of
	// the polygon record
	initial := struct {
		Shape     ShapeType
		Box       BoundaryBox
		NumParts  int32
		NumPoints int32
	}{}

	if err := binary.Read(bytes.NewReader(record), binary.LittleEndian, &initial); err != nil {
		return fmt.Errorf("error parsing polygon shape initial struct: %w", err)
	}

	partsAndPoints := struct {
		Parts  []int32
		Points []Point
	}{}

	partsAndPoints.Parts = make([]int32, initial.NumParts)
	partsAndPoints.Points = make([]Point, initial.NumPoints)

	// parse points and parts
	if err := binary.Read(bytes.NewReader(record[44:]), binary.LittleEndian, &partsAndPoints); err != nil {
		return fmt.Errorf("error parsing polygon parts and points array: %w", err)
	}

	// assign
	p.shape = initial.Shape
	p.box = initial.Box
	p.numParts = initial.NumParts
	p.numPoints = initial.NumPoints
	p.parts = partsAndPoints.Parts
	p.points = partsAndPoints.Points

	return nil
}

// LengthBytes returns the polygon's byte length.
func (p *Polygon) LengthBytes() int32 {
	return 0
}

// EmptyPolygon returns and empty or default Polygon shape.
func EmptyPolygon() Polygon {
	return Polygon{
		shape:     POLYGON,
		box:       EmptyBoundaryBox(),
		numParts:  0,
		numPoints: 0,
		parts:     nil,
		points:    nil}
}
