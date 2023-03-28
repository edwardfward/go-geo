package records

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
)

type NullShape struct {
	header RecordHeader
	shape  ShapeType
}

const (
	NULLSHAPELENGTH = 4
)

func (n *NullShape) Parse(record []byte, header RecordHeader) error {
	if len(record) != NULLSHAPELENGTH {
		nullShapeFail := errors.New("[nullshape.go] nullshape parse failed")

		return fmt.Errorf("%w: incorrect number of bytes: received %d, expected %d",
			nullShapeFail, len(record), NULLSHAPELENGTH)
	}

	// read the first four bytes of the record to determine ShapeType
	var parsedShape ShapeType

	err := binary.Read(bytes.NewReader(record), binary.LittleEndian, &parsedShape)
	if err != nil {
		return fmt.Errorf("nullshape record shapetype parse failed: %w ", err)
	}

	// check to make sure parsed shape type matches record ShapeType
	if n.shape != parsedShape {
		parseShapeType := errors.New("parsed shape type did not match null shape type")

		return fmt.Errorf("%w: parsed: %d needed %d", parseShapeType,
			parsedShape, NULLSHAPE)
	}

	n.header = header

	return nil
}

// RecordNumber returns the null shape's record number.
func (n *NullShape) RecordNumber() int32 {
	return n.header.recordNumber
}

// LengthBytes returns the null shape record's length in bytes.
func (n *NullShape) LengthBytes() int32 {
	return n.header.contentLength * WORDMULTIPLE // content length is number of 16-bit (2 byte) words
}

// RecordType returns the null shape's shape type.
func (n *NullShape) RecordType() ShapeType {
	return n.shape
}

// EmptyNullShape returns an empty or default null shape.
func EmptyNullShape() NullShape {
	return NullShape{header: EmptyRecordHeader(), shape: NULLSHAPE}
}
