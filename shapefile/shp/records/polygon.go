package records

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
)

type Polygon struct {
	header    RecordHeader
	shape     ShapeType
	box       BoundaryBox
	numParts  int32
	numPoints int32
	parts     []int32
	points    []Point
}

func (p *Polygon) Parse(record []byte, header RecordHeader) error {
	// parse shape type
	err := binary.Read(bytes.NewReader(record[0:4]), binary.LittleEndian, &p.shape)
	if err != nil {
		return fmt.Errorf("%w: failed to parse polygon record shape type", err)
	}
	// check shape type
	if p.shape != POLYGON {
		var polygonShapeParseError = errors.New("incorrect shape type parsed for polygon")

		return fmt.Errorf("%w: parsed %d", polygonShapeParseError, &p.shape)
	}

	// parse boundary box
	p.box, err = ParseBoundaryBox(record[4:36])
	if err != nil {
		return fmt.Errorf("%w: error parsing polygon boundary box", err)
	}

	// polygon number of parts
	err = binary.Read(bytes.NewReader(record[36:40]), binary.LittleEndian, &p.numParts)
	if err != nil || p.numParts < 0 {
		return fmt.Errorf("%w: error parsing number of parts for polygon", err)
	}

	// polygon number of points
	err = binary.Read(bytes.NewReader(record[40:44]), binary.LittleEndian,
		&p.numPoints)
	if err != nil || p.numPoints < 0 {
		return fmt.Errorf("%w: error parsing number of points for polygon %x",
			err, record[40:44])
	}
	// polygon part array
	p.parts, err = ParseParts(record[44:44+p.numParts*INT32LENGTH], p.numParts)
	if err != nil {
		return fmt.Errorf("%w: error parsing polygon parts array", err)
	}

	// polygon point array
	p.points, err = ParsePoints(record[44+p.numParts*INT32LENGTH:], p.numPoints)
	if err != nil {
		return fmt.Errorf("%w: error parsing polygon points array", err)
	}

	return nil
}

// LengthBytes returns the polygon's byte length.
func (p *Polygon) LengthBytes() int32 {
	return p.header.contentLength * WORDMULTIPLE
}

// RecordNumber returns the polygon's record number.
func (p *Polygon) RecordNumber() int32 {
	return p.header.recordNumber
}

// EmptyPolygon returns and empty or default Polygon shape.
func EmptyPolygon() Polygon {
	return Polygon{header: EmptyRecordHeader(),
		shape:     POLYGON,
		box:       EmptyBoundaryBox(),
		numParts:  0,
		numPoints: 0,
		parts:     nil,
		points:    nil}
}
