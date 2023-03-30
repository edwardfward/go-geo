package records

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
)

type RecordHeader struct {
	RecordNumber  int32 // record number big endian
	ContentLength int32 // 16-bit words big endian
}

const (
	RECORDHEADERLENGTH = 8 // bytes
)

func (h *RecordHeader) LengthBytes() int {
	return int(h.ContentLength * WORDMULTIPLE)
}

func (h *RecordHeader) Parse(headerSlice []byte) error {
	// check record header is 4 words (8 bytes) long
	if len(headerSlice) != RECORDHEADERLENGTH {
		incorrectHeaderLength := errors.New("unable to parse record header")

		return fmt.Errorf("%w: expected %d bytes, received %d", incorrectHeaderLength,
			RECORDHEADERLENGTH,
			len(headerSlice))
	}

	// parse record
	err := binary.Read(bytes.NewReader(headerSlice), binary.BigEndian, h)
	if err != nil {
		return fmt.Errorf("%w: error parsing main record header", err)
	}

	return nil
}

// EmptyRecordHeader returns an empty RecordHeader.
func EmptyRecordHeader() RecordHeader {
	return RecordHeader{RecordNumber: 0, ContentLength: 0}
}
