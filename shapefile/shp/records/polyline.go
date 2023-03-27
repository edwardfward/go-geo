package records

import (
	"bytes"
	"encoding/binary"
	"errors"
)

type PolyLine struct {
	header    RecordHeader
	shape     ShapeType
	box       BoundaryBox
	numParts  int32
	numPoints int32
	parts     []int32
	points    []Point
}

func (p *PolyLine) Parse(record []byte, header RecordHeader) error {
	// check length of record equals header contentlength in bytes
	if len(record) != header.LengthBytes() {
		return errors.New("error parsing polyline record, record bytes did not match " +
			"header bytes")
	}

	var shape ShapeType

	// parse shape type
	err := binary.Read(bytes.NewReader(record[0:4]), binary.LittleEndian, &shape)
	if err != nil {
		return errors.New("failed to parse polyline shape type")
	}

	// check shape type
	if shape != p.shape {
		return errors.New("parsed shape type does not match")
	}

	// parse polyline boundary box
	p.box, err = ParseBoundaryBox(record[4:36])
	if err != nil {
		return err
	}

	// parse number of polyline parts
	err = binary.Read(bytes.NewReader(record[36:40]), binary.LittleEndian, &p.numParts)
	if err != nil {
		return err
	}

	p.parts = make([]int32, p.numParts)

	// parse number of points
	err = binary.Read(bytes.NewReader(record[40:44]), binary.LittleEndian, &p.numPoints)
	if err != nil {
		return err
	}

	p.parts, err = ParseParts(record[44:44+p.numParts*POINTLENGTH], p.numParts)
	if err != nil {
		return err
	}

	p.points, err = ParsePoints(record[44+p.numParts*POINTLENGTH:], p.numPoints)
	if err != nil {
		return err
	}

	return nil
}

func (p *PolyLine) RecordNumber() int32 {
	return p.header.recordNumber
}

func (p *PolyLine) LengthBytes() int32 {
	return p.header.contentLength * WORDMULTIPLE
}

func EmptyPolyLine() PolyLine {
	return PolyLine{
		header:    EmptyRecordHeader(),
		shape:     POLYLINE,
		box:       EmptyBoundaryBox(),
		numParts:  0,
		numPoints: 0,
		parts:     nil, points: nil,
	}
}
