package records

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
)

type RecordHeader struct {
	recordNumber  int32 // record number big endian
	contentLength int32 // 16-bit words big endian
}

const (
	RECORDHEADERLENGTH = 8 // bytes
)

func (h *RecordHeader) LengthBytes() int {
	return int(h.contentLength * WORDMULTIPLE)
}

func (h *RecordHeader) RecordNumber() int {
	return int(h.recordNumber)
}

func (h *RecordHeader) Parse(headerSlice []byte) error {
	// check record header is 4 words (8 bytes) long
	if len(headerSlice) != RECORDHEADERLENGTH {
		incorrectHeaderLength := errors.New("unable to parse record header")

		return fmt.Errorf("%w: expected %d bytes, received %d", incorrectHeaderLength,
			RECORDHEADERLENGTH,
			len(headerSlice))
	}

	// parse record number
	err := binary.Read(bytes.NewReader(headerSlice[0:4]), binary.BigEndian,
		&h.recordNumber)
	if err != nil {
		return fmt.Errorf("%w: unable to read record number bytes", err)
	}

	// parse contentlength in 16-bit (2 byte) words
	err = binary.Read(bytes.NewReader(headerSlice[4:8]), binary.BigEndian,
		&h.contentLength)
	if err != nil {
		return fmt.Errorf("%w: unable to read content length bytes", err)
	}

	return nil
}

// EmptyRecordHeader returns an empty RecordHeader.
func EmptyRecordHeader() RecordHeader {
	return RecordHeader{recordNumber: 0, contentLength: 0}
}
