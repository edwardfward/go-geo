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

	// look into making this a function, repeated
	parts := make([]int32, initial.NumParts)
	points := make([]Point, initial.NumPoints)

	partsStart := 44
	partsEnd := partsStart + int(initial.NumParts)*4

	// parse parts
	if err := binary.Read(bytes.NewReader(record[partsStart:partsEnd]),
		binary.LittleEndian, &parts); err != nil {
		return fmt.Errorf("error parsing polygon parts array: %w", err)
	}

	// parse points
	if err := binary.Read(bytes.NewReader(record[partsEnd:]),
		binary.LittleEndian, &points); err != nil {
		return fmt.Errorf("error parsing polygon points array: %w", err)
	}

	// assign
	p.shape = initial.Shape
	p.box = initial.Box
	p.numParts = initial.NumParts
	p.numPoints = initial.NumPoints
	p.parts = parts
	p.points = points

	return nil
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
