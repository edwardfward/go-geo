package records

import (
	"bytes"
	"encoding/binary"
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
		return fmt.Errorf("nullshape parse failed: incorrect number of bytes (%d) for nullshape",
			len(record))
	}

	// read the first four bytes of the record to determine ShapeType
	var parsedShape ShapeType

	err := binary.Read(bytes.NewReader(record), binary.LittleEndian, &parsedShape)
	if err != nil {
		return fmt.Errorf("nullshape record shapetype parse failed: %w ", err)
	}

	// check to make sure parsed shape type matches record ShapeType
	if n.shape != parsedShape {
		return fmt.Errorf("parsed shapetype (%d) did not match null shapetype (%d)",
			parsedShape, n.shape)
	}

	n.header = header

	return nil
}

func (n *NullShape) RecordNumber() int32 {
	return n.header.recordNumber
}

func (n *NullShape) LengthBytes() int32 {
	return n.header.contentLength * WORDMULTIPLE // content length is number of 16-bit (2 byte) words
}

func (n *NullShape) RecordType() ShapeType {
	return n.shape
}
