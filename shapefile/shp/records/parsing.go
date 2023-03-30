package records

import (
	"errors"
	"fmt"
	"io"
	"os"
)

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
				err, recordHeader.RecordNumber)
		}

		record, recordError := GetShapeType(shape)
		if recordError != nil {
			return records, fmt.Errorf("%w: unable to get record shape type",
				recordError)
		}

		if err = record.Parse(recordBytes); err != nil {
			return records, fmt.Errorf("%w: error parsing record", err)
		}

		records = append(records, record)
	}

	return records, nil
}
