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

// ParsePoint returns a single point.
func ParsePoint(pointBytes []byte) (Point, error) {
	// check length to verify 16-bytes received
	if len(pointBytes) != POINTLENGTH {
		pointParseError := errors.New("error parsing point: incorrect length byte" +
			"array received")

		return EmptyPoint(),
			fmt.Errorf("%w: expected %d received %d", pointParseError,
				len(pointBytes), POINTLENGTH)
	}

	point := EmptyPoint()
	// parse x-coordinate
	err := binary.Read(bytes.NewReader(pointBytes[0:8]), binary.LittleEndian, &point.x)
	if err != nil {
		return point,
			fmt.Errorf("new point parse error: unable to parse x-coordinate: %w",
				err)
	}
	// parse y-coordinate
	err = binary.Read(bytes.NewReader(pointBytes[8:16]), binary.LittleEndian, &point.y)
	if err != nil {
		return point,
			fmt.Errorf("%w: unable to parse y-coordinate", err)
	}

	// completed parsing x and y
	return point, nil
}

// ParsePoints returns one or more points for complex shapes (i.e. polyline).
func ParsePoints(points []byte, numPoints int32) ([]Point, error) {
	// check points bytes is the correct length for the number of points requested
	if len(points) != int(numPoints*POINTLENGTH) {
		var pointParseError = errors.New("failed to parse point")

		return nil, fmt.Errorf("%w: received %d bytes for %d points (%d-bytes)",
			pointParseError, len(points), numPoints, numPoints*POINTLENGTH)
	}

	pointArray := make([]Point, numPoints)

	// parse points
	for index := int32(0); index < numPoints; index++ {
		o, err := ParsePoint(points[POINTLENGTH*index : POINTLENGTH*index+POINTLENGTH])
		if err != nil {
			return nil, err
		}

		pointArray[index] = o
	}

	return pointArray, nil
}

// EmptyPoint returns an empty or default Point.
func EmptyPoint() Point {
	return Point{header: EmptyRecordHeader(), shape: POINT, x: 0, y: 0}
}
