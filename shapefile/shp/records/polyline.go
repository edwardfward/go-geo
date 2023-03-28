package records

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
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
		var polylineByteLengthError = errors.New("incorrect number of bytes received")

		return fmt.Errorf("%w: received %d expected %d", polylineByteLengthError,
			len(record), header.LengthBytes())
	}

	// parse shape type
	err := binary.Read(bytes.NewReader(record[0:4]), binary.LittleEndian, &p.shape)
	if err != nil {
		return errors.New("failed to parse polyline shape type")
	}

	// check shape type
	if p.shape != POLYLINE {
		return errors.New("parsed shape type does not match")
	}

	// parse polyline boundary box
	p.box, err = ParseBoundaryBox(record[4:36])
	if err != nil {
		return fmt.Errorf("%w: failed to parse polyline boundary box", err)
	}

	// parse number of polyline parts
	err = binary.Read(bytes.NewReader(record[36:40]), binary.LittleEndian, &p.numParts)
	if err != nil {
		return fmt.Errorf("%w: failed to read polyline number of parts", err)
	}

	p.parts = make([]int32, p.numParts)

	// parse number of points
	err = binary.Read(bytes.NewReader(record[40:44]), binary.LittleEndian, &p.numPoints)
	if err != nil {
		return fmt.Errorf("%w: failed to read number of polyline points", err)
	}

	p.parts, err = ParseParts(record[44:44+p.numParts*INT32LENGTH], p.numParts)
	if err != nil {
		return fmt.Errorf("%w: failed to read number of polyline parts", err)
	}

	p.points, err = ParsePoints(record[44+p.numParts*INT32LENGTH:], p.numPoints)
	if err != nil {
		return fmt.Errorf("%w: failed to parse polyline points", err)
	}

	return nil
}

// RecordNumber returns the polyline's record number.
func (p *PolyLine) RecordNumber() int32 {
	return p.header.recordNumber
}

// LengthBytes returns the length of the polyline record in bytes
// does not include the header (4 bytes).
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
