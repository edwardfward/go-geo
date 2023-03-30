package records

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
)

type PointRecord struct {
	Shape ShapeType
	point Point
}

type Point struct {
	X float64
	Y float64
}

const (
	POINTRECORDLENGTH = 20 // bytes
	POINTLENGTH       = 16 // bytes
)

func (p *PointRecord) Parse(record []byte) error {
	// check the slice is the right size shape type int32 + two float64s
	if len(record) != POINTRECORDLENGTH {
		return fmt.Errorf("incorrect number of bytes received: %d",
			len(record))
	}

	if err := binary.Read(bytes.NewReader(record), binary.LittleEndian, p); err != nil {
		return fmt.Errorf("error parsing point: %w", err)
	}

	return nil
}

func (p *Point) Parse(pointBytes []byte) error {
	// check the slice is 16-bytes
	if len(pointBytes) != POINTLENGTH {
		var byteError error = errors.New("incorrect number of bytes received")

		return fmt.Errorf("error parsing point: %w", byteError)
	}

	if err := binary.Read(bytes.NewReader(pointBytes), binary.LittleEndian, p); err != nil {
		return fmt.Errorf("error parsing point: %w", err)
	}

	return nil
}

// EmptyPoint returns an empty or default Point.
func EmptyPoint() Point {
	return Point{X: 0, Y: 0}
}

func EmptyPointRecord() PointRecord {
	return PointRecord{Shape: POINT, point: Point{X: 0, Y: 0}}
}
