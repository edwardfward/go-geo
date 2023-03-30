package records

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
)

type NullShape struct {
	Shape ShapeType
}

const (
	NULLSHAPELENGTH = 4
)

func (n *NullShape) Parse(record []byte) error {
	// check the binary slice is the correct size
	if len(record) != NULLSHAPELENGTH {
		nullShapeFail := errors.New("incorrect number of bytes received for nullshape")

		return fmt.Errorf("%w: received %d needed %d", nullShapeFail,
			len(record), NULLSHAPELENGTH)
	}
	// parse null shape
	if err := binary.Read(bytes.NewReader(record), binary.LittleEndian, n); err != nil {
		return fmt.Errorf("%w: error parsing nullshape", err)
	}

	return nil
}

// LengthBytes returns the null shape record's length in bytes.
func (n *NullShape) LengthBytes() int32 {
	return NULLSHAPELENGTH // content length is number of 16-bit (2 byte) words
}

// EmptyNullShape returns an empty or default null shape.
func EmptyNullShape() NullShape {
	return NullShape{Shape: NULLSHAPE}
}
