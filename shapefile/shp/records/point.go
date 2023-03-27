package records

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
)

type Point struct {
	header RecordHeader
	shape  ShapeType
	x      float64
	y      float64
}

const (
	POINTLENGTH = 16 // bytes
)

func (p *Point) Parse(record []byte, header RecordHeader) error {
	PointX := record[0:8]  // float64
	PointY := record[8:16] // float64

	err := binary.Read(bytes.NewReader(PointX), binary.LittleEndian, p.x)
	if err != nil {
		return errors.New("failed to parse x coordinate of point record")
	}

	err = binary.Read(bytes.NewReader(PointY), binary.LittleEndian, p.y)
	if err != nil {
		return errors.New("failed to parse y coordinate of point record")
	}

	p.header = header

	return nil
}

func (p *Point) RecordNumber() int32 {
	return p.header.recordNumber
}

func (p *Point) LengthBytes() int32 {
	return p.header.contentLength * WORDMULTIPLE // content length is number of 16-bit (2-byte) words
}

func ParsePoint(point []byte) (Point, error) {
	// check length to verify 16-bytes received
	if len(point) != POINTLENGTH {
		return Point{x: 0, y: 0},
			fmt.Errorf("new point parse error: expected 16-bytes, received: %d",
				len(point))
	}

	p := Point{x: 0, y: 0}
	// parse x-coordinate
	err := binary.Read(bytes.NewReader(point[0:8]), binary.LittleEndian, &p.x)
	if err != nil {
		return p,
			fmt.Errorf("new point parse error: unable to parse x-coordinate: %w",
				err)
	}
	// parse y-coordinate
	err = binary.Read(bytes.NewReader(point[8:16]), binary.LittleEndian, &p.y)
	if err != nil {
		return p,
			fmt.Errorf("new point parse error: unable to parse y-coordinate: %v",
				err)
	}

	// completed parsing x and y
	return p, nil
}
