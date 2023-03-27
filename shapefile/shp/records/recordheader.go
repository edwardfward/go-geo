package records

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type RecordHeader struct {
	recordNumber  int32 // record number big endian
	contentLength int32 // 16-bit words big endian
}

const (
	RECORDHEADERLENGTH = 8 // bytes
)

func ParseRecordHeader(headerSlice []byte) (RecordHeader, error) {
	header := RecordHeader{recordNumber: 0, contentLength: 0}

	// check record header is 4 words (8 bytes) long
	if len(headerSlice) != RECORDHEADERLENGTH {
		return header,
			fmt.Errorf("error parsing record header: incorrect bytes received: %d",
				len(headerSlice))
	}

	// parse record number
	err := binary.Read(bytes.NewReader(headerSlice[0:4]), binary.BigEndian,
		&header.recordNumber)
	if err != nil {
		return header, fmt.Errorf("unable to parse record number from record header: %w",
			err)
	}

	// parse contentlength in 16-bit (2 byte) words
	err = binary.Read(bytes.NewReader(headerSlice[4:8]), binary.BigEndian,
		&header.contentLength)
	if err != nil {
		return header, fmt.Errorf("unable to parse content length from record header: %w",
			err)
	}

	return header, nil
}
