package records

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"os"
)

func ParseParts(parts []byte, numParts int32) ([]int32, error) {
	// check right number of bytes received to parts number of parts
	if len(parts) != int(numParts*INT32LENGTH) {
		var partsParseError = errors.New("parts parse failure")

		return nil, fmt.Errorf("%w : received %d bytes for %d parts", partsParseError,
			len(parts), numParts)
	}

	partsArray := make([]int32, numParts)

	for index := 0; index < int(numParts); index++ {
		var part int32

		err := binary.Read(bytes.NewReader(parts[index*INT32LENGTH:index*INT32LENGTH+INT32LENGTH]),
			binary.LittleEndian, &part)
		if err != nil {
			return nil, fmt.Errorf("%w: failed to parse part", err)
		}

		partsArray[index] = part
	}

	return partsArray, nil
}

func ParseRecords(file *os.File, shape ShapeType) ([]Record, error) {
	var records []Record

	for {
		// read record header bytes from shapefile
		recordHeaderBytes := make([]byte, RECORDHEADERLENGTH)

		// check to ensure correct number of bytes read
		numBytes, err := file.Read(recordHeaderBytes)
		if numBytes != RECORDHEADERLENGTH || err != nil {
			// reached the end of the file
			if errors.Is(err, io.EOF) {
				break
			}

			return records, fmt.Errorf("%w: unable to parse record header", err)
		}

		// create record header
		recordHeader := EmptyRecordHeader()

		err = recordHeader.Parse(recordHeaderBytes)
		if err != nil {
			return records, fmt.Errorf("%w: unable to parse record header", err)
		}

		// create record
		recordBytes := make([]byte, recordHeader.LengthBytes())

		numBytes, err = file.Read(recordBytes)

		if numBytes != recordHeader.LengthBytes() || err != nil {
			return records, fmt.Errorf("%w: unable to read record bytes for record %d",
				err, recordHeader.RecordNumber())
		}

		record, recordError := GetShapeType(shape)
		if recordError != nil {
			return records, fmt.Errorf("%w: unable to get record shape type",
				recordError)
		}

		err = record.Parse(recordBytes, recordHeader)

		if err != nil {
			return records, fmt.Errorf("%w: error parsing record", err)
		}

		records = append(records, record)
	}

	return records, nil
}
